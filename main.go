package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// オプションのフラグ定義
	hideHidden := flag.String("h", "y", "hide hidden file(y) / show hidden file(n)")
	maxDepth := flag.Int("t", 3, "depth of file structure")
	outputFile := flag.String("f", "", "output file name")

	// フラグのパース
	flag.Parse()

	// ディレクトリパスの取得（デフォルトはカレントディレクトリ）
	root := "."
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}

	// 出力先が指定されていない場合は標準出力
	var out *os.File
	if *outputFile != "" {
		// ファイルの存在確認
		if _, err := os.Stat(*outputFile); err == nil {
			// ファイルが存在している場合
			fmt.Printf("Warning:  '%s' is already exist.\n", *outputFile)
			fmt.Print("Overwrite？ (y/n): ")

			var userInput string
			fmt.Scanln(&userInput) // ユーザー入力を受け取る

			if userInput != "y" && userInput != "Y" {
				fmt.Println("Overwrite canceled")
				return
			}
		}
		var err error
		out, err = os.Create(*outputFile)
		if err != nil {
			fmt.Println("error: Failed to create file.:", err)
			return
		}
		defer out.Close() // 処理が終わったらファイルを閉じる
	} else {
		out = os.Stdout
	}

	// 指定されたパスが存在するか確認
	if _, err := os.Stat(root); os.IsNotExist(err) {
		fmt.Println("error: :The specified directory does not exist.", root)
		return
	}

	// ディレクトリ構造を表示
	printDirectoryStructure(root, "", "├── ", "│   ", "└── ", *hideHidden == "y", *maxDepth, 1, out)
}

func printDirectoryStructure(path, indent, branch, pipe, lastBranch string, hideHidden bool, maxDepth, currentDepth int, out *os.File) {
	// 深さ制限を超えたら終了
	if currentDepth > maxDepth {
		return
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintln(out, "error:", err)
		return
	}

	for i, entry := range entries {
		// 隠しファイルを非表示にする場合はスキップ
		if hideHidden && entry.Name()[0] == '.' {
			continue
		}

		isLast := i == len(entries)-1
		if isLast {
			fmt.Fprintln(out, indent+lastBranch+entry.Name())
		} else {
			fmt.Fprintln(out, indent+branch+entry.Name())
		}

		// サブディレクトリがあれば再帰的に処理
		if entry.IsDir() {
			newIndent := indent + pipe
			if isLast {
				newIndent = indent + "    "
			}
			printDirectoryStructure(filepath.Join(path, entry.Name()), newIndent, branch, pipe, lastBranch, hideHidden, maxDepth, currentDepth+1, out)
		}
	}
}
