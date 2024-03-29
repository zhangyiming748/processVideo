package processVideo

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Convert2H265(in GetFileInfo.Info, threads string) {
	info := GetFileInfo.GetVideoFileInfo(in.FullPath)
	if info.Code == "HEVC" {
		slog.Info(fmt.Sprintf("跳过hevc文件:%v", in.FullPath))
		return
	}
	prefix := strings.Trim(in.FullPath, in.FullName)
	middle := "h265"
	os.MkdirAll(strings.Join([]string{prefix, middle}, ""), os.ModePerm)
	out := strings.Join([]string{prefix, middle, string(os.PathSeparator), in.FullName}, "")
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("出现错误", slog.String("输入文件", in.FullPath), slog.String("输出文件", out))
		}
	}()
	mp4 := strings.Join([]string{strings.Trim(out, in.ExtName), "mp4"}, "")
	slog.Debug("调试", slog.String("输入文件", in.FullPath), slog.String("输出文件", out))

	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libx265", "-threads", threads, "-tag:v", "hvc1", mp4)
	if runtime.GOOS == "darwin" {
		slog.Debug("匹配到苹果设备,使用硬件加速", slog.String("文档", "https://developer.apple.com/documentation/videotoolbox"))
		cmd = exec.Command("ffmpeg", "-threads", threads, "-hwaccel", "videotoolbox", "-i", in.FullPath, "-c:v", "libx265", "-threads", threads, "-tag:v", "hvc1", mp4)
	}
	// info := GetFileInfo.GetVideoFileInfo(in.FullPath)
	if info.Width > 1920 && info.Height > 1920 {
		slog.Warn("视频大于1080P需要使用其他程序先处理视频尺寸", slog.Any("原视频", in))
		return
	}
	slog.Debug("生成的命令", slog.String("command", fmt.Sprint(cmd)))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("cmd.StdoutPipe", slog.Any("产生错误", err))
		return
	}
	if err = cmd.Start(); err != nil {
		slog.Warn("cmd.Run", slog.Any("产生错误", err))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		fmt.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("命令执行中", slog.Any("产生错误", err))
		return
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err = os.RemoveAll(in.FullPath); err != nil {
		slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
	} else {
		slog.Info("删除成功", slog.Any("源文件", in.FullName))
	}
}
func ConvertOne(src, dst, threads string) {
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", src, "-c:v", "libx265", "-tag:v", "hvc1", "-threads", threads, dst)
	slog.Debug("生成的命令", slog.String("command", fmt.Sprint(cmd)))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("cmd.StdoutPipe", slog.Any("产生错误", err))
		return
	}
	if err = cmd.Start(); err != nil {
		slog.Warn("cmd.Run", slog.Any("产生错误", err))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err = stdout.Read(tmp)
		t := string(tmp)
		fmt.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("命令执行中", slog.Any("产生错误", err))
		return
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err = os.RemoveAll(src); err != nil {
		slog.Warn("删除失败", slog.Any("源文件", src), slog.Any("产生错误", err))
	} else {
		slog.Info("删除成功", slog.Any("源文件", src))
	}
}
