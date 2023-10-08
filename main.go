package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

func tryCatch(task func(), errCallback func(R any)) {
	defer func() {
		if r := recover(); r != nil {
			errCallback(r)
			debug.PrintStack()
		}
	}()
	task()
}

func mainTask() {
	prefix, _ := os.Getwd()
	sourceFile, err := os.Open(os.Args[1])
	if err != nil {
		panic("无法打开源文件")
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(prefix + "/cn.srt")
	if err != nil {
		panic("无法创建新文件")
	}
	defer targetFile.Close()

	scanner := bufio.NewScanner(sourceFile)
	var block []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			block = append(block, line)
		} else {
			for i := 0; i < len(block)-1; i++ {
				targetFile.WriteString(block[i] + "\n")
			}
			targetFile.WriteString("\n")
			block = nil
		}

		if err := scanner.Err(); err != nil {
			panic("扫描出错")
		}
	}
}

func main() {
	tryCatch(mainTask, func(R any) {
		fmt.Println(R)
	})
}
