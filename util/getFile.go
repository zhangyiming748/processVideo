package util

import (
	"github.com/zhangyiming748/log"
	"os"
	"path"
	"strings"
)

type File struct {
	FullPath string // 文件的绝对路径
	Size     int64  // 文件大小
	FullName string // 文件名
	ExtName  string // 扩展名
}

/*
获取指定目录下符合条件的文件
*/
func GetMultiFiles(dir, pattern string) []File {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Debug.Printf("读取文件目录产生的错误:%v\n", err)
	}
	var aim []File
	if strings.Contains(pattern, ";") {
		exts := strings.Split(pattern, ";")
		for _, file := range files {
			if strings.HasPrefix(file.Name(), ".") {
				log.Info.Println("跳过隐藏文件", file.Name())
				continue
			}
			ext := path.Ext(file.Name())
			//log.Info.Printf("extname is %v\n", ext)
			for _, ex := range exts {
				if strings.Contains(ext, ex) {
					//aim = append(aim, file.Name())
					f := &File{
						FullPath: strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)),
						Size:     0,
						FullName: file.Name(),
						ExtName:  ext,
					}
					aim = append(aim, *f)
				}
			}
		}
	} else {
		for _, file := range files {
			if strings.HasPrefix(file.Name(), ".") {
				log.Info.Println("跳过隐藏文件", file.Name())
				continue
			}
			ext := path.Ext(file.Name())
			//log.Info.Printf("extname is %v\n", ext)
			if strings.Contains(ext, pattern) {
				//aim = append(aim, file.Name())
				f := &File{
					FullPath: strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)),
					Size:     0,
					FullName: file.Name(),
					ExtName:  ext,
				}
				aim = append(aim, *f)
			}
		}
	}
	log.Debug.Printf("有效的目标文件: %+v \n", aim)
	return aim
}
