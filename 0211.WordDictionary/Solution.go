package leetcode

type WordDictionary struct {
	IsWord   bool
	Children map[byte]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{false, make(map[byte]*WordDictionary)}
}

func (this *WordDictionary) AddWord(word string) {
	parent := this

	for i := range word {
		if child, ok := parent.Children[word[i]]; ok {
			parent = child
		} else {
			child = &WordDictionary{false, make(map[byte]*WordDictionary)}
			parent.Children[word[i]] = child
			parent = child
		}
	}

	parent.IsWord = true
}

func (this *WordDictionary) Search(word string) bool {
	parent := this
	for i := range word {
		if word[i] == '.' {
			isMatched := false

			for _, child := range parent.Children {
				if child.Search(word[i+1:]) {
					isMatched = true
				}
			}

			return isMatched
		} else if _, ok := parent.Children[word[i]]; !ok {
			return false
		}

		parent = parent.Children[word[i]]
	}

	return len(parent.Children) == 0 || parent.IsWord
}
