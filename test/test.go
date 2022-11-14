package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type TheData struct {
	Value     string //值
	Key       string //键
	ValueSize int    //值长度
	KeySize   int    //键长度
	Ts        int64  //时间戳
}

type KeyDir struct {
	// FileNo string  文件编号
	ValuePos  int64 // 值在文件中的位置
	ValueSize int   //值的长度
}

func (t *TheData) tsLen() int {
	return len(strconv.FormatInt(t.Ts, 10))
}

func (t *TheData) keySizeLen() int {
	return len(strconv.Itoa(t.KeySize))
}

func (t *TheData) valueSizeLen() int {
	return len(strconv.Itoa(t.ValueSize))
}

//	func (t *TheData) valuePos() int {
//		return t.tsLen() + t.keySizeLen() + t.valueSizeLen()
//	}
//
// 数据排列：value key ts
func (t *TheData) toString() string {
	return t.Value + t.Key + strconv.Itoa(t.ValueSize) + strconv.Itoa(t.KeySize) + strconv.FormatInt(t.Ts, 10)
}

func main() {
	var pos int64 = 0
	keydir := make(map[string]KeyDir)
	file, err := os.OpenFile("./data.txt", os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开文件失败：err: ", err)
		return
	}
	data1 := TheData{
		"b1",
		"a1",
		2,
		2,
		time.Now().UnixMilli(),
	}
	data2 := TheData{
		"b2",
		"a2",
		2,
		2,
		time.Now().UnixMilli(),
	}

	keydir[data1.Key] = KeyDir{pos, data1.ValueSize}
	file.WriteString(data1.toString())
	pos += int64(len(data1.toString()))
	keydir[data2.Key] = KeyDir{pos, data2.ValueSize}
	file.WriteString(data2.toString())
	pos += int64(len(data2.toString()))

	buf := make([]byte, data1.ValueSize)
	file.ReadAt(buf, keydir["a1"].ValuePos)
	s := string(buf)
	fmt.Println(s)

	buf = make([]byte, keydir["a2"].ValueSize)
	file.ReadAt(buf, keydir["a2"].ValuePos)
	s = string(buf)
	fmt.Println(s)
}
