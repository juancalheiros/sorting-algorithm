package dojo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorted1(t *testing.T) {
	//Assert
	assert := assert.New(t)
	array := []int{10}
	expected := []int{10}

	//Arrange
	result := bubbleSorted(array)

	//Act
	assert.Equal(expected, result, "should [10] return sorted [10]")
}

func TestSorted2(t *testing.T) {
	//Assert
	assert := assert.New(t)
	array := []int{10, 0}
	expected := []int{0, 10}

	//Arrange
	result := bubbleSorted(array)

	//Act
	assert.Equal(expected, result, "should [10, 0] return sorted [0,10]")
}
func TestSorted3(t *testing.T) {
	//Assert
	assert := assert.New(t)
	array := []int{2, 0, 1}
	expected := []int{0, 1, 2}

	//Arrange
	result := bubbleSorted(array)

	//Act
	assert.Equal(expected, result, "should [2,0,1] return sorted [0,1,2]")
}

func TestSorted4(t *testing.T) {
	//Assert
	assert := assert.New(t)
	array := []int{2, 0, -4}
	expected := []int{-4, 0, 2}

	//Arrange
	result := bubbleSorted(array)

	//Act
	assert.Equal(expected, result, "should [2,0,-4] return sorted [-4,0,2]")
}

func TestSorted5(t *testing.T) {
	//Assert
	assert := assert.New(t)
	array := []int{10, 0, -3}
	expected := []int{-3, 0, 10}

	//Arrange
	result := bubbleSorted(array)

	//Act
	assert.Equal(expected, result, "should [10, 0, -3] return sorted [-3,0,10]")
}

func TestSorted6(t *testing.T) {
	//Assert
	assert := assert.New(t)
	array := []int{5, 2, 10, 0}
	expected := []int{0, 2, 5, 10}

	//Arrange
	result := bubbleSorted(array)

	//Act
	assert.Equal(expected, result, "should [5, 2, 10, 0] return sorted [0,2,5,10]")
}

func TestSorted7(t *testing.T) {
	//Assert
	assert := assert.New(t)
	array := []int{5, 2, 10, 0, 40, -20, 11}
	expected := []int{-20, 0, 2, 5, 10, 11, 40}

	//Arrange
	result := bubbleSorted(array)

	//Act
	assert.Equal(expected, result, "should [5, 2, 10, 0, 40, -20, 11] return sorted [-20,0,2,5,10,11,40]")
}

func TestChangePosition1(t *testing.T) {
	//Assert
	assert := assert.New(t)
	arr := []int{5, 2}
	pointer := &arr
	index := 0
	expected := []int{2, 5}
	expectedPointerArr := &expected

	//Arrange
	result := canChangePosition(pointer, index)

	//Act
	assert.Equal(expectedPointerArr, result, "should [5, 2] return sorted [2,5]")
}

func TestChangePosition2(t *testing.T) {
	//Assert
	assert := assert.New(t)
	arr := []int{5, 2, 10}
	pointer := &arr
	index := 1
	expected := []int{5, 2, 10}
	expectedPointerArr := &expected

	//Arrange
	result := canChangePosition(pointer, index)

	//Act
	assert.Equal(expectedPointerArr, result, "should [5, 2, 10] return sorted [5,2,10]")
}

func TestChangePosition3(t *testing.T) {
	//Assert
	assert := assert.New(t)
	arr := []int{5, 2, 10, 15, 2, -5}
	pointer := &arr
	index := 3
	expected := []int{5, 2, 10, 2, 15, -5}
	expectedPointerArr := &expected

	//Arrange
	result := canChangePosition(pointer, index)

	//Act
	assert.Equal(expectedPointerArr, result, "should 5, 2, 10,15,2,-5] return sorted [5, 2, 10,2,15,-5]")
}
