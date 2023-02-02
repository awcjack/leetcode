type MyHashSet struct {
	content []int
}

func Constructor() MyHashSet {
	hashMapSlice := make([]int, 0)

	return MyHashSet{content: hashMapSlice}
}

func (this *MyHashSet) Add(key int) {
	for _, v := range this.content {
		if v == key {
			return
		}
	}
	this.content = append(this.content, key)
}

func (this *MyHashSet) Remove(key int) {
	for i, v := range this.content {
		if v == key {
			copy(this.content[i:], this.content[i+1:])
			this.content = this.content[:len(this.content)-1]
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, v := range this.content {
		if v == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
