package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"strings"
)

type Text struct {
	Raw           string
	WordList      []string
	WordMap       map[string]int
	WordCountList [][]interface{}
}

func NewText(raw string) *Text {
	return &Text{Raw: raw}
}

func (t *Text) Participle() {
	t.WordList = participle(t.Raw)
}

func participle(s string) []string {
	var words []string
	x := gojieba.NewJieba()
	defer x.Free()
	words = x.CutForSearch(s, false)
	return words
}

func (t *Text) CountWord() {
	nn := make(map[string]int, 0)
	tem := countWord(t.WordList)
	for k, v := range tem {
		if _, ok := nn[k]; ok {
			fmt.Println(k, " is exist")
		}
		nn[k] = v
	}
	t.WordMap = nn
	t.WordCountList = sortWordCount(t.WordMap)
}

func countWord(wordList []string) map[string]int {
	out := make(map[string]int, 0)
	for _, v := range wordList {
		//nv := strings.Trim(v, " ")
		if len(v) >= 6 {
			out[v]++
		}
	}
	return out
}

//从多到少排列
func sortWordCount(wordCountList map[string]int) [][]interface{} {
	return SortByValue(wordCountList)
}

func (t *Text) GetWordsList(maxLen, maxCount, minCount int) []string {
	return getWordList(t.WordCountList, maxLen, maxCount, minCount)
}

func (t *Text) GetWordsString(maxLen, maxCount, minCount int, sep string) string {
	outList := getWordList(t.WordCountList, maxLen, maxCount, minCount)
	if sep == ""{
		sep = "，"
	}
	return strings.
}

func getWordList(data [][]interface{}, maxLen, maxCount, minCount int) []string {
	out := make([]string, 0, 0)
	for _, v := range data {
		if minCount > 0 {
			if v[0].(int) < minCount {
				break
			}
		}
		if maxCount > 0 {
			if v[0].(int) > maxCount {
				continue
			}
		}
		out = append(out, v[1].(string))
		if maxLen > 0 {
			if len(out) >= maxLen {
				break
			}
		}
	}
	return out
}