package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 引数を取得
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	if _, err := os.Stat(root); os.IsNotExist(err) {
		fmt.Println("エラー: 指定されたディレクトリは存在しません:", root)
		return
	}

	printDirectoryStructure(root, "", "├── ", "│   ", "└── ")
}

func printDirectoryStructure(path, indent, branch, pipe, lastBranch string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	var filteredEntries []os.DirEntry
	for _, entry := range entries {
		if entry.Name()[0] != '.' {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	for i, entry := range filteredEntries {
		isLast := i == len(filteredEntries)-1
		if isLast {
			fmt.Println(indent + lastBranch + entry.Name())
		} else {
			fmt.Println(indent + branch + entry.Name())
		}

		if entry.IsDir() {
			newIndent := indent + pipe
			if isLast {
				newIndent = indent + "    "
			}
			printDirectoryStructure(filepath.Join(path, entry.Name()), newIndent, branch, pipe, lastBranch)
		}
	}
}
