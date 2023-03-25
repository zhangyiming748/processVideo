package processVideo

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/processVideo/convert"
	"github.com/zhangyiming748/replace"
	"github.com/zhangyiming748/voiceAlert"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	MB = 1048576
)
const (
	SMALL  = 1 * MB
	MIDDLE = 20 * MB
	BIG    = 30 * MB
	HUGE   = 500 * MB
)

func init() {
	logLevel := os.Getenv("LEVEL")
	//var level slog.Level
	var opt slog.HandlerOptions
	switch logLevel {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Info("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}

	}
	file := "processVideo.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	//defer logf.Close() //如果不关闭可能造成内存泄露
	logger := slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
	slog.SetDefault(logger)
}

/*
转换一个手动输入路径的视频为h265
*/
func ProcessVideo(fullpath, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	dst := strings.Join([]string{path.Dir(fullpath), "h265"}, string(os.PathSeparator))
	os.Mkdir(dst, 0777)
	filename := path.Base(fullpath)
	target := strings.Join([]string{dst, filename}, string(os.PathSeparator))
	convert.ConvertOne(fullpath, target, threads)
}
func ProcessVideos(dir, pattern, threads string, focus bool) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for i, file := range files {
		slog.Info(fmt.Sprintf("正在处理第 %d/%d 个视频\n", i+1, len(files)))
		if focus {
			slog.Debug(fmt.Sprintln("异步获取视频帧数"))
			go GetFileInfo.CountFrame(&file)
		}
		convert.Convert2H265(file, threads)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}

func ProcessAllVideos(root, pattern, threads string, focus bool) {
	ProcessVideos(root, pattern, threads, focus)
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		ProcessVideos(folder, pattern, threads, focus)
	}
}

/*
批量去除视频片头
*/
func CutHead(src, pattern, startAt string) {
	infos := GetFileInfo.GetAllFileInfo(src, pattern)
	for i, info := range infos {
		slog.Info(fmt.Sprintf("正在处理第%d/%d个文件\n", i+1, len(infos)))
		dir := strings.Trim(info.FullPath, info.FullName)
		newBase := strings.Join([]string{dir, "afterHead"}, "")
		os.Mkdir(newBase, 0777)
		newFullPath := strings.Join([]string{newBase, info.FullName}, string(os.PathSeparator))
		doCut(info.FullPath, newFullPath, startAt)
	}
}
func doCut(in, out, startAt string) {
	slog.Info("输入文件", in, "输出文件", out)
	cmd := exec.Command("ffmpeg", "-ss", startAt, "-i", in, out)
	slog.Debug("生成命令", slog.Any("命令", cmd))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("cmd.StdoutPipe产生的错误", err)
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
}
