package convert

import (
	"github.com.zhangyiming748/processVideo/util"
	"testing"
)

func TestConvH265(t *testing.T) {
	var f util.File
	f.FullPath = "/Users/zen/Downloads/男孩子打胶教程.MP4"
	f.FullName = "男孩子打胶教程.MP4"
	f.ExtName = ".MP4"
	Convert2H265(f, "2")
}
