package processVideo

import "github.com.zhangyiming748/processVideo/util"

func ProcessVideos(dir, pattern, threads string) {

	util.GetFileInfo(util.GetMultiFiles("/Users/zen/Github/processVideo", "go;log"))
}
