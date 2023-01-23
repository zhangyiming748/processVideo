package processVideo

import (
	"fmt"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/processVideo/convert"
	"github.com/zhangyiming748/processVideo/util"
	"runtime"
	"strconv"
	"time"
)

const (
	MB = 1048576
)
const (
	SMALL  = 200 * MB
	MIDDLE = 800 * MB
	BIG    = 1200 * MB
	HUGE   = 2000 * MB
)

func ProcessVideos(dir, pattern string) {
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)

	thread := runtime.NumCPU() / 2
	threads := strconv.Itoa(thread)
	var files []util.File
	files = util.GetFileInfo(util.GetMultiFiles(dir, pattern))
	for _, file := range files {
		frame := util.DetectFrame(file)
		log.Debug.Printf("文件帧数约%d\n", frame)
		if file.Size < SMALL {
			convert.Convert2AV1(file, threads)
		} else if file.Size < MIDDLE {
			convert.Convert2VP9(file, threads)
		} else if file.Size < BIG {
			convert.Convert2VP8(file, threads)
		} else if file.Size < HUGE {
			convert.Convert2H265(file, threads)
		} else {
			convert.Convert2H265(file, threads)
		}
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	fmt.Printf("整个任务用时 %v 分\n", during)
}
