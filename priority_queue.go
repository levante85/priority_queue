package pqueue

// New instantiate and returns a new priority queue
func New(comp Comparator) *Heap {
	return &Heap{make([]interface{}, 1), comp}
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
	for 2*k <= q.Size() {
		j := 2 * k

		if j < q.Size() && q.compare(j, j+1, q.slice) {
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
	q.swim(q.Size())
}

// PopTop removes and returns the max or min ( depending on the comparator function )
// element in the queue
func (q *Heap) PopTop() interface{} {
	top := q.slice[1]
	q.swap(1, q.Size())
	q.slice = q.slice[0:q.Size()]
	q.sink(1)

	return top
}

// Top returns the max or min element ( depending on the comparator function )
// in the queue but doesn't remove it
func (q *Heap) Top() interface{} {
	return q.slice[1]
}

// Size returns the size of the underling slice with deleted elements
func (q *Heap) Size() int {
	return len(q.slice) - 1
}

// Empty returns whether the priority queue contains no elements
func (q *Heap) Empty() bool {
	return q.Size() == 0
}
