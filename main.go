package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Version = "0.1.0"
)

var (
	links4K  []string
	links8K  []string
	linksAll []string
)

// main 主程序入口
func main() {
	fmt.Println(`
┌────────────────────────────┐
│    ed2k Links Extractor    │
└────────────────────────────┘`)
	fmt.Printf(":: Version %s\n\n", Version)

	if len(os.Args) == 1 {
		fmt.Println("提示：需要指定目录")
		return
	}

	var err error
	var fi os.FileInfo
	var folders []os.DirEntry

	path := strings.Trim(os.Args[1], " ")
	fmt.Printf("> 读取目录：%s\n\n", path)

	if folders, err = os.ReadDir(path); err != nil {
		fmt.Printf("! 读取目录错误：%s\n", err)
		return
	}

	// 遍历目录
	for _, file := range folders {
		if fi, err = file.Info(); err != nil {
			fmt.Println("! 读取文件信息错误：", err)
			continue
		}

		if fi.Name()[0] == '.' {
			continue
		}

		extension := fi.Name()[len(fi.Name())-3:]
		if extension != "txt" {
			continue
		}

		if err = readLinks(path, fi.Name()); err != nil {
			fmt.Println("! 处理文件错误：", err)
			// 结束
			return
		}
	}

	// 写入文件
	if len(links4K) > 0 {
		if err = writeFile(links4K, fmt.Sprintf("%s/_ed2k_4k.txt", path)); err != nil {
			return
		}
	}
	if len(links8K) > 0 {
		if err = writeFile(links8K, fmt.Sprintf("%s/_ed2k_8k.txt", path)); err != nil {
			return
		}
	}
	if len(linksAll) > 0 {
		if err = writeFile(linksAll, fmt.Sprintf("%s/_ed2k_all.txt", path)); err != nil {
			return
		}
	}
}

// readList 处理文件内容
func readLinks(path, fileName string) (err error) {
	fmt.Printf("> 分析文件：%s/%s\n", path, fileName)

	// 打开文件
	var file *os.File
	if file, err = os.Open(fmt.Sprintf("%s/%s", path, fileName)); err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("! 关闭文件错误：", err)
		}
	}(file)

	// 创建一个 Scanner 对象，用于逐行读取文件内容
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	var video4K []string
	var video8K []string
	for scanner.Scan() {
		line := scanner.Text() // 获取当前行的文本内容

		// 跳过压缩包
		if strings.Contains(line, ".rar") || strings.Contains(line, ".zip") {
			continue
		}

		// 跳过非视频
		if !strings.Contains(line, "mp4") {
			continue
		}

		// 提取 ed2k 链接
		if strings.HasPrefix(line, "ed2k://") {
			if strings.Contains(strings.ToLower(line), "8k") {
				video8K = append(video8K, line)
				fmt.Println("8K:", line)
			} else {
				video4K = append(video4K, line)
				fmt.Println("4K:", line)
			}
		}
	}

	// 检查是否出现了读取文件的错误
	if err = scanner.Err(); err != nil {
		return
	}

	if len(video4K) > 0 {
		links4K = append(links4K, video4K...)
	}
	if len(video8K) > 0 {
		links8K = append(links8K, video8K...)
		linksAll = append(linksAll, video8K...)
	} else {
		linksAll = append(linksAll, video4K...)
	}

	fmt.Printf("\n:: %s/%s 处理完成\n\n", path, fileName)

	return
}

// writeFile 写入文件
func writeFile(lines []string, fileName string) (err error) {
	// 将数组的每一行拼接成一个字符串，每行之间使用换行符分隔
	content := strings.Join(lines, "\n")

	//fmt.Println(content)

	// 将拼接好的内容写入到文件中
	if err = os.WriteFile(fileName, []byte(content), 0666); err != nil {
		fmt.Println("! 写入文件时发生错误:", err)
		return
	}

	fmt.Printf(":: 内容已成功写入到文件：%s\n", fileName)

	return
}
