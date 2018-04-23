package pqueue

// New instantiate and returns a new priority queue
func New(comp Comparator) *Heap {
	return &Heap{make([]interface{}, 1), comp, 0}
}

// Comparator is the type of the comparision function that needs to be
// passed to the heap, if the comparison is equals to less than ie
// values[i] < values[j] the queue becomes a max priority queue instead
// if values[i] > values[j] the queue becomes a min priority queue.
// An example is given in the test file
type Comparator func(i, j int, values []interface{}) bool

// Heap is the priority queue concrete implementation
type Heap struct {
	slice   []interface{}
	compare Comparator
	num     int
}

func (q *Heap) swap(i, j int) {
	q.slice[i], q.slice[j] = q.slice[j], q.slice[i]
}

func (q *Heap) swim(k int) {
	for k > 1 && q.compare(k/2, k, q.slice) {
		q.swap(k/2, k)
		k /= 2
	}
}

func (q *Heap) sink(k int) {
	for 2*k <= len(q.slice) {
		j := 2 * k
		if j < len(q.slice) && q.compare(j, j+1, q.slice) {
			j++
		}
		if !q.compare(k, j, q.slice) {
			break
		}
		q.swap(k, j)
		k = j
	}
}

// Insert adds an element to the queue
func (q *Heap) Insert(element interface{}) {
	q.slice = append(q.slice, element)
	q.num++
	q.swim(q.num)
}

// PopTop removes and returns the max or min ( depending on the comparator function )
// element in the queue
func (q *Heap) PopTop() interface{} {
	max := q.slice[1]
	q.num--
	q.swap(1, q.num)
	q.slice[q.num+1] = -1
	q.sink(1)

	return max
}

// Top returns the max or min element ( depending on the comparator function )
// in the queue but doesn't remove it
func (q *Heap) Top() interface{} {
	return q.slice[1]
}

// Size returns the size of the underling slice with deleted elements
func (q *Heap) Size() int {
	return len(q.slice)
}

// Count returns the actual number of elements present in the queue
func (q *Heap) Count() int {
	return q.num
}
