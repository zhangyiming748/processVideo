package convert

import (
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/processVideo/util"
	"github.com/zhangyiming748/replace"
	"os/exec"
	"testing"
)

func TestConvH265(t *testing.T) {
	var f util.File
	f.FullPath = "/Users/zen/Downloads/男孩子打胶教程.MP4"
	f.FullName = "男孩子打胶教程.MP4"
	f.ExtName = ".MP4"
	Convert2H265(f, "2")
}
func TestCircle(t *testing.T) {
	cmd := exec.Command("/opt/homebrew/bin/bash", "-c", "echo foo && echo bar")
	log.Debug.Printf("生成的命令是:%s\n", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Debug.Panicf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		log.Debug.Panicf("cmd.Run产生的错误:%v\n", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		//log.Info.Printf("正在处理第 %d/%d 个文件: %s\n", index+1, total, file)
		t := string(tmp)
		t = replace.Replace(t)
		log.Info.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Debug.Panicf("命令执行中有错误产生:%v\n", err)
	}
}
