package xxx

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

// ReadFile 读取文件中全部内容
func ReadFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

// ReadLines 读取文件，返回行数组
func ReadLines(filename string) ([]string, error) {
	return ReadLinesOffsetN(filename, 0, -1)
}

// ReadLinesOffsetN 读取文件，返回offset行之后n行，n<0表示不限制返回行数
func ReadLinesOffsetN(filename string, offset uint, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer func() {
		_ = f.Close()
	}()

	var ret []string

	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}

	return ret, nil
}
