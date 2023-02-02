type keyValuePair struct {
	key   int
	value int
}

type MyHashMap struct {
	content []*keyValuePair
}

func Constructor() MyHashMap {
	hashMapSlice := make([]*keyValuePair, 0)

	return MyHashMap{content: hashMapSlice}
}

func (this *MyHashMap) Put(key int, value int) {
	for _, v := range this.content {
		if v.key == key {
			v.value = value
			return
		}
	}
	this.content = append(this.content, &keyValuePair{key, value})
}

func (this *MyHashMap) Get(key int) int {
	for _, v := range this.content {
		if v.key == key {
			return v.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	for i, v := range this.content {
		if v.key == key {
			copy(this.content[i:], this.content[i+1:])
			this.content = this.content[:len(this.content)-1]
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
