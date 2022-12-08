package entry

import (
	"encoding/binary"
	"github.com/duke-git/lancet/v2/convertor"
	"strconv"
)

type Entry struct {
	KeyLen   uint   //键长度
	ValueLen uint   //值长度
	Kind     uint   //操作类型
	Key      string //键
	Value    string //值
}

const (
	PUT = 1
	DEL = 2
)

const UsizeLen = strconv.IntSize / 8
const EntryHeadLen = UsizeLen*2 + 1

func (entry *Entry) size() uint {
	return EntryHeadLen + entry.KeyLen + entry.ValueLen
}

// 对数据进行二进制编码
func (entry *Entry) encode() []byte {

	buf := make([]byte, entry.size())

	bufKeyLen, _ := convertor.ToBytes(entry.KeyLen)
	buf = append(buf, bufKeyLen...)

	bufValueLen, _ := convertor.ToBytes(entry.ValueLen)
	buf = append(buf, bufValueLen...)

	bufKind, _ := convertor.ToBytes(entry.Kind)
	buf = append(buf, bufKind[UsizeLen-1])

	buf = append(buf, []byte(entry.Key)...)

	buf = append(buf, []byte(entry.Value)...)

	return buf
}

// 对二进制数据进行解码
func decode(buf []byte) *Entry {

	KeyLen := uint(binary.BigEndian.Uint64(buf[0:UsizeLen]))
	ValueLen := uint(binary.BigEndian.Uint64(buf[UsizeLen : UsizeLen*2]))
	Kind := uint(binary.BigEndian.Uint64(buf[UsizeLen*2 : UsizeLen*2+1]))
	Key := string(buf[UsizeLen*2+1 : UsizeLen*2+1+KeyLen])
	Value := string(buf[UsizeLen*2+1+KeyLen : UsizeLen*2+1+KeyLen+ValueLen])

	return &Entry{KeyLen, ValueLen, Kind, Key, Value}
}
