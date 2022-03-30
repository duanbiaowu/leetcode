package leetcode

// Trie Tree Node
type Trie struct {
	IsWord   bool
	Children map[byte]*Trie
}

// Constructor Initialize your data structure here.
func Constructor() Trie {
	return Trie{false, make(map[byte]*Trie)}
}

// Insert a word into the trie.
func (t *Trie) Insert(word string) {
	parent := t
	for i := range word {
		if child, ok := parent.Children[word[i]]; ok {
			parent = child
		} else {
			newChild := &Trie{false, make(map[byte]*Trie)}
			parent.Children[word[i]] = newChild
			parent = newChild
		}
	}

	parent.IsWord = true
}

// Search Returns if the word is in the trie.
func (t *Trie) Search(word string) bool {
	parent := t
	for i := range word {
		if child, ok := parent.Children[word[i]]; ok {
			parent = child
		} else {
			return false
		}
	}
	return parent.IsWord
}

// StartsWith Returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	parent := t
	for i := range prefix {
		if child, ok := parent.Children[prefix[i]]; ok {
			parent = child
		} else {
			return false
		}
	}
	return true
}
