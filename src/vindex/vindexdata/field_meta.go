package vindexdata

import (
	"reflect"
	"strconv"
)

type nameLengthPair struct {
	Name   string
	Length int
	Type   string
}

func structFields() []nameLengthPair {
	s := VIndexItem{}
	reflType := reflect.TypeOf(s)
	lenFields := reflType.NumField()
	ret := make([]nameLengthPair, lenFields)

	for i := 0; i < lenFields; i++ {
		field := reflType.Field(i)
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

func rowLength(pairs []nameLengthPair) int64 {
	var ret int
	for i := range pairs {
		ret += pairs[i].Length
	}

	return int64(ret)
}

// RowLength RowLength
func RowLength() int64 {
	return rowLength(structFields())
}
