package lib

import (
	"fmt"
)

type ByteIterator struct {
	lookup   map[byte]int
	data     []byte
	length   int
	index    int
	isString bool
}

const (
	BUF_OUT_OF_RANGE = 0
	IDX_OUT_OF_RANGE = -1
)

func NewByteIterator(data interface{}) *ByteIterator {
	var bytesData []byte
	var isString bool
	lookup := make(map[byte]int)

	switch v := data.(type) {
	case []byte:
		bytesData = v
	case string:
		bytesData = []byte(v)
		isString = true
	default:
		panic("unsupported data type")
	}

	for i := 0; i < len(bytesData); i++ {
		lookup[bytesData[i]] = i
	}

	return &ByteIterator{
		data:     bytesData,
		length:   len(bytesData),
		index:    0,
		lookup:   lookup,
		isString: isString,
	}
}

func (i *ByteIterator) Reset() {
	i.index = 0
}

func (i *ByteIterator) Len() int {
	return i.length
}

func (i *ByteIterator) Print() {
	fmt.Println(string(i.data))
}

func (i *ByteIterator) HasNext() bool {
	return i.index < i.length
}

// consume the current char and increment index by 1
func (i *ByteIterator) Next() *ByteIterator {
	return i.Advance(1)
}

// go to the previous char
func (i *ByteIterator) Prev() *ByteIterator {
	i.index -= 1
	return i
}

func (i *ByteIterator) GetIdx() int {
	if i.index < 0 || i.index >= i.length {
		return IDX_OUT_OF_RANGE
	}
	return i.index
}

// advance the internal index by skip steps. Returns the modified internal state
func (i *ByteIterator) Advance(skip int) *ByteIterator {
	i.index += skip
	return i
}

// Peek k steps/offset ahead
func (i *ByteIterator) Peek(args ...int) byte {
	var peekOffset = 0
	if len(args) > 0 {
		peekOffset = args[0]
	}
	return i.Get(i.index + peekOffset)
}

func (i *ByteIterator) End() byte {
	return i.Get(i.length - 1)
}

// This is a O(1) operation we could flip the complexities since regexes are not usually long strings
func (i *ByteIterator) Find(char byte) bool {
	_, ok := i.lookup[char]
	return ok
}

func (i *ByteIterator) Get(pos int) byte {
	if pos >= i.length {
		return BUF_OUT_OF_RANGE
	}
	return i.data[pos]
}
