package ops

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Create() {
	file, err := os.Create("./data/data.txt")
	if err != nil {
		fmt.Printf("创建文件失败，err: %v\n", err)
	}
	defer file.Close()
	fmt.Printf("创建成功\n")

}

func OpenAppend() *os.File {
	file, err := os.OpenFile("./data/data.txt", os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("文件打开失败, err: %v", err)
	}
	return file
}

func Rm() {
	fmt.Printf("删掉最后一行\n")
}

func Add(str string) {
	file := OpenAppend()
	defer file.Close()
	file.WriteString(str + "\n")
	fmt.Printf("追加一行, 内容是: %v\n", str)
}

func Find(subStr string) []string {
	file, err := os.Open("./data/data.txt")
	if err != nil {
		fmt.Printf("读取文件出错, %v", err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	var lines []string
	for fileScanner.Scan() {
		str := fileScanner.Text()
		if strings.Contains(str, subStr) {
			lines = append(lines, str)
		}
	}
	fmt.Printf("检索条件：%v\n", subStr)
	return lines
}
