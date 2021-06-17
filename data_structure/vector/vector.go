package vector

import (
	"errors"
	"fmt"
	"sync"
)

const kDefaultCapacity = 16

type Vector struct {
	kData []int
	mutex sync.Mutex
}

func NewVector() *Vector {
	return &Vector{kData: make([]int, kDefaultCapacity, kDefaultCapacity)}
}

func NewVectorWithCapacity(capacity uint) (*Vector, error) {
	return &Vector{kData: make([]int, capacity, capacity)}, nil
}

func (vec *Vector) Size() int {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	return len(vec.kData)
}

func (vec *Vector) Capacity() int {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	return cap(vec.kData)
}

func (vec *Vector) Empty() bool {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	if len(vec.kData) == 0 {
		return true
	} else {
		return false
	}
}

func (vec *Vector) At(idx int) (int, error) {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	if idx > len(vec.kData) || idx > cap(vec.kData) {
		return -1, errors.New("[ERROR] Out of Range!")
	} else {
		return vec.kData[idx], nil
	}
}

func (vec *Vector) Push(val int) {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	if len(vec.kData) > cap(vec.kData) {
		cnt := 0
		fmt.Println("Resizing for %d", cnt)
		vec.resize()
		cnt+
	}
	vec.kData[len(vec.kData)-1] = val
}

func (vec *Vector) Insert(val int, idx int) error {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	if val < 0 {
		return errors.New("[ERROR] Out of Range!")
	}
	if len(vec.kData) > cap(vec.kData) {
		vec.resize()
	}
	for i := len(vec.kData); i > idx; i-- {
		vec.kData[i] = vec.kData[i-1]
	}

	vec.kData[idx] = val
	return nil
}

func (vec *Vector) Prepend(val int) error {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	return vec.Insert(val, 0)
}

func (vec *Vector) Pop() (int, error) {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	if len(vec.kData) == 0 {
		return -1, errors.New("[ERROR] Out of range!")
	}
	val := vec.kData[len(vec.kData)-1]
	vec.kData = vec.kData[:len(vec.kData)-2]
	return val, nil
}

func (vec *Vector) Delete(idx int) error {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	if idx < 0 || idx >= len(vec.kData) || idx >= cap(vec.kData) {
		return errors.New("[ERROR] Out of range!")
	}
	for i := idx; i < len(vec.kData)-2; i++ {
		vec.kData[i] = vec.kData[i+1]
	}
	vec.kData = vec.kData[:len(vec.kData)-2]
	return nil
}

func (vec *Vector) Remove(val int) error {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	for idx, v := range vec.kData {
		if v == val {
			vec.Delete(idx)
			return nil
		}
	}
	return errors.New("[ERROR] Can't find value in the vector")
}

func (vec *Vector) Find(val int) (int, error) {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	for idx, v := range vec.kData {
		if v == val {
			return idx, nil
		}
	}
	return -1, errors.New("[ERROR] Can't find value in the vector")
}

func (vec *Vector) resize() {
	vec.mutex.Lock()
	defer vec.mutex.Unlock()
	temp := make([]int, 0, 2*cap(vec.kData))
	for idx, val := range vec.kData {
		temp[idx] = val
	}
	vec.kData = temp[:]
}
