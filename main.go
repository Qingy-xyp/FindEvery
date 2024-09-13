package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 搜索目录中的文件
func searchFiles(directory string, extensions []string) []string {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, ext := range extensions {
			if strings.HasSuffix(info.Name(), ext) {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("指定路径错误:", err)
	}
	return files
}

// 搜索文件内容
func searchContent(filePath, content string) []string {
	var matchingLines []string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("[-] 打开文件%s错误\n", filePath)
		return matchingLines
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, content) {
			lineOutput := fmt.Sprintf("[+] 文件第%d行: %s", lineNum, line)
			matchingLines = append(matchingLines, lineOutput)
			// 同时将结果输出到控制台
			fmt.Println(lineOutput)
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("[-] 读取文件错误%s: %v\n", filePath, err)
	}
	return matchingLines
}

// 将匹配结果写入文件
func writeToFile(outputFile, filePath string, matchingLines []string) {
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("[-] Error writing to file %s\n", outputFile)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "[!] 文件地址: %s\n", filePath)
	for _, line := range matchingLines {
		fmt.Fprintln(writer, line)
	}
	fmt.Fprintln(writer)
	writer.Flush()
}

func main() {
	name := flag.String("n", "", "指定需要查找的文件类型，如：.txt,.text,.ini,.yaml,.yml,.php,.jsp,.java,.xml,.sql,.properties")
	content := flag.String("c", "", "指定要查找的字段(不支持多个字段)，如：password=,jdbc:,user=,key=,ssh-,ldap:,mysqli_connect,sk-")
	outputFile := flag.String("o", "findout.txt", "Specify output file")
	directory := flag.String("d", "./", "Target directory")
	flag.Parse()

	if *name == "" || *content == "" {
		fmt.Println("[-] 请输入参数\n-n 指定需要查找的文件类型，如：.txt,text,.ini,.yaml,.yml,.php,.jsp,.java,.xml,.sql,.properties\n-c 指定要查找的字段(不支持多个字段)，如：password=,jdbc:,user=,key=,ssh-,ldap:,mysqli_connect,sk-\n示例：go run main.go -n .txt,.text,.ini,.yaml,.yml,.php,.jsp,.java,.xml,.sql,.properties -c \"password=\" -d /Users/admin/Desktop/")
		return
	}

	extensions := strings.Split(*name, ",")
	fmt.Println("[!] 快速启动中...")

	files := searchFiles(*directory, extensions)
	for _, filePath := range files {
		matchingLines := searchContent(filePath, *content)
		if len(matchingLines) > 0 {
			writeToFile(*outputFile, filePath, matchingLines)
		}
	}

	fmt.Println("[!] 详细结果保存至resout.txt")
}
