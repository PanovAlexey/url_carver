package tools

import "reflect"

type toolService struct {
}

func GetToolService() toolService {
	return toolService{}
}

func (service toolService) GetKeysSliceByStringToBoolMap(m map[string]bool) []string {
	mapKeyValuesSlice := reflect.ValueOf(m).MapKeys()
	mapKeyStringSlice := make([]string, 0)

	for _, value := range mapKeyValuesSlice {
		mapKeyStringSlice = append(mapKeyStringSlice, value.String())
	}

	return mapKeyStringSlice
}
