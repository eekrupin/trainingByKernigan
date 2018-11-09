package main

import (
	"strconv"
	"testing"
)

func init() {
	for i := 1; i <= 10000; i++ {
		cache.set(strconv.Itoa(i), strconv.Itoa(i))
	}
}

func BenchmarkSet(b *testing.B) {

	for i := 1; i <= b.N; i++ {
		go func() {
			cache.set(strconv.Itoa(i), strconv.Itoa(i))
			cache.set(strconv.Itoa(i+1), strconv.Itoa(i+1))
			cache.set(strconv.Itoa(i+2), strconv.Itoa(i+2))
		}()
	}

}

func BenchmarkGet(b *testing.B) {

	for c := 1; c <= b.N; c++ {

		for i := 1; i <= 1000; i++ {
			go func() {
				cache.set(strconv.Itoa(i), strconv.Itoa(i))
			}()
		}
	}

}

func BenchmarkDelete(b *testing.B) {

	for c := 1; c <= b.N; c++ {

		for i := 1; i <= 1000; i++ {
			go func() {
				cache.delete(strconv.Itoa(i))
			}()
		}
	}

}

func BenchmarkExists(b *testing.B) {

	for c := 1; c <= b.N; c++ {

		for i := 1; i <= 1000; i++ {
			go func() {
				cache.exists(strconv.Itoa(i))
			}()
		}
	}

}
