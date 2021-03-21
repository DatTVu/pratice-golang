package vector

import (
	"log"
	"sync"
)

const kDefaultCapacity = 16

type Vector struct {
	kCapacity int
	kSize     int
	kData     []int
	mutex     sync.Mutex
}

func (vec Vector) Size() int {
	return vec.kSize
}

func (vec Vector) Capacity() int {
	return vec.kCapacity
}

func (vec Vector) Empty() bool {
	if vec.kSize == 0 {
		return true
	} else {
		return false
	}
}

func (vec Vector) At(idx int) int {
	if idx > vec.kSize || idx > vec.kCapacity {
		log.Panic("What T")
	}
}
