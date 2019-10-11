package vindex

import "github.com/tett23/kinsro/src/vindex/writer"

// VIndexItem VIndexItem
type VIndexItem struct {
	Filename string
}

// VIndex VIndex
type VIndex []VIndexItem

// Append Append
func (vindex *VIndex) Append(item VIndexItem) {
	writer.Append(writer.BinaryIndexItem{
		Filename: item.Filename,
	})
}
