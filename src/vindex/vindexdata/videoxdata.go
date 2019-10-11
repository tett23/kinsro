package vindexdata

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// BinaryIndexItem BinaryIndexItem
type BinaryIndexItem struct {
	Filename string `bi_length:"100" bi_type:"string"`
	Date     uint64 `bi_length:"8" bi_type:"uint"`
}

// BinaryIndexItems BinaryIndexItems
type BinaryIndexItems []BinaryIndexItem

// BinaryIndexIterable BinaryIndexIterable
type BinaryIndexIterable interface {
	Next() *BinaryIndexItem
}

type nameLengthPair struct {
	Name   string
	Length int
	Type   string
}

// Next Next
func (item BinaryIndexItem) Next() *BinaryIndexItem {
	return &item
}

func structFields() []nameLengthPair {
	s := BinaryIndexItem{}
	t := reflect.TypeOf(s)
	lenFields := t.NumField()

	ret := make([]nameLengthPair, lenFields)
	for i := 0; i < lenFields; i++ {
		field := t.Field(i)
		fieldLength, err := strconv.Atoi(field.Tag.Get("bi_length"))
		if err != nil {
			panic(err)
		}
		typeName := field.Tag.Get("bi_type")

		ret[i] = nameLengthPair{
			Name:   field.Name,
			Length: fieldLength,
			Type:   typeName,
		}
	}

	return ret
}

func rowLength(pairs []nameLengthPair) int {
	var ret int
	for i := range pairs {
		ret += pairs[i].Length
	}

	return ret
}

// NewBinaryIndexItem NewBinaryIndexItem
func NewBinaryIndexItem(filename string, date uint64) BinaryIndexItem {
	return BinaryIndexItem{
		Filename: filename,
		Date:     date,
	}
}

// ToBinary ToBinary
func (item BinaryIndexItem) ToBinary() []byte {
	pairs := structFields()
	totalLen := rowLength(pairs)
	ret := make([]byte, totalLen)
	var offset int
	r := reflect.ValueOf(item)
	indirect := reflect.Indirect(r)

	for _, pair := range structFields() {
		value := indirect.FieldByName(pair.Name)
		colLen := pair.Length
		data := make([]byte, colLen)

		switch pair.Type {
		case "string":
			tmp := []byte(value.String())

			dataLen := len(tmp)
			cond := colLen - dataLen

			for i := 0; i < cond; i++ {
				tmp = append(tmp, '\000')
			}
			for i := range tmp {
				data[i] = tmp[i]
			}
		case "uint":
			tmp := uint2bytes(item.Date, 8)
			for i := range tmp {
				data[i] = tmp[i]
			}
		}

		for i := range data {
			ret[i+int(offset)] = data[i]
		}

		offset += colLen
	}

	return ret
}

// NewBinaryIndexItemFromBinary NewBinaryIndexItemFromBinary
func NewBinaryIndexItemFromBinary(data []byte) (BinaryIndexItem, error) {
	ret := BinaryIndexItem{}
	pairs := structFields()
	totalLen := rowLength(pairs)
	if len(pairs) > totalLen {
		return ret, errors.Errorf("Invalid data")
	}

	rawValue := reflect.ValueOf(&ret).Elem()
	offset := 0
	for _, pair := range pairs {
		colRawBytes := make([]byte, pair.Length)
		cond := offset + pair.Length
		for i := offset; i < cond; i++ {
			colRawBytes[i-offset] = data[i]
		}

		switch pair.Type {
		case "string":
			eolIndex := 0
			for eolIndex = range colRawBytes {
				if colRawBytes[eolIndex] == '\000' {
					break
				}
			}

			value := string(colRawBytes[:eolIndex])
			rawValue.FieldByName(pair.Name).SetString(value)
		case "uint":
			value := bytes2uint(colRawBytes...)
			rawValue.FieldByName(pair.Name).SetUint(value)
		}

		offset += pair.Length
	}

	return ret, nil
}

func str2bytes(str string) []byte {
	bytes := make([]byte, 8)
	for i, e := range strings.Fields(str) {
		b, _ := strconv.ParseUint(e, 16, 64)
		bytes[i] = byte(b)
	}
	return bytes
}

func bytes2str(bytes ...byte) string {
	strs := []string{}
	for _, b := range bytes {
		strs = append(strs, fmt.Sprintf("%02x", b))
	}
	return strings.Join(strs, " ")
}

func bytes2uint(bytes ...byte) uint64 {
	padding := make([]byte, 8-len(bytes))
	i := binary.BigEndian.Uint64(append(padding, bytes...))
	return i
}

func uint2bytes(i uint64, size int) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, i)
	return bytes[8-size : 8]
}

func bytes2int(bytes ...byte) int64 {
	if 0x7f < bytes[0] {
		mask := uint64(1<<uint(len(bytes)*8-1) - 1)

		bytes[0] &= 0x7f
		i := bytes2uint(bytes...)
		i = (^i + 1) & mask
		return int64(-i)

	}

	i := bytes2uint(bytes...)
	return int64(i)
}

func int2bytes(i int, size int) []byte {
	var ui uint64
	if 0 < i {
		ui = uint64(i)
	} else {
		ui = (^uint64(-i) + 1)
	}
	return uint2bytes(ui, size)
}
