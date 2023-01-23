package util

import (
	"fmt"
	"os"
	"testing"
)

func TestGetMultiFiles(t *testing.T) {
	ret := GetMultiFiles("/Users/zen/Github/processVideo", "go;log")
	t.Log(ret)
}
func TestGetFileInfo(t *testing.T) {
	f, _ := os.Stat("/Users/zen/Github/processVideo/.gitignore")
	fmt.Printf("name: %v\n", f.Name())
	fmt.Printf("mode: %v\n", f.Mode())
	fmt.Printf("isdir: %v\n", f.IsDir())
	fmt.Printf("modtime: %v\n", f.ModTime())
	fmt.Printf("size: %v\n", f.Size())
	fmt.Printf("sys: %v\n", f.Sys())

}

func TestPreload(t *testing.T) {
	f := &File{
		FullPath: "/Users/zen/Github/processVideo/DB/004 - 【ヒマリとくま】一個サービスしとくよ.mp4",
		Size:     0,
		FullName: "004 - 【ヒマリとくま】一個サービスしとくよ.mp4",
		ExtName:  ".mp4",
	}
	DetectFrame(*f)
}
