type MyHashSet struct {
	content []bool
}

func Constructor() MyHashSet {
	hashMapSlice := make([]bool, 1000001, 1000001)

	return MyHashSet{content: hashMapSlice}
}

func (this *MyHashSet) Add(key int) {
	this.content[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.content[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.content[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
