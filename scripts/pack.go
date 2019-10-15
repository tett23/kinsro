package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/pkg/errors"
)

const templateString = `package assets

var AssetIndex = {{.}}
`

func main() {
	vindexPath := os.Getenv("VINDEX_PATH")
	if vindexPath == "" {
		panic("")
	}

	file, err := os.Open(vindexPath)
	if err != nil {
		panic(errors.Wrapf(err, "os.Open failed. vindexPath=%v", vindexPath))
	}

	stat, err := file.Stat()
	if err != nil {
		panic(errors.Wrapf(err, "stat failed"))
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		panic(errors.Wrapf(err, "read failed. vindexPath=%v", vindexPath))
	}

	wf, err := os.Create("./src/js_main/assets/assets.go")
	if err != nil {
		panic(errors.Wrapf(err, "os.Create failed"))
	}
	defer wf.Close()

	tmpl, err := template.New("test").Parse(templateString)
	if err != nil {
		panic(errors.Wrapf(err, "Parse failed"))
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, fmt.Sprintf("%#v", data)); err != nil {
		panic(errors.Wrapf(err, "template failed"))
	}

	_, err = wf.Write(buf.Bytes())
	if err != nil {
		panic(errors.Wrapf(err, "Write failed"))
	}
}
