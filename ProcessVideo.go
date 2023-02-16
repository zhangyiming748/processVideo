package processVideo

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
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

func ProcessVideos(dir, pattern, threads string, focus bool) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "文件转换失败")
		}
	}()
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)

	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for i, file := range files {
		log.Debug.Printf("符合条件的第%d个文件:%+v\n", i+1, file)
	}
	for i, file := range files {
		//frame := util.DetectFrame(file)
		log.Debug.Printf("正在处理第 %d/%d 个视频\n", i+1, len(files))
		if focus {
			go GetFileInfo.CountFrame(&file)
		}
		convert.Convert2H265(file, threads)
		go voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "单个文件转换完成")
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	go voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "单个目录下文件全部转换完成")
	log.Debug.Printf("整个任务用时 %v 分\n", during)
}

func ProcessAllVideos(root, pattern, threads string, focus bool) {
	ProcessVideos(root, pattern, threads, focus)
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		ProcessVideos(folder, pattern, threads, focus)
	}
}
