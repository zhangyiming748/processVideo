package processVideo

import (
	"testing"
)

func TestMaster(t *testing.T) {
	ProcessAllVideos("/Users/zen/Downloads/整理", "mp4;avi;wmv", "10", false)
}

func TestProcessAllVideos(t *testing.T) {

}
