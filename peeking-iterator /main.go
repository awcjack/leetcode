/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter    *Iterator
	nextVal int
}

func Constructor(iter *Iterator) *PeekingIterator {
	if iter.hasNext() {
		nextVal := iter.next()
		return &PeekingIterator{iter, nextVal}
	}
	return &PeekingIterator{iter, 0}
}

func (this *PeekingIterator) hasNext() bool {
	if this.nextVal != 0 {
		return true
	}
	return false
}

func (this *PeekingIterator) next() int {
	returnVal := this.nextVal
	if this.iter.hasNext() {
		nextVal := this.iter.next()
		this.nextVal = nextVal
		return returnVal
	}
	this.nextVal = 0
	return returnVal
}

func (this *PeekingIterator) peek() int {
	return this.nextVal
}
