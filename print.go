package jittoku

import "github.com/davecgh/go-spew/spew"

// PrintNyan 標準エラー出力に猫を表示する
func PrintNyan() {
	println("########## 😺 ##########")
}

// Dump spew.Dumpを呼ぶ。標準出力にダンプ。
func Dump(a ...interface{}) {
	spew.Dump(a...)
}