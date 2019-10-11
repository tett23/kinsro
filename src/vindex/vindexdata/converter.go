package vindexdata

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

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
