package processVideo

import (
	"path"
	"testing"
)

func TestMaster(t *testing.T) {
	ConvAllVideos2H265("/Users/zen/Downloads/Telegram", "mp4;avi;wmv", "10", false)
}

func TestProcessAllVideos(t *testing.T) {
	ConvVideos2H265("/Users/zen/Downloads/bilibili", "mp4;avi;wmv", "10", false)

}
func TestDir(t *testing.T) {
	fp := "/Users/zen/Downloads/Telegram Desktop/水岛津实/33.mp4"
	ret := path.Dir(fp)
	t.Log(ret)
}
func TestOne(t *testing.T) {
	fp := "/Volumes/T7/slacking/Telegram/DOA/Christie x Kasumi x Marie Rose/我的影片13.mp4"
	threads := "10"
	ConvVideo2H265(fp, threads)
}
func TestOnces(t *testing.T) {
	threads := "10"
	fps := []string{
		"/Volumes/T7/slacking/Telegram/DOA/Christie x Marie Rose/我的影片18.mp4",
		"/Volumes/T7/slacking/Telegram/DOA/Kasumi/0000-0152.avi",
		"/Volumes/T7/slacking/Telegram/DOA/Leifang/Leifang_Splits.mp4",
		"/Volumes/T7/slacking/Telegram/DOA/Momiji/MomijiPoVold.mov",
	}
	for _, fp := range fps {
		ConvVideo2H265(fp, threads)
	}
}
func TestCutHead(t *testing.T) {
	src := "/Users/zen/Downloads/head"
	pattern := "mp4"
	CutHead(src, pattern, "00:00:06.500")
}

func TestGetENV(t *testing.T) {
	ConvVideos2H265("/Users/zen/Downloads/Telegram Desktop/shemale/solo", "mp4;mov", "4", false)
}
func TestProcessAllH265(t *testing.T) {
	root := "/Users/zen/Downloads/4x3/Telegram"
	pattern := "mp4"
	ProcessAllH265(root, pattern)
}
