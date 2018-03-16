package priority_queue

import "testing"

func debug(s []int, t *testing.T) {
	t.Logf("Content %+v\n", s)
}

func TestInsert(t *testing.T) {
	pq := New()
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range nums {
		pq.Insert(nums[n])
	}

	if pq.Count() != len(nums) {
		t.Fatal("Inserting went wrong!")
	}

	debug(pq.slice, t)
}

func TestInsertMax(t *testing.T) {
	pq := New()
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range nums {
		pq.Insert(nums[n])
	}

	if pq.Count() != len(nums) {
		t.Fatalf("Inserting went wrong %v\n", pq.Size())
	}

	if max := pq.Max(); max != 9 {
		t.Fatalf("Max not working %v\n", max)
	}
}

func TestInsertPopMax(t *testing.T) {
	pq := New()
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range nums {
		pq.Insert(nums[n])
	}

	if pq.Count() != len(nums) {
		t.Fatal("Inserting went wrong!")
	}

	if max := pq.PopMax(); max != 9 || pq.Count() != 8 {
		t.Fatalf("PopMax not working, max %v, size: %v\n", max, pq.Size())
	}
}

func TestInsertPopMaxMulti(t *testing.T) {
	pq := New()
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range nums {
		pq.Insert(nums[n])
	}

	if pq.Count() != len(nums) {
		t.Fatal("Inserting went wrong!")
	}

	if max := pq.PopMax(); max != 9 || pq.Count() != 8 {
		t.Fatalf("PopMax not working, max %v, size: %v\n", max, pq.Size())
	}
	if max := pq.PopMax(); max != 8 || pq.Count() != 7 {
		t.Fatalf("PopMax not working, max %v, size: %v\n", max, pq.Size())
	}

	debug(pq.slice, t)
}
