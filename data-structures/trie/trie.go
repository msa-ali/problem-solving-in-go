package trie

const ROOT = "*"

type TrieNode struct {
	Char     string
	Children map[string]*TrieNode
	isEnd    bool
}

func NewTrie() *TrieNode {
	return &TrieNode{
		Char:     ROOT, // Root Node
		Children: make(map[string]*TrieNode, 0),
		isEnd:    false,
	}
}

func (root *TrieNode) Insert(word string) {
	temp := root
	length := len([]rune(word))
	if length > 0 {
		for i, c := range word {
			c := string(c)
			node, ok := temp.Children[c]

			if !ok {
				temp.Children[c] = &TrieNode{
					Char:     c,
					Children: make(map[string]*TrieNode, 0),
				}
				if i == length-1 {
					temp.Children[c].isEnd = true
				}
				temp = temp.Children[c]
				continue
			}

			if i == length-1 {
				node.isEnd = true
			}
			temp = node
		}
	}
}

func (root *TrieNode) StartsWith(prefix string) bool {
	temp := root
	length := len([]rune(prefix))
	if length == 0 {
		return false
	}
	for i, c := range prefix {
		c := string(c)
		node, ok := temp.Children[c]
		if !ok {
			return false
		}
		if i == length-1 {
			return true
		}
		temp = node
	}
	return false
}

func (root *TrieNode) Search(word string) bool {
	temp := root
	length := len([]rune(word))
	if length == 0 {
		return false
	}
	for i, c := range word {
		c := string(c)
		node, ok := temp.Children[c]
		if !ok {
			return false
		}
		if i == length-1 {
			return node.isEnd
		}
		temp = node
	}
	return false
}

func (root *TrieNode) GetAllWordsStartsWith(prefix string) (result []string) {
	temp := root
	length := len([]rune(prefix))
	if length == 0 {
		return result
	}
	for i, c := range prefix {
		c := string(c)
		node, ok := temp.Children[c]
		if !ok {
			return result
		}
		if i == length-1 {

		}
		temp = node
	}
	return result
}

// func getAllWordsInTrie(node *TrieNode, prefix string) (result []string) {
// 	ch := make(chan string)
// 	done := make(chan struct {})

// 	if node == nil || len(node.Children) == 0 {
// 		return result
// 	}

// 	for c,t := range node.Children {
// 		go func(c string, node *TrieNode) {

// 		}(c,t)
// 	}

// 	for {
// 		select {
// 		case foundString := <-ch:
// 			result = append(result, foundString)
// 		case <-done:

// 	}
// }
