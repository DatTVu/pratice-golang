package vector

import "testing"

func TestNewVector(t *testing.T) {
	vec := NewVector()
	if vec == nil {
		t.Errorf("Could not construct new vector")
	}

	if len(vec.kData) != kDefaultCapacity {
		t.Errorf("Length of the vector is incorrect, want: %d got: %d", 0, len(vec.kData))
	}

	if cap(vec.kData) != kDefaultCapacity {
		t.Errorf("Capacity of the vector is incorrect, want: %d got: %d", kDefaultCapacity, cap(vec.kData))
	}
}

func TestNewVectorWithCapacity(t *testing.T) {
	tables := []uint{
		1,
		5,
		10,
		103,
		5003,
	}

	for _, table := range tables {
		vec, _ := NewVectorWithCapacity(table)
		if vec == nil {
			t.Errorf("Could not construct new vector")
		}

		if len(vec.kData) != int(table) {
			t.Errorf("Length of the vector is incorrect, want: %d got: %d", 0, len(vec.kData))
		}

		if cap(vec.kData) != int(table) {
			t.Errorf("Capacity of the vector is incorrect, want: %d got: %d", table, cap(vec.kData))
		}
	}
}

func TestSize(t *testing.T) {
	tables := []struct {
		size     int
		capacity uint
	}{
		{10, 10},
		{10, 20},
		{1231232, 2000000},
	}

	for _, table := range tables {
		vec, _ := NewVectorWithCapacity(table.capacity)
		for i := 0; i < table.size; i++ {
			vec.kData[i] = i
		}
		if vec.Size() != int(table.capacity) {
			t.Errorf("Wrong vector size, want: %d, got: %d", table.size, vec.Size())
		}
	}
}

func TestPush(t *testing.T) {
	tables := []struct {
		size     int
		capacity uint
	}{
		{10, 10},
		{10, 20},
		{12312321, 20000000},
	}

	for _, table := range tables {
		vec, _ := NewVectorWithCapacity(table.capacity)
		for i := 0; i < int(table.capacity); i++ {
			vec.Push(i)
			if vec.kData[i] != i {
				t.Errorf("Push the wrong value, want: %d, got: %d", i, vec.kData[i])
			}
		}
	}
}

// func TestInsert(t *testing.T) {
// 	tables := []struct {
// 		val int
// 		idx int
// 	}{
// 		{123, 4444123},
// 		{10, 20},
// 		{12312321, 20000000},
// 	}

// 	for _, table := range tables {
// 		vec, _ := NewVectorWithCapacity(20000001)
// 		err := vec.Insert(table.val, table.idx)
// 		if err != nil {
// 			t.Errorf("Trouble when inserting value: %d into position: %d. Error message is: %v", table.val, table.idx, err)
// 		}

// 		val, _ := vec.At(table.idx)
// 		if val != table.val {
// 			t.Errorf("Trouble when inserting value: %d into position: %d, got: %d", table.val, table.idx, val)
// 		}
// 	}
// }

// func TestPrepend(t *testing.T) {
// 	vec, _ := NewVectorWithCapacity(10000)
// 	for i := 0; i < 12000; i++ {
// 		err := vec.Prepend(i)
// 		if err != nil {
// 			t.Errorf("Trouble when prepend the vector: %v", err)
// 		}
// 		val, _ := vec.At(0)
// 		if val != i {
// 			t.Errorf("Trouble when prepend the vector, want: %d got: %d", i, val)
// 		}
// 	}
// }

// func TestPop(t *testing.T) {
// 	vec, _ := NewVectorWithCapacity(10000)
// 	for i := 0; i < 12000; i++ {
// 		vec.kData[i] = i
// 	}
// 	for i := 11999; i >= 0; i-- {
// 		val, err := vec.Pop()
// 		if err != nil {
// 			t.Errorf("Trouble when pop the vector: %v", err)
// 		}

// 		valtest, _ := vec.At(i)

// 		if val != i {
// 			t.Errorf("Trouble when pop the vector, want: %d got: %d", i, valtest)
// 		}
// 	}
// }
