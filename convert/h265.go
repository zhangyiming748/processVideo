package convert

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/replace"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Convert2H265(in GetFileInfo.Info, threads string) {
	info := GetFileInfo.GetVideoFileInfo(in.FullPath)
	if info.Code == "HEVC" {
		slog.Info(fmt.Sprintf("跳过hevc文件:%v\n", in.FullPath))
		return
	}
	prefix := strings.Trim(in.FullPath, in.FullName)
	middle := "h265"
	os.MkdirAll(strings.Join([]string{prefix, middle}, ""), os.ModePerm)
	out := strings.Join([]string{prefix, middle, "/", in.FullName}, "")
	defer func() {
		if err := recover(); err != nil {
			slog.Warn(fmt.Sprintf("出现错误的输入文件\"%v\"\n输出文件\"%v\"\n", in.FullPath, out))
		}
	}()
	mp4 := strings.Join([]string{strings.Trim(out, in.ExtName), "mp4"}, ".")
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libx265", "-threads", threads, "-tag:v", "hvc1", mp4)
	if runtime.GOOS == "darwin" {
		slog.Debug("匹配到苹果设备,使用硬件加速", "https://developer.apple.com/documentation/videotoolbox")
		cmd = exec.Command("ffmpeg", "-threads", threads, "-hwaccel", "videotoolbox", "-i", in.FullPath, "-c:v", "libx265", "-threads", threads, "-tag:v", "hvc1", mp4)
	}
	// info := GetFileInfo.GetVideoFileInfo(in.FullPath)
	if info.Width > 1920 && info.Height > 1920 {
		slog.Debug("视频大于1080P需要使用其他程序先处理视频尺寸", in)
		return
	}
	slog.Debug("生成的命令", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("cmd.StdoutPipe产生错误", err)
		return
	}
	if err = cmd.Start(); err != nil {
		slog.Warn("cmd.Run产生的错误", err)
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = replace.Replace(t)
		slog.Debug("ffmpeg", t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("命令执行中有错误产生", err)
		return
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err := os.RemoveAll(in.FullPath); err != nil {
		slog.Warn("删除源文件失败", err)
	} else {
		slog.Info("删除源文件", in.FullName)
	}
}
func ConvertOne(src, dst, threads string) {
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", src, "-c:v", "libx265", "-threads", threads, dst)
	slog.Debug("生成的命令", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("cmd.StdoutPipe产生错误", err)
		return
	}
	if err = cmd.Start(); err != nil {
		slog.Warn("cmd.Run产生的错误", err)
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = replace.Replace(t)
		slog.Debug("ffmpeg运行时的输出", t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("命令执行中有错误产生", err)
		return
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err := os.RemoveAll(src); err != nil {
		slog.Warn("删除源文件失败", err)
	} else {
		slog.Info("删除源文件", src)
	}
}
