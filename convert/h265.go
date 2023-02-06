package convert

import (
	"github.com/zhangyiming748/GetFileInfo"
	u "github.com/zhangyiming748/getInfo/util"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/replace"
	"os"
	"os/exec"
	"strings"
)

func Convert2H265(in GetFileInfo.Info, threads string) {
	prefix := strings.Trim(in.FullPath, in.FullName)
	middle := "h265"
	os.MkdirAll(strings.Join([]string{prefix, middle}, ""), os.ModePerm)
	out := strings.Join([]string{prefix, middle, "/", in.FullName}, "")
	mp4 := strings.Join([]string{strings.Trim(out, in.ExtName), "mp4"}, ".")
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libx265", "-threads", threads, mp4)
	if u.BiggerThenFHD(in.FullPath) {
		// 	//ffmpeg -i 1.mp4 -strict -2 -vf scale=-1:1080 4.mp4
		cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libx265", "-strict", "2", "-vf", "scale=-1:1080", "-threads", threads, mp4)
	}
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
		log.Info.Printf("%v\b", t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Debug.Panicf("命令执行中有错误产生:%v\n", err)
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err := os.RemoveAll(in.FullPath); err != nil {
		log.Debug.Printf("删除源文件失败:%v\n", err)
	} else {
		log.Debug.Printf("删除源文件:%v\n", in.FullName)
	}
}
