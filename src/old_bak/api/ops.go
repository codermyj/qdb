package api

import (
	"bufio"
	"encoding/json"
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
func OpenAppend(path string) *os.File {
	file, err := os.OpenFile(path, os.O_APPEND, 0666)
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
	dataFile := OpenAppend("./data/data.txt")
	metaFile := OpenAppend("./data/meta.txt")
	defer dataFile.Close()
	defer metaFile.Close()
	dataFile.WriteString(data.toString())
	memData[data.Key] = KeyDir{data.Key, *pos, data.ValueSize}
	*pos += int64(data.length())
	meta, _ := json.Marshal(memData[data.Key])
	metaFile.WriteString(string(meta) + "\n")
}

// LoadData
// 启动时加载到内存的数据
func LoadData() map[string]KeyDir {
	file, err := os.Open("./data/meta.txt")
	if err != nil {
		fmt.Printf("读取文件出错, %v", err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	//var lines []string
	var dir KeyDir
	meta := make(map[string]KeyDir)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		err = json.Unmarshal([]byte(str), &dir)
		if err != nil {
			//fmt.Printf("解析数据出错, %v", err)
			continue
		}
		meta[dir.Key] = dir
		//if meta.Value != "HAS_BEEN_DELETED" {
		//	m[data.Key] = data.Value
		//} else {
		//	delete(m, data.Key)
		//}
	}
	//fmt.Printf("检索条件：%v\n", subStr)
	return meta
}

// GetData
// 查询一条数据/*
func GetData(key string, keyDir map[string]KeyDir) (string, bool) {
	file := OpenAppend("./data/data.txt")
	defer file.Close()
	dir, ok := keyDir[key]
	buf := make([]byte, dir.ValueSize)
	_, err := file.ReadAt(buf, dir.ValuePos)
	if err != nil {
		fmt.Println("读取数据失败！")
	}
	return string(buf), ok
}
