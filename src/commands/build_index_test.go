package commands

import (
	"fmt"
	"testing"
)

func TestBuildVIndex(t *testing.T) {
	// panic("aaaa")
	// t.Logf("%+v, \n", 1)
	items, err := BuildVIndex([]string{})
	fmt.Printf("%+v, \n", items)
	fmt.Println("aa", items)
	fmt.Println("bb", err)
}
