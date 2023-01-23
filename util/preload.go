package util

import (
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/replace"
	"os/exec"
	"strconv"
)

func Preload() {

}

/*
输出文件全部帧数
"ffprobe", "-v" ,"error" ,"-count_frames", "-select_streams", "v:0" ,"-show_entries", "stream=nb_read_frames", "-of", "default=nokey=1:noprint_wrappers=1" ,"input.mp4"
*/
func DetectFrame(f File) int {

	cmd := exec.Command("ffprobe", "-v", "error", "-count_frames", "-select_streams", "v:0", "-show_entries", "stream=nb_read_frames", "-of", "default=nokey=1:noprint_wrappers=1", f.FullPath)
	log.Debug.Printf("生成的命令是:%s\n", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Warn.Panicf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		log.Warn.Panicf("cmd.Run产生的错误:%v\n", err)
	}

	tmp := make([]byte, 1024)
	stdout.Read(tmp)
	t := string(tmp)
	t = replace.Replace(t)
	if atoi, err := strconv.Atoi(t); err == nil {
		return atoi
	}
	log.Warn.Println("读取文件帧数出错")
	return 0
}
