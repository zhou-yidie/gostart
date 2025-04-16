package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// traverseDirectory 递归遍历目录
func traverseDirectory(path string, level int) error {
	// 读取目录内容
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	// 遍历每个条目
	for _, entry := range entries {
		// 根据当前层级生成缩进
		indent := ""
		for i := 0; i < level; i++ {
			indent += "  "
		}
		// 打印条目名称
		fmt.Printf("%s%s\n", indent, entry.Name())
		// 如果条目是目录，递归调用 traverseDirectory
		if entry.IsDir() {
			subPath := filepath.Join(path, entry.Name())
			err := traverseDirectory(subPath, level+1)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	// 请将该路径替换为你要遍历的目录路径
	rootPath := "/root/code/gostart/ch12"
	err := traverseDirectory(rootPath, 0)
	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
	}
}
