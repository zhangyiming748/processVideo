package processVideo

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/replace"
	"github.com/zhangyiming748/voiceAlert"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"path"
	"strings"
)

/*
转换一个手动输入路径的视频为h265
*/
func ConvVideo2H265(fullpath, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	dst := strings.Join([]string{path.Dir(fullpath), "h265"}, string(os.PathSeparator))
	os.Mkdir(dst, 0777)
	filename := path.Base(fullpath)
	target := strings.Join([]string{dst, filename}, string(os.PathSeparator))
	ConvertOne(fullpath, target, threads)
}
func ConvVideos2H265(dir, pattern, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for i, file := range files {
		slog.Info(fmt.Sprintf("正在处理第 %d/%d 个视频", i+1, len(files)))
		Convert2H265(file, threads)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}

func ConvAllVideos2H265(root, pattern, threads string) {
	ConvVideos2H265(root, pattern, threads)
	folders := GetAllFolder.List(root)
	for _, folder := range folders {
		ConvVideos2H265(folder, pattern, threads)
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
	slog.Info("准备剪切", slog.Any("输入文件", in), slog.Any("输出文件", out))
	cmd := exec.Command("ffmpeg", "-ss", startAt, "-i", in, out)
	slog.Debug("生成命令", slog.Any("命令", cmd))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("cmd.StdoutPipe", slog.Any("错误", err))
		return
	}
	if err = cmd.Start(); err != nil {
		slog.Warn("cmd.Run", slog.Any("产生的错误", err))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = replace.Replace(t)
		fmt.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("命令执行中", slog.Any("错误", err))
		return
	}
}
