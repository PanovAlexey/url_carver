package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetChannelsMapService(t *testing.T) {
	t.Run("Test channels map service creating", func(t *testing.T) {
		channelsMapService := GetChannelsMapService()
		structType := fmt.Sprintf("%T", channelsMapService)
		assert.Equal(t, structType, "*services.ChannelsMapService")
	})
}

func Test_GetMap(t *testing.T) {
	t.Run("Test channels map after creating", func(t *testing.T) {
		channelsMapService := GetChannelsMapService()
		channelsMap := channelsMapService.GetMap()

		assert.Equal(t, len(channelsMap), 0)
	})
}

func Test_AddChannelAndGetChannelByName(t *testing.T) {
	t.Run("Test add channel and get channel by name", func(t *testing.T) {
		channelsMapService := GetChannelsMapService()
		channelTestName := "test"
		channelNonExistsTestName := "test-wrong"
		ch := channelsMapService.addChannel(channelTestName)

		assert.Equal(t, ch, channelsMapService.GetChannelByName(channelTestName))
		assert.Equal(
			t,
			channelsMapService.GetChannelByName(channelNonExistsTestName),
			channelsMapService.GetChannelByName(channelNonExistsTestName),
		)
	})
}
