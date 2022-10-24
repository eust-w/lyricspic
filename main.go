package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"strings"
)

func main() {
	var s string
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()
	s = "长江大桥长江大桥长江大桥"
	words = x.CutForSearch(s, !use_hmm)
	fmt.Println(s)
	fmt.Println("搜索引擎模式:", strings.Join(words, "/"), len(words))
}
