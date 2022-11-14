package api

import (
	"fmt"
	"os"
)

// Create
// 创建一个新文件/*
func Create() {
	file, err := os.Create("./data/data.txt")
	if err != nil {
		fmt.Printf("创建文件失败，err: %v\n", err)
	}
	defer file.Close()
	fmt.Printf("创建成功\n")

}

// OpenAppend
// 以追加方式打开一个文件
// /*
func OpenAppend() *os.File {
	file, err := os.OpenFile("./data/data.txt", os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("文件打开失败, err: %v", err)
	}
	return file
}

// RmData
// 删除一条数据
///*
//func RmData(data TheData, memData map[string]KeyDir) {
//	data.Value = "HAS_BEEN_DELETED"
//	SetData(data, memData)
//	delete(memData, data.Key)
//}

// SetData
// 更新一条数据
// /*
func SetData(data TheData, memData map[string]KeyDir, pos *int64) {
	file := OpenAppend()
	defer file.Close()
	file.WriteString(data.toString())
	memData[data.Key] = KeyDir{*pos, data.ValueSize}
	*pos += int64(data.length())
}

// LoadData
// 启动时加载到内存的数据
// /*
//func LoadData() map[string]string {
//	file, err := os.Open("./data/data.txt")
//	if err != nil {
//		fmt.Printf("读取文件出错, %v", err)
//	}
//	defer file.Close()
//	fileScanner := bufio.NewScanner(file)
//	//var lines []string
//	var data Data
//	m := make(map[string]string)
//	for fileScanner.Scan() {
//		str := fileScanner.Text()
//		err = json.Unmarshal([]byte(str), &data)
//		if err != nil {
//			//fmt.Printf("解析数据出错, %v", err)
//			continue
//		}
//		if data.Value != "HAS_BEEN_DELETED" {
//			m[data.Key] = data.Value
//		} else {
//			delete(m, data.Key)
//		}
//	}
//	//fmt.Printf("检索条件：%v\n", subStr)
//	return m
//}

// GetData
// 查询一条数据/*
func GetData(subStr string, m map[string]KeyDir) (string, bool) {
	s, ok := m[subStr]
	return s, ok
}
