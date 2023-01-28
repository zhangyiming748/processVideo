package processVideo

import (
	"testing"
)

func TestMaster(t *testing.T) {
	//ret1 := util.GetMultiFiles("/Users/zen/Github/processVideo", "go;log")
	//t.Logf("%+v\n", ret1)
	//ret2 := util.GetFileInfo(ret1)
	//t.Logf("%+v\n", ret2)
	ProcessVideos("/Users/zen/Downloads/Telegram Desktop/Nier/2B", "mp4;avi;wmv", "2")
}
func TestListFiles(t *testing.T) {
	path := "/Users/zen/Github/processVideo"
	listFolders(path)
}

func TestProcessAllVideos(t *testing.T) {
	ProcessAllVideos("/Users/zen/Downloads/Telegram Desktop/Nier", "mp4", "2")
}
