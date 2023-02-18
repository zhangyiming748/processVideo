package processVideo

import (
	"testing"
)

func TestMaster(t *testing.T) {
	ProcessAllVideos("/Users/zen/Downloads/Telegram/telegram/本庄玲", "mp4;avi;wmv", "10", false, true)
}

func TestProcessAllVideos(t *testing.T) {

}
