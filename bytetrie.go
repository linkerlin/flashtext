package flashtext

type bytetrie struct {
	key  byte
	next [256]*bytetrie
	word string
}

func NewByteTrie(b byte) *bytetrie {
	return &bytetrie{
		key:  b,
	}
}

func (t *bytetrie) addKeyword(keyword string) {
	var pos = t
	for _, b := range []byte( keyword ){
		pos = pos.getOrSet(b)
	}
	pos.word = keyword
}

func (t *bytetrie) removeKeyword(keyword string) {
	var pos = t
	for _, b := range []byte( keyword ) {
		pos = pos.getOrSet(b)
	}
	//fake delete
	pos.word = ""
}

func (t *bytetrie) getOrSet(b byte) *bytetrie {
	if t.next[b] != nil{
		return t.next[b]
	}
	next := NewByteTrie(b)
	t.next[b] = next
	return next
}

func (t *bytetrie) exists(keyword string) bool {
	var keytrie = t
	for _, b := range []byte( keyword ) {
		if keytrie.next[b]!=nil {
			keytrie = keytrie.next[b]
		} else {
			return false
		}
	}
	return keytrie.word != ""
}
