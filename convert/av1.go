package convert

import (
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/replace"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"strings"
)

func Convert2AV1(in GetFileInfo.Info, threads string) {
	prefix := strings.Trim(in.FullPath, in.FullName)
	middle := "av1"
	os.MkdirAll(strings.Join([]string{prefix, middle}, ""), os.ModePerm)
	out := strings.Join([]string{prefix, middle, "/", in.FullName}, "")
	mkv := strings.Join([]string{strings.Trim(out, in.ExtName), "mkv"}, ".")

	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libaom-av1", "-crf", "30", "-threads", threads, mkv)
	slog.Warn("生成的命令", cmd)
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
		slog.Debug("ffmpeg程序输出", t)
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
		log.Debug.Printf("删除源文件失败:%v\n", err)
		slog.Warn("删除源文件失败", err, in.FullPath)
	} else {
		log.Debug.Printf("删除源文件:%v\n", in.FullName)
		slog.Info("删除源文件", in.FullPath)
	}
}
