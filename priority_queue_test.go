package pqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func debug(s []interface{}, t *testing.T) {
	t.Logf("Content %+v\n", s)
}

func compare(i, j int, values []interface{}) bool {
	return values[i].(int) < values[j].(int)
}

func TestInsert(t *testing.T) {
	pq := New(compare)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Count(), len(num), "Inserting went wrong")
	debug(pq.slice, t)
}

func TestInserttop(t *testing.T) {
	pq := New(compare)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Count(), len(num), "Inserting went wrong")
	top := pq.Top().(int)
	debug(pq.slice, t)
	assert.Equal(t, top, 9, "Inserting went wrong")

}

func TestInsertPopTop(t *testing.T) {
	pq := New(compare)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Count(), len(num), "Inserting went wrong")
	assert.Equal(t, pq.PopTop().(int), 9, "Pop not working")
	assert.Equal(t, pq.Count(), 8, "PopTop went wrong")

}

func TestInsertPopTopMulti(t *testing.T) {
	pq := New(compare)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Count(), len(num), "Inserting went wrong")
	assert.Equal(t, pq.PopTop().(int), 9, "Pop not working")
	assert.Equal(t, pq.Count(), 8, "PopTop went wrong")
	assert.Equal(t, pq.PopTop().(int), 8, "Pop not working")
	assert.Equal(t, pq.Count(), 7, "PopTop went wrong")

	debug(pq.slice, t)
}
