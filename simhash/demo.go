package main

import (
	"fmt"
	
	"github.com/yanyiwu/gosimhash"
)

func main() {
	hasher := gosimhash.New(
		"./dict/jieba.dict.utf8",
		"./dict/hmm_model.utf8",
		"./dict/idf.utf8",
		"./dict/stop_words.utf8")
	defer hasher.Free()
	sentence := "我来到北京清华大学"
	fingerprint := hasher.MakeSimhash(sentence, 5)
	fmt.Printf("%s simhash: %x\n", sentence, fingerprint)
}
