package processVideo

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/pretty"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"strings"
)

/*
h265视频添加标签
非h265视频生成报告
*/
func ProcessAllH265(root, pattern string) {
	folders := GetAllFolder.List(root)
	folders = append(folders, root)
	pretty.P(folders)
	report, err := os.OpenFile("report.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		slog.Error("打开文件失败", slog.Any("错误文本", err))
	}
	defer report.Close()
	for _, folder := range folders {
		files := GetFileInfo.GetAllVideoFileInfo(folder, pattern)
		for _, file := range files {
			reportOne(file, report)
		}
	}
}

func reportOne(fi GetFileInfo.Info, report *os.File) {
	if fi.Code == "HEVC" {
		if fi.VTag == "hvc1" {
			slog.Info("跳过正常文件", slog.String("文件名", fi.FullPath))
		} else {
			slog.Info("准备添加标签", slog.String("文件名", fi.FullPath))
			addTag(fi)
		}
	} else {
		slog.Info("记录非hevc的视频文件", slog.String("文件名", fi.FullPath))
		report.WriteString(fmt.Sprintf("%s\n", fi.FullPath))
	}
}
func processOne(fi GetFileInfo.Info, report *os.File) {
	if fi.Code == "HEVC" {
		if fi.VTag == "hvc1" {
			slog.Info("跳过正常文件", slog.String("文件名", fi.FullPath))
		} else {
			slog.Info("准备添加标签", slog.String("文件名", fi.FullPath))
			addTag(fi)
		}
	} else {
		slog.Info("记录非hevc的视频文件", slog.String("文件名", fi.FullPath))
		report.WriteString(fmt.Sprintf("%s\n", fi.FullPath))
	}
}

func addTag(fi GetFileInfo.Info) {
	prefix := strings.Trim(fi.FullPath, fi.FullName) // 带 /
	dst := strings.Join([]string{prefix, "tag"}, "")
	os.Mkdir(dst, 0777)
	target := strings.Join([]string{dst, fi.FullName}, string(os.PathSeparator))
	cmd := exec.Command("ffmpeg", "-i", fi.FullPath, "-c:v", "copy", "-c:a", "copy", "-c:s", "copy", "-tag:v", "hvc1", target)
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("combined", slog.Any("cmd.Run() failed", err), slog.String("out", string(out)))
		return
	}
	slog.Debug("combined", slog.String("out", string(out)))
	if err = os.RemoveAll(fi.FullPath); err != nil {
		slog.Warn("删除失败", slog.Any("源文件", fi.FullPath), slog.Any("错误", err))
	} else {
		slog.Info("删除成功", slog.Any("源文件", fi.FullName))
	}
}
func processTag(fi GetFileInfo.Info) {
	prefix := strings.Trim(fi.FullPath, fi.FullName) // 带 /
	dst := strings.Join([]string{prefix, "tag"}, "")
	os.Mkdir(dst, 0777)
	target := strings.Join([]string{dst, fi.FullName}, string(os.PathSeparator))
	cmd := exec.Command("ffmpeg", "-i", fi.FullPath, "-c:v", "libx265", "-c:a", "copy", "-c:s", "copy", "-tag:v", "hvc1", target)
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("combined", slog.Any("cmd.Run() failed", err), slog.String("out", string(out)))
		return
	}
	slog.Info("combined", slog.String("out", string(out)))
	if err = os.RemoveAll(fi.FullPath); err != nil {
		slog.Warn("删除失败", slog.Any("源文件", fi.FullPath), slog.Any("错误", err))
	} else {
		slog.Info("删除成功", slog.Any("源文件", fi.FullName))
	}
}
