package util

import (
	"os"
)

/*
获取文件的元数据
*/
func GetFileInfo(files []File) []File {
	//fullname := strings.Join([]string{dir, fname}, string(os.PathSeparator))
	//mate, err := os.Stat(fullname)
	//if err != nil {
	//	log.Debug.Printf("读取文件%s元数据出错:%v\n", fullname, err)
	//}
	//size := mate.Size()
	//fmt.Println(size)
	//if size < 209715200 {
	//	convert.Convert2AV1()
	//} else {
	//	convert.Convert2H265()
	//}
	var fullInfos []File
	for _, file := range files {
		mate, _ := os.Stat(file.FullPath)
		mate.Size()
		fullInfo := &File{
			FullPath: file.FullPath,
			Size:     mate.Size(),
			FullName: file.FullName,
			ExtName:  file.ExtName,
		}
		fullInfos = append(fullInfos, *fullInfo)
	}
	return fullInfos
}
