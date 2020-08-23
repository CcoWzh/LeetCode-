package main

import "fmt"

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	n := len(word)
	for i := 0; i < n; i++ {
		cur := word[i] - 'a'
		if this.next[cur] == nil {
			this.next[cur] = new(Trie)
		}
		this = this.next[cur]
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	n := len(word)
	for i := 0; i < n; i++ {
		cur := word[i] - 'a'
		node = node.next[cur]
		if node == nil {
			return false
		}
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
//和 search 操作类似，只是不需要判断最后一个字符结点的isEnd，
//因为既然能匹配到最后一个字符，那后面一定有单词是以它为前缀的。
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	n := len(prefix)
	for i := 0; i < n; i++ {
		cur := prefix[i] - 'a'
		node = node.next[cur]
		if node == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // 返回 true
	fmt.Println(trie.Search("app"))     // 返回 false
	fmt.Println(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // 返回 true

}
