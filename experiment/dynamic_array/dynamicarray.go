package dynamicarray

import (
	"errors"
)

var defaultCapacity = 10

type DynamicArray struct {
	size        int
	capacity    int
	elementData []interface{}
}

func (da *DynamicArray) CheckRangeFromIndex(index int) error {
	if index >= da.size || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}

func (da *DynamicArray) NewCapacity() {
	if da.capacity == 0 {
		da.capacity = defaultCapacity
	} else {
		da.capacity = da.capacity << 1
	}

	newDataElement := make([]interface{}, da.capacity)

	copy(newDataElement, da.elementData)

	da.elementData = newDataElement
}

func (da *DynamicArray) CheckArrayLen(index int) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}
	return nil
}

func (da *DynamicArray) Put(index int, element interface{}) error {
	da.CheckArrayLen(index)

	da.elementData[index] = element

	return nil
}

func (da *DynamicArray) Add(element interface{}) {
	if da.size == da.capacity {
		da.NewCapacity()
	}
	da.elementData[da.size] = element
	da.size++
}

func (da *DynamicArray) Remove(index int) error {
	da.CheckArrayLen(index)

	copy(da.elementData[index:], da.elementData[index+1:da.size])
	da.elementData[da.size-1] = nil

	da.size--

	return nil
}

func (da *DynamicArray) Get(index int) (interface{}, error) {
	da.CheckArrayLen(index)

	return da.elementData[index], nil
}

func (da *DynamicArray) IsEmpty() bool {
	return da.size == 0
}

func (da *DynamicArray) GetData() []interface{} {
	return da.elementData[:da.size]
}
