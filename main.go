package main

import "github.com/hashicorp/golang-lru/v2/lru"

func main() {
	newLru, err := lru.New(100000)
}
