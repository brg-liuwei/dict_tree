package dict_tree_test

import (
	"testing"

	"dict_tree"
)

func ExampleDictTree() {
	tree := dict_tree.NewDictTree()
	tree.Add("中华人民共和国", "大陆")
	tree.Add("中华民国", "台湾")
	tree.Add("中华人民", "五十六个民族")
	tree.Display(" ")
	// Output:
	// 中 华 人 民 #五十六个民族#
	//             共 和 国 #大陆#
	//       民 国 #台湾#
}

func BenchmarkRetrieval(b *testing.B) {
	tmap := map[string]string{
		"中华人民共和国": "大陆",
		"中华民国":    "台湾",
		"中华人民":    "五十六个民族",
		"人民币":     "RMB",
		"人民群众":    "people",
		"人民":      "人民",
		"人民法院":    "呵呵",
		"人民公仆":    "周永康",
		"人民英雄":    "任志强",
		"人民民主":    "中国梦",
		"人民民主专政":  "梦",
	}

	tree := dict_tree.NewDictTree()
	for k, v := range tmap {
		tree.Add(k, v)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Retrieval("雄关漫道真如铁，而今迈步从头越")
	}
}
