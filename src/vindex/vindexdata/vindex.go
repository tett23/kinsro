package vindexdata

import "github.com/pkg/errors"

// VIndex VIndex
type VIndex []VIndexItem

// ToBinary ToBinary
func (vindex VIndex) ToBinary() ([]byte, error) {
	pairs := structFields()
	totalRowLen := int(rowLength(pairs))
	ret := make([]byte, totalRowLen*len(vindex))

	for i := range vindex {
		bin := vindex[i].ToBinary()
		offset := totalRowLen * i

		for j := range bin {
			ret[offset+j] = bin[j]
		}
	}

	return ret, nil
}

// NewVIndexFromBinary NewVIndexFromBinary
func NewVIndexFromBinary(data []byte) (VIndex, error) {
	pairs := structFields()
	totalRowLen := int(rowLength(pairs))
	if len(data)%totalRowLen != 0 {
		return nil, errors.Errorf("Invalid binary")
	}

	rowsLen := len(data) / totalRowLen
	ret := make(VIndex, rowsLen)

	for i := 0; i < rowsLen; i++ {
		rawRow := data[i*totalRowLen : (i+1)*totalRowLen]
		vindexItem, err := NewBinaryIndexItemFromBinary(rawRow)
		if err != nil {
			return ret, err
		}

		ret[i] = *vindexItem
	}

	return ret, nil
}
