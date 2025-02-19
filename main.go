package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "."
	// 指定されたパスが存在するか確認
	if _, err := os.Stat(root); os.IsNotExist(err) {
		fmt.Println("エラー: 指定されたディレクトリは存在しません:", root)
		return
	}

	// ディレクトリ構造を表示
	printDirectoryStructure(root, "", "├── ", "│   ", "└── ")
}

func printDirectoryStructure(path, indent, branch, pipe, lastBranch string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("エラー:", err)
	}

	for i, entry := range entries {
		if entry.Name()[0] == '.' {
			continue
		}

		isLast := i == len(entries)-1
		if isLast {
			fmt.Println(indent + lastBranch + entry.Name())
		} else {
			fmt.Println(indent + branch + entry.Name())
		}

		// サブディレクトリがある場合、再帰的に処理
		if entry.IsDir() {
			newIndent := indent + pipe
			if isLast {
				newIndent = indent + "    "
			}
			printDirectoryStructure(filepath.Join(path, entry.Name()), newIndent, branch, pipe, lastBranch)
		}
	}
}
