package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"dict_tree"
)

func loadCityTag(file string) (map[string]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	m := make(map[string]string)
	reader := bufio.NewReader(f)
	for lineNo := 1; ; lineNo++ {
		line, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err == io.EOF {
				return m, nil
			}
			return nil, err
		}

		arr := bytes.SplitN(line, []byte(" "), 2)
		if len(arr) != 2 {
			fmt.Fprintln(os.Stderr, "line", lineNo, "format error: ", string(line))
			continue
		}
		m[string(bytes.TrimSpace(arr[0]))] = string(bytes.TrimSpace(arr[1]))
	}
}

func dealRecord(tree *dict_tree.DictTree, file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Fprintln(os.Stderr, "read record file error: ", err)
		}
		line = strings.TrimSpace(line)
		if tag, ok := tree.Retrieval(line); ok {
			fmt.Println(line, " ### ", tag)
		}
	}
}

func main() {
	tree := dict_tree.NewDictTree()
	cityTag, err := loadCityTag("data/cities")
	if err != nil {
		panic(err)
	}
	for city, tag := range cityTag {
		tree.Add(city, tag)
	}

	tree.Display(" ")

	dealRecord(tree, "data/records")
}
