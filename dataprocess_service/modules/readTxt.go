package modules

import (
	"bufio"
	"os"
)

func ReadFile(filePath string) ([]string, error) {
	//打开文件
	fi, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	buf := bufio.NewScanner(fi)
	// 循环读取
	var lineArr []string
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		lineArr = append(lineArr, line)
	}

	return lineArr, nil
}
