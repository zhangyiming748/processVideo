package processVideo

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/processVideo/convert"
	"github.com/zhangyiming748/voiceAlert"
	"os"
	"path"
	"strings"
)

const (
	MB = 1048576
)
const (
	SMALL  = 1 * MB
	MIDDLE = 20 * MB
	BIG    = 30 * MB
	HUGE   = 500 * MB
)

/*
转换一个手动输入路径的视频为h265
*/
func ProcessVideo(fullpath, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	dst := strings.Join([]string{path.Dir(fullpath), "h265"}, string(os.PathSeparator))
	os.Mkdir(dst, 0777)
	filename := path.Base(fullpath)
	target := strings.Join([]string{dst, filename}, string(os.PathSeparator))
	log.Debug.Printf("src = %v\t dst = %v\n", fullpath, target)
	convert.ConvertOne(fullpath, target, threads)
}
func ProcessVideos(dir, pattern, threads string, focus bool) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for i, file := range files {
		log.Debug.Printf("符合条件的第%d个文件:%+v\n", i+1, file)
	}
	for i, file := range files {
		log.Debug.Printf("正在处理第 %d/%d 个视频\n", i+1, len(files))
		if focus {
			go GetFileInfo.CountFrame(&file)
		}
		convert.Convert2H265(file, threads)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}

func ProcessAllVideos(root, pattern, threads string, focus bool) {
	ProcessVideos(root, pattern, threads, focus)
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		ProcessVideos(folder, pattern, threads, focus)
	}
}
