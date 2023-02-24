package convert

import (
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/replace"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Convert2H265(in GetFileInfo.Info, threads string, fast bool) {
	info := GetFileInfo.GetVideoFileInfo(in.FullPath)
	if info.Code == "HEVC" {
		log.Debug.Printf("跳过hevc文件:%v\n", in.FullPath)
		return
	}
	prefix := strings.Trim(in.FullPath, in.FullName)
	middle := "h265"
	os.MkdirAll(strings.Join([]string{prefix, middle}, ""), os.ModePerm)
	out := strings.Join([]string{prefix, middle, "/", in.FullName}, "")

	defer func() {
		if err := recover(); err != nil {
			log.Warn.Printf("出现错误的输入文件\"%v\"\n输出文件\"%v\"\n", in.FullPath, out)
		}
	}()
	mp4 := strings.Join([]string{strings.Trim(out, in.ExtName), "mp4"}, ".")
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libx265", "-threads", threads, mp4)
	if runtime.GOOS == "darwin" && fast {
		cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "hevc_videotoolbox", "-pix_fmt", "yuv420p10le", "-threads", threads, mp4)
	}
	// info := GetFileInfo.GetVideoFileInfo(in.FullPath)
	if info.Width > 1920 && info.Height > 1920 {
		cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libx265", "-strict", "2", "-vf", "scale=-1:1080", "-threads", threads, mp4)
		if runtime.GOOS == "darwin" && fast {
			cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "hevc_videotoolbox", "-pix_fmt", "yuv420p10le", "-strict", "2", "-vf", "scale=-1:1080", "-threads", threads, mp4)
		}
	}
	log.Debug.Printf("生成的命令是:%s\n", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Warn.Panicf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		log.Warn.Panicf("cmd.Run产生的错误:%v\n", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		//log.Info.Printf("正在处理第 %d/%d 个文件: %s\n", index+1, total, file)
		t := string(tmp)
		t = replace.Replace(t)
		log.TTY.Printf("%v\b", t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Warn.Panicf("命令执行中有错误产生:%v\n", err)
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err := os.RemoveAll(in.FullPath); err != nil {
		log.Warn.Printf("删除源文件失败:%v\n", err)
	} else {
		log.Debug.Printf("删除源文件:%v\n", in.FullName)
	}
}
func ConvertOnce(src, dst, threads string) {
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", src, "-c:v", "libx265", "-threads", threads, dst)
	log.Debug.Printf("生成的命令是%v\n", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Warn.Panicf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		log.Warn.Panicf("cmd.Run产生的错误:%v\n", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		//log.Info.Printf("正在处理第 %d/%d 个文件: %s\n", index+1, total, file)
		t := string(tmp)
		t = replace.Replace(t)
		log.TTY.Printf("%v\b", t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Warn.Panicf("命令执行中有错误产生:%v\n", err)
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err := os.RemoveAll(src); err != nil {
		log.Warn.Printf("删除源文件失败:%v\n", err)
	} else {
		log.Debug.Printf("删除源文件:%v\n", src)
	}
}
