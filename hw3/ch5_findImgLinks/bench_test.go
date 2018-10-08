package main

import (
	"testing"
)

func BenchmarkHello(b *testing.B) {
	url := "https://goods.ru/"
	resp, err := getResponse(url)
	defer resp.Body.Close()

	doc := getDoc(err, resp)

	for c := 0; c <= 10000; c++ {
		_ = getLinksMap(doc, "img", "src")
	}

	//for key := range links {
	//	fmt.Println(key)
	//}
	//1        1 431 808 000 ns/op
	//		   1 728 597 000 ns/op
	//		   1 237 290 400 ns/op
}
