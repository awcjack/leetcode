type MyHashMap struct {
	content []int
}

func Constructor() MyHashMap {
	hashMapSlice := make([]int, 1000001, 1000001)
	for i, _ := range hashMapSlice {
		hashMapSlice[i] = -1
	}

	return MyHashMap{content: hashMapSlice}
}

func (this *MyHashMap) Put(key int, value int) {
	this.content[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.content[key]
}

func (this *MyHashMap) Remove(key int) {
	this.content[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
