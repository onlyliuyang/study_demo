package main

import "fmt"

type Trie struct {
	IsLeaf   bool
	Children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

// insert a word into the trie
func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}

	if this == nil {
		t := Constructor()
		this = &t
	}

	key := getKey(word[0])
	if this.Children[key] == nil {
		t := Constructor()
		this.Children[key] = &t
	}

	if len(word) == 1 {
		this.Children[key].IsLeaf = true
	} else {
		this.Children[key].Insert(word[1:])
	}
}

//returns if the word is in the trie
func (this *Trie) Search(word string) bool {
	if word == "" {
		return false
	}

	key := getKey(word[0])
	if this.Children[key] == nil {
		return false
	}

	if len(word) == 1 {
		return this.Children[key].IsLeaf
	}
	return this.Children[key].Search(word[1:])
}

//returns if there is any word in the trie that starts with the given prefix
func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return false
	}

	key := getKey(prefix[0])
	if this.Children[key] == nil {
		return false
	}

	if len(prefix) == 1 {
		return true
	}

	return this.Children[key].StartsWith(prefix[1:])
}

func getKey(v byte) int {
	if v >= 'A' && v <= 'Z' {
		return int(v - 'A')
	}

	if v >= 'a' && v <= 'z' {
		return int(v - 'a')
	}
	return 0
}

func main() {
	obj := Constructor()
	obj.Insert("/hello/world")
	fmt.Println(obj.Search("/hello/world"))
	fmt.Println(obj.StartsWith("/hello"))
}
