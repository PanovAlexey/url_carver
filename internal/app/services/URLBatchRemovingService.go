package services

import (
	"github.com/PanovAlexey/url_carver/internal/app/services/tools"
	"log"
	"sync"
)

const ChannelWithRemovingURLsName = "url_removing"
const workersURLsRemovingCount = 6

type batchURLsRemovingService struct {
	databaseRepository databaseURLRepositoryInterface
}

func GetBatchURLsRemovingService(databaseRepository databaseURLRepositoryInterface) batchURLsRemovingService {
	return batchURLsRemovingService{
		databaseRepository: databaseRepository,
	}
}

func (service batchURLsRemovingService) RemoveByShortURLSlice(URLSlice []string, userID int) error {
	queueMap := GetChannelsMapService()
	globalInputChannel := queueMap.GetChannelByName(ChannelWithRemovingURLsName) // channel will close in Main

	go func() {
		for _, shortURL := range URLSlice {
			globalInputChannel <- shortURL

			log.Println("Short URL \"" + shortURL + "\" added to queue.")
		}
	}()

	fanOutChannels := getSliceChannelsByFanOut(globalInputChannel)

	workerChannels := make([]chan string, 0, workersURLsRemovingCount)
	for _, channel := range fanOutChannels {
		worker := service.newURMRemovingWorker(channel, userID)
		workerChannels = append(workerChannels, worker)
	}

	// this goroutine is to avoid to wait an answer
	go func() {
		for range fanInChannels(workerChannels...) {
			//fmt.Println("Read from fanIn: " + v)
		}
	}()

	return nil
}

func (service batchURLsRemovingService) newURMRemovingWorker(inputChannel <-chan string, userID int) chan string {
	outChannel := make(chan string)
	toolService := tools.GetToolService()

	go func() {
		for shortURLvalue := range inputChannel {
			resultDeletingMap := make(map[string]bool)
			resultDeletingMap[shortURLvalue] = false

			// @ToDo: add to SQL query batch 3-5 items at a time
			deletedDatabaseURLs, err := service.databaseRepository.
				DeleteURLsByShortValueSlice(toolService.GetKeysSliceByStringToBoolMap(resultDeletingMap), userID)

			if err != nil {
				log.Println("error while URL (" + shortURLvalue + ") deleting: " + err.Error())
				continue
			}

			for _, databaseURL := range deletedDatabaseURLs {
				_, ok := resultDeletingMap[databaseURL.GetShortURL()]

				if ok {
					resultDeletingMap[databaseURL.GetShortURL()] = true
				}
			}

			for shortURLValue, isDeleted := range resultDeletingMap {
				if isDeleted {
					log.Println("Short URL \"" + shortURLValue + "\" deleted successfully")
				} else {
					log.Println("Warning. Short URL \"" + shortURLValue + "\" has not been deleted")
				}
			}

			outChannel <- shortURLvalue
		}

		close(outChannel)
	}()

	return outChannel
}

func getSliceChannelsByFanOut(inputChannel chan string) []chan string {
	batchChannelsSlice := make([]chan string, 0, workersURLsRemovingCount)

	for i := 0; i < workersURLsRemovingCount; i++ {
		outputChannel := make(chan string)
		batchChannelsSlice = append(batchChannelsSlice, outputChannel)
	}

	go func() {
		defer func(batchChannelsSlice []chan string) {
			for _, outputChannel := range batchChannelsSlice {
				close(outputChannel)
			}
		}(batchChannelsSlice)

		for i := 0; ; i++ {
			if i == len(batchChannelsSlice) {
				i = 0
			}

			shortURLValue, ok := <-inputChannel

			if !ok {
				return
			}

			outputChannel := batchChannelsSlice[i]
			outputChannel <- shortURLValue
		}
	}()

	return batchChannelsSlice
}

func fanInChannels(inputChannels ...chan string) chan string {
	resultOutChannel := make(chan string)

	go func() {
		wg := &sync.WaitGroup{}

		for _, inputChannel := range inputChannels {
			wg.Add(1)

			go func(inputChannel chan string) {
				defer wg.Done()
				for item := range inputChannel {
					resultOutChannel <- item
				}
			}(inputChannel)
		}

		wg.Wait()
		close(resultOutChannel)
	}()

	return resultOutChannel
}
