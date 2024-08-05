package lib

type ByteIterator struct {
	data     []byte
	length   int
	index    int
	isString bool
}

const (
	BUF_OUT_OF_RANGE = 0
)

func NewIterator(data interface{}) *ByteIterator {
	var bytesData []byte
	var isString bool

	switch v := data.(type) {
	case []byte:
		bytesData = v
	case string:
		bytesData = []byte(v)
		isString = true

	default:
		panic("unsupported data type")
	}

	return &ByteIterator{
		data:     bytesData,
		length:   len(bytesData),
		index:    0,
		isString: isString,
	}
}

func (i *ByteIterator) Reset() {
	i.index = 0
}

func (i *ByteIterator) Len() int {
	return i.length
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
	if i.index == 0 {
		return i
	}
	i.index -= 1
	return i
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

func (i *ByteIterator) Get(pos int) byte {
	if pos >= i.length {
		return BUF_OUT_OF_RANGE
	}
	return i.data[pos]
}
