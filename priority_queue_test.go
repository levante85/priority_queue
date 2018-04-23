package pqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func less(i, j int, values []interface{}) bool {
	l, _ := values[i].(int)
	r, _ := values[j].(int)
	return l < r
}

func TestInsert(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Size(), len(num), "Inserting went wrong")
}

func TestInserttop(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Size(), len(num), "Inserting went wrong")
	top := pq.Top().(int)
	assert.Equal(t, top, 9, "Inserting went wrong")

}

func TestPopTopEdge3Elements(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3}
	for n := range num {
		pq.Insert(num[n])
	}

	_ = pq.PopTop()
}

func TestInsertPopTop(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Size(), len(num), "Inserting went wrong")
	assert.Equal(t, pq.PopTop().(int), 9, "Pop not working")
	assert.Equal(t, pq.Size(), 8, "PopTop went wrong")

}

func TestInsertPopTopMulti(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Size(), len(num), "Inserting went wrong")
	assert.Equal(t, pq.PopTop().(int), 9, "Pop not working")
	assert.Equal(t, pq.Size(), 8, "PopTop went wrong")
	assert.Equal(t, pq.PopTop().(int), 8, "Pop not working")
	assert.Equal(t, pq.Size(), 7, "PopTop went wrong")

}

func TestInsertPopTopAll(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for n := range num {
		pq.Insert(num[n])
	}

	for n := len(num); !pq.Empty(); n-- {
		v := pq.PopTop()
		assert.Equal(t, n, v.(int), "Not ok should be equal")
	}
}

func TestSink(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for n := range num {
		pq.Insert(num[n])
	}

	_ = pq.PopTop()
	_ = pq.PopTop()

	res := []int{9, 8, 6, 7, 3, 2, 5, 1, 4}
	for i := 1; i < len(pq.slice); i++ {
		v := pq.slice[i].(int)
		if res[i-1] != v {
			t.Log(res)
			t.Log(pq.slice[1:])
			t.Log(v, res[i])
			t.Fatal("The slice isn't in the state it should be")
		}
	}
}
