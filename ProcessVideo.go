package processVideo

import (
	"github.com.zhangyiming748/processVideo/convert"
	"github.com.zhangyiming748/processVideo/util"
	"runtime"
	"strconv"
)

const (
	MB = 1048576
)

func ProcessVideos(dir, pattern string) {
	thread := runtime.NumCPU() / 2
	threads := strconv.Itoa(thread)
	var files []util.File
	files = util.GetFileInfo(util.GetMultiFiles(dir, pattern))
	for _, file := range files {
		if file.Size < 200*MB {
			convert.Convert2AV1(file, threads)
		} else {
			convert.Convert2H265(file, threads)
		}
	}
}
