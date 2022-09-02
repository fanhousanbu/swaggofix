package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

var swagPath string

func main() {
	flag.StringVar(&swagPath, "path", "./swagger/swagger.json", "swagger文件路径，默认./swagger/swagger.json")
	flag.Parse()

	c, err := os.ReadFile(swagPath)
	if err != nil {
		panic(err)
	}
	input := string(c)

	result := repl(&input)
	file, err := os.OpenFile(swagPath, os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if n, err := io.WriteString(file, *result); err != nil {
		panic(err)
	} else {
		fmt.Printf("修复成功，写入字节数：%d", n)
	}
}

// repl 替换
func repl(orig *string) *string {

	exp := regexp.MustCompile(`/:([\w+]*)`)

	result := exp.ReplaceAllString(*orig, "/{$1}")

	return &result
}
