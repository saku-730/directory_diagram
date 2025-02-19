package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// オプションのフラグ定義
	hideHidden := flag.String("h", "y", "隠しファイルを表示しない (y) / 表示する (n)")
	maxDepth := flag.Int("t", 1, "表示する階層の深さ")
	outputFile := flag.String("f", "", "ファイルに出力する場合のファイル名")

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
			fmt.Printf("警告: ファイル '%s' はすでに存在します。\n", *outputFile)
			fmt.Print("上書きしますか？ (y/n): ")

			var userInput string
			fmt.Scanln(&userInput) // ユーザー入力を受け取る

			if userInput != "y" && userInput != "Y" {
				fmt.Println("ファイルの上書きはキャンセルされました。")
				return
			}
		}
		var err error
		out, err = os.Create(*outputFile)
		if err != nil {
			fmt.Println("エラー: ファイルの作成に失敗しました:", err)
			return
		}
		defer out.Close() // 処理が終わったらファイルを閉じる
	} else {
		out = os.Stdout
	}

	// 指定されたパスが存在するか確認
	if _, err := os.Stat(root); os.IsNotExist(err) {
		fmt.Println("エラー: 指定されたディレクトリは存在しません:", root)
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
		fmt.Fprintln(out, "エラー:", err)
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
