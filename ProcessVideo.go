package processVideo

import (
	"fmt"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/processVideo/convert"
	"github.com/zhangyiming748/processVideo/util"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	MB = 1048576
)
const (
	SMALL  = 50 * MB
	MIDDLE = 20 * MB
	BIG    = 30 * MB
	HUGE   = 500 * MB
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
		} else {
			convert.Convert2H265(file, threads)
		}
		//if file.Size < SMALL {
		//	convert.Convert2AV1(file, threads)
		//} else if file.Size < MIDDLE {
		//	// convert.Convert2VP9(file, threads)
		//	convert.Convert2H265(file, threads)
		//} else if file.Size < BIG {
		//	// convert.Convert2VP8(file, threads)
		//	convert.Convert2H265(file, threads)
		//} else if file.Size < HUGE {
		//	convert.Convert2H265(file, threads)
		//} else {
		//	convert.Convert2H265(file, threads)
		//}
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	log.Debug.Printf("整个任务用时 %v 分\n", during)
}

func ProcessAllVideos(root, pattern string) {
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)

	thread := runtime.NumCPU() / 2
	threads := strconv.Itoa(thread)
	var files []util.File
	folders := listFolders(root)
	for _, src := range folders {
		files = util.GetFileInfo(util.GetMultiFiles(src, pattern))
		for _, file := range files {
			frame := util.DetectFrame(file)
			log.Debug.Printf("文件帧数约%d\n", frame)
			if file.Size < SMALL {
				convert.Convert2AV1(file, threads)
			} else {
				convert.Convert2H265(file, threads)
			}
			//if file.Size < SMALL {
			//	convert.Convert2AV1(file, threads)
			//} else if file.Size < MIDDLE {
			//	convert.Convert2VP9(file, threads)
			//} else if file.Size < BIG {
			//	convert.Convert2VP8(file, threads)
			//} else if file.Size < HUGE {
			//	convert.Convert2H265(file, threads)
			//} else {
			//	convert.Convert2H265(file, threads)
			//}
		}
	}

	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	log.Debug.Printf("整个任务用时 %v 分\n", during)
}

func listFolders(dirname string) []string {
	fileInfos, _ := os.ReadDir(dirname)
	var folders []string
	for _, fi := range fileInfos {
		filename := strings.Join([]string{dirname, fi.Name()}, "/") //拼写当前文件夹中所有的文件地址
		//fmt.Println(filename)                                       //打印文件地址
		if fi.IsDir() { //判断是否是文件夹 如果是继续调用把自己的地址作为参数继续调用
			if strings.Contains(fi.Name(), "h265") {
				continue
			}
			fmt.Printf("获取到的文件夹:%v\n", filename)
			folders = append(folders, filename)
			listFolders(filename) //递归调用
		}
	}
	return folders
}
