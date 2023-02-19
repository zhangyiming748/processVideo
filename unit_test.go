package processVideo

import (
	"testing"
)

func TestMaster(t *testing.T) {
	ProcessAllVideos("/Users/zen/Downloads/Telegram Desktop/水岛津实", "mp4;avi;wmv", "10", false, false)
}

func TestProcessAllVideos(t *testing.T) {

}
