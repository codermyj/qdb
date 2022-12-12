package api

import "strconv"

type TheData struct {
	Value     string //值
	Key       string //键
	ValueSize int    //值长度
	KeySize   int    //键长度
	Ts        int64  //时间戳
}

type KeyDir struct {
	// FileNo string  文件编号(待实现)
	Key       string `json:"key"`
	ValuePos  int64  `json:"value_pos"`  // 值在文件中的位置
	ValueSize int    `json:"value_size"` //值的长度
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

func (t *TheData) toString() string {
	return t.Value + t.Key + strconv.Itoa(t.ValueSize) + strconv.Itoa(t.KeySize) + strconv.FormatInt(t.Ts, 10)
}

func (t *TheData) length() int {
	return len(t.toString())
}
