package services

var globalChannelsMapStorage *ChannelsMapService

type ChannelsMapService struct {
	queueMap map[string]chan string
}

func GetChannelsMapService() *ChannelsMapService {
	if globalChannelsMapStorage != nil {
		return globalChannelsMapStorage
	} else {
		queueMap := make(map[string]chan string)
		globalChannelsMapStorage = &ChannelsMapService{queueMap: queueMap}

		return globalChannelsMapStorage
	}
}

func (service ChannelsMapService) GetMap() map[string]chan string {
	return service.queueMap
}

func (service ChannelsMapService) GetChannelByName(name string) chan string {
	channel, ok := service.queueMap[name]

	if !ok {
		return service.addChannel(name)
	}

	return channel
}

func (service ChannelsMapService) addChannel(name string) chan string {
	service.queueMap[name] = make(chan string)

	return service.queueMap[name]
}
