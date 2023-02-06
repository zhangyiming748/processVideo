package processVideo

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/getInfo"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/processVideo/convert"
	"github.com/zhangyiming748/voiceAlert"
	"time"
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

func ProcessVideos(dir, pattern, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Voice(voiceAlert.FAILED)
		}
	}()
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)

	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for i, file := range files {
		log.Debug.Printf("符合条件的第%d个文件:%+v\n", i+1, file)
	}
	for _, file := range files {
		//frame := util.DetectFrame(file)
		info := GetFileInfo.GetVideoFileInfo(file.FullPath)
		if info.Code == "HEVC" {
			continue
		}
		go getInfo.GetVideoFrame(file.FullPath)
		convert.Convert2H265(file, threads)
		voiceAlert.Voice(voiceAlert.SUCCESS)
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	voiceAlert.Voice(voiceAlert.COMPLETE)
	log.Debug.Printf("整个任务用时 %v 分\n", during)
}

func ProcessAllVideos(root, pattern, threads string) {
	ProcessVideos(root, pattern, threads)
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		ProcessVideos(folder, pattern, threads)
	}
}
