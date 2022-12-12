package kv

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"io"
	"os"
	"strconv"
)

const USIZE_LEN = strconv.IntSize / 8
const ENTRY_HEAD_LEN = USIZE_LEN*2 + 1
const STORAGE_FILE_PREFIX = "qfs"

const (
	PUT = 1
	DEL = 2
)

type Entry struct {
	KeyLen   uint   //键长度
	ValueLen uint   //值长度
	Kind     uint   //操作类型
	Key      string //键
	Value    string //值
}

func NewEntry(key string, value string, kind uint) *Entry {
	return &Entry{
		uint(len(key)),
		uint(len(value)),
		kind,
		key,
		value}
}

func (entry *Entry) size() uint {
	return ENTRY_HEAD_LEN + entry.KeyLen + entry.ValueLen
}

// 对数据进行二进制编码
func (entry *Entry) encode() []byte {

	buf := make([]byte, 0, entry.size())

	bufKeyLen, _ := convertor.ToBytes(entry.KeyLen)
	buf = append(buf, bufKeyLen...)

	bufValueLen, _ := convertor.ToBytes(entry.ValueLen)
	buf = append(buf, bufValueLen...)

	bufKind, _ := convertor.ToBytes(entry.Kind)
	buf = append(buf, bufKind[USIZE_LEN-1])

	buf = append(buf, []byte(entry.Key)...)

	buf = append(buf, []byte(entry.Value)...)

	fmt.Println("encode之后：buf", buf)

	return buf
}

// 对二进制数据进行Entry头部解码
func decode(buf []byte) *Entry {

	//测试
	fmt.Println("测试内容buf：", buf)
	KeyLen := uint(binary.BigEndian.Uint64(buf[0:USIZE_LEN]))
	//测试
	fmt.Println("测试内容KeyLen：", KeyLen)
	ValueLen := uint(binary.BigEndian.Uint64(buf[USIZE_LEN : USIZE_LEN*2]))
	//测试
	fmt.Println("测试内容ValueLen: ", ValueLen)
	Kind := uint(buf[USIZE_LEN*2])
	fmt.Println("测试内容Kind: ", Kind)
	//Key := string(buf[USIZE_LEN*2+1 : USIZE_LEN*2+1+KeyLen])
	//Value := string(buf[USIZE_LEN*2+1+KeyLen : USIZE_LEN*2+1+KeyLen+ValueLen])

	return &Entry{KeyLen, ValueLen, Kind, "", ""}
}

type Storage interface {
	get(key string) (string, error)
	put(key string, val string) error
	remove(key string) error
}

type SimplifiedBitcask struct {
	dataPathBuf    string
	reader         *os.File
	writer         *os.File
	index          map[string]uint64
	pendingCompact uint64
	pos            uint64
}

func OpenSimplifiedBitcask(path string) (*SimplifiedBitcask, error) {
	dataPathBuf := path + STORAGE_FILE_PREFIX + ".data"
	readOpen, err := os.Open(dataPathBuf)
	if err != nil {
		return &SimplifiedBitcask{}, err
	}
	reader := readOpen
	writeOpen, err := os.OpenFile(dataPathBuf, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return &SimplifiedBitcask{}, err
	}
	writer := writeOpen

	instance := SimplifiedBitcask{
		dataPathBuf,
		reader,
		writer,
		make(map[string]uint64),
		0,
		0}
	instance.loadIndex()
	return &instance, nil
}

func (s *SimplifiedBitcask) loadIndex() error {
	var offset uint64 = 0
	for {
		entry, err := s.readAt(offset)
		if err != nil {
			if err == io.EOF {
				s.pos = offset
				return nil
			} else {
				return err
			}
		} else {
			if entry.Kind == DEL {
				delete(s.index, entry.Key)
				continue
			}
			size := uint64(entry.size())
			s.index[entry.Key] = offset
			offset += size
		}
	}

}

func (s *SimplifiedBitcask) read(key string) (*Entry, error) {
	pos, ok := s.index[key]
	var err error
	if ok {
		return s.readAt(pos)
	}
	err = errors.New("the key not found: " + key)
	return nil, err
}

func (s *SimplifiedBitcask) readAt(offset uint64) (*Entry, error) {
	//解码Entry头
	bufHead := make([]byte, ENTRY_HEAD_LEN)
	_, err := s.reader.ReadAt(bufHead, int64(offset))
	if err != nil {
		return nil, err
	}
	entry := decode(bufHead)

	pos := offset + ENTRY_HEAD_LEN

	//解码Key
	bufKey := make([]byte, entry.KeyLen)
	_, err = s.reader.ReadAt(bufKey, int64(pos))
	if err != nil {
		return nil, err
	}
	entry.Key = string(bufKey)

	pos += uint64(entry.KeyLen)

	//解码Value
	bufValue := make([]byte, entry.ValueLen)
	_, err = s.reader.ReadAt(bufValue, int64(pos))
	if err != nil {
		return nil, err
	}
	entry.Value = string(bufValue)

	return entry, nil
}

func (s *SimplifiedBitcask) get(key string) (string, error) {
	entry, err := s.read(key)
	if err != nil {
		return "key not found", err
	}
	return entry.Value, err
}

func (s *SimplifiedBitcask) write(entry *Entry) error {
	key := entry.Key
	s.index[key] = s.pos
	buf := entry.encode()
	return s.writeData(buf)
}

func (s *SimplifiedBitcask) writeData(buf []byte) error {
	//测试
	fmt.Println("写入数据：", string(buf), buf)
	n, err := s.writer.Write(buf)
	if err != nil {
		return err
	}
	s.pos += uint64(n)
	return nil
}

func (s *SimplifiedBitcask) put(key string, val string) error {
	entry := NewEntry(key, val, PUT)
	return s.write(entry)
}

func (s *SimplifiedBitcask) remove(key string) error {
	_, ok := s.index[key]
	if ok {
		entry := NewEntry(key, "", DEL)
		s.write(entry)
		delete(s.index, key)
		return nil
	}
	return errors.New("key not found")
}
