package vindexdata

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"golang.org/x/text/unicode/norm"
)

// Digest MD5
type Digest [16]byte

// MarshalJSON MarshalJSON
func (digest Digest) MarshalJSON() ([]byte, error) {
	return json.Marshal(digest.Hex())
}

// Hex Hex
func (digest Digest) Hex() string {
	return fmt.Sprintf("%x", digest)
}

// VIndexItem VIndexItem
type VIndexItem struct {
	Filename string `json:"filename" bi_length:"2000" bi_type:"string"`
	Date     uint64 `json:"date" bi_length:"8" bi_type:"uint"`
	Storage  string `json:"storage" bi_length:"100" bi_type:"string"`
	Digest   Digest `json:"digest" bi_length:"16" bi_type:"digest"`
}

// NewVIndexItem NewVIndexItem
func NewVIndexItem(storage string, date uint64, filename string) VIndexItem {
	normalized := norm.NFC.String(filename)

	return VIndexItem{
		Storage:  storage,
		Filename: normalized,
		Date:     date,
		Digest:   md5.Sum([]byte(normalized)),
	}
}

// HexDigest HexDigest
func (item VIndexItem) HexDigest() string {
	return item.Digest.Hex()
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
		case "digest":
			for i := 0; i < pair.Length; i++ {
				data[i] = byte(value.Index(i).Uint())
			}
		}

		for i := range data {
			ret[i+offset] = data[i]
		}

		offset += colLen
	}

	return ret
}

// MarshalJSON MarshalJSON
// func (item VIndexItem) MarshalJSON() ([]byte, error) {
// }

// NewBinaryIndexItemFromBinary NewBinaryIndexItemFromBinary
func NewBinaryIndexItemFromBinary(data []byte) (*VIndexItem, error) {
	ret := VIndexItem{}
	pairs := structFields()
	totalLen := int(rowLength(pairs))
	if len(data) != totalLen {
		return nil, errors.Errorf("Invalid data. len(data)=%v totalLen=%v", len(data), totalLen)
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
		case "digest":
			for i := range colRawBytes {
				rawValue.FieldByName(pair.Name).Index(i).SetUint(uint64(colRawBytes[i]))
			}
			// rawValue.FieldByName(pair.Name).SetBytes(colRawBytes)
			// rawValue.FieldByName(pair.Name).data = colRowBytes
		}

		offset += pair.Length
	}

	return &ret, nil
}
