package dojo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInsertSort1(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1}

	result := insertSort(arr)
	expected := []int{1}
	assert.Equal(expected, result, "must be array [1] should expected [1]")
}

func TestIsInsertSort2(t *testing.T) {
	assert := assert.New(t)
	arr := []int{2, 1}

	result := insertSort(arr)
	expected := []int{1, 2}
	assert.Equal(expected, result, "must be array [2,1] should expected [1,2]")
}

func TestIsInsertSort3(t *testing.T) {
	assert := assert.New(t)
	arr := []int{2, 1, 3}

	result := insertSort(arr)
	expected := []int{1, 2, 3}
	assert.Equal(expected, result, "must be array [2,1,3] should expected [1,2,3]")
}
func TestIsInsertSort4(t *testing.T) {
	assert := assert.New(t)
	arr := []int{40, 37, 95, 42, 23, 51, 27}

	result := insertSort(arr)
	expected := []int{23, 27, 37, 40, 42, 51, 95}
	assert.Equal(expected, result, "must be array [40,37,95,42,23,51,27] should expected [23,27,37,40,42,51,95]")
}

func TestIscanChangePosition1(t *testing.T) {
	assert := assert.New(t)
	arr := []int{2, 1, 3}
	valueInit := 1
	position := 0

	insertPositionResp, arrResp := canChangePosition(position, valueInit, arr)

	insertPositionExpected := -1
	arrExpected := []int{2, 2, 3}
	assert.Equal(insertPositionExpected, insertPositionResp, "must be init position 0 should expected init position -1")
	assert.Equal(arrExpected, arrResp, "must be array [2,1,3] should expected [2,2,3]")
}

func TestIscanChangePosition2(t *testing.T) {
	assert := assert.New(t)
	arr := []int{2, 1, 3}
	valueInit := 3
	position := 1

	insertPositionResp, arrResp := canChangePosition(position, valueInit, arr)

	insertPositionExpected := 1
	arrExpected := []int{2, 1, 3}
	assert.Equal(insertPositionExpected, insertPositionResp, "must be init position 1 should expected init position 1")
	assert.Equal(arrExpected, arrResp, "must be array [2,1,3] should expected [2,1,3]")
}
