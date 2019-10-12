package vindexdata

import (
	"reflect"

	"github.com/pkg/errors"
)

// VIndexItem VIndexItem
type VIndexItem struct {
	Filename string `bi_length:"10000" bi_type:"string"`
	Date     uint64 `bi_length:"8" bi_type:"uint"`
	Storage  string `bi_length:"100" bi_type:"string"`
}

// NewVIndexItem NewVIndexItem
func NewVIndexItem(storage string, date uint64, filename string) VIndexItem {
	return VIndexItem{
		Storage:  storage,
		Filename: filename,
		Date:     date,
	}
}

// ToBinary ToBinary
func (item VIndexItem) ToBinary() []byte {
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
			tmp := uint2bytes(value.Uint(), 8)
			for i := range tmp {
				data[i] = tmp[i]
			}
		}

		for i := range data {
			ret[i+offset] = data[i]
		}

		offset += colLen
	}

	return ret
}

// NewBinaryIndexItemFromBinary NewBinaryIndexItemFromBinary
func NewBinaryIndexItemFromBinary(data []byte) (VIndexItem, error) {
	ret := VIndexItem{}
	pairs := structFields()
	totalLen := rowLength(pairs)
	if len(data) != totalLen {
		return ret, errors.Errorf("Invalid data. len(data)=%v totalLen=%v", len(data), totalLen)
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
			value := make([]byte, 0, len(colRawBytes))
			for i := range colRawBytes {
				if colRawBytes[i] == '\000' {
					break
				}

				value = append(value, colRawBytes[i])
			}

			rawValue.FieldByName(pair.Name).SetString(string(value))
		case "uint":
			value := bytes2uint(colRawBytes...)
			rawValue.FieldByName(pair.Name).SetUint(value)
		}

		offset += pair.Length
	}

	return ret, nil
}
