package tests

import (
	"conway-v2/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcNumNeighborsWithNeighbors(t *testing.T) {
	height := 3
	width := 3

	cells := make([][]int, height)

	for i := 0; i < height; i++ {
		cells[i] = make([]int, width)
	}

	// 010
	// 101
	// 010
	cells[0][1] = 1
	cells[1][0] = 1
	cells[1][2] = 1
	cells[2][1] = 1

	neighbors := utils.CalcNumNeighbors(1, 1, height, width, &cells)

	assert.Equal(t, 4, neighbors)
}

func TestCalcNumNeighborsWithNoNeighbors(t *testing.T) {
	height := 3
	width := 4

	cells := make([][]int, height)

	for i := 0; i < height; i++ {
		cells[i] = make([]int, width)
	}

	// 1000
	// 0000
	// 0000
	cells[0][0] = 1

	neighbors := utils.CalcNumNeighbors(1, 2, height, width, &cells)

	assert.Equal(t, 0, neighbors)
}

func TestCalcNumNeighborsWrappedAround(t *testing.T) {
	height := 3
	width := 3

	cells := make([][]int, height)

	for i := 0; i < height; i++ {
		cells[i] = make([]int, width)
	}

	// 100
	// 100
	// 000
	cells[0][0] = 1
	cells[1][0] = 1

	neighbors := utils.CalcNumNeighbors(1, 2, height, width, &cells)

	assert.Equal(t, 2, neighbors)
}

func TestCalCurrentStateCond1(t *testing.T) {
	neighbors := 1
	prev_state := 1
	state := utils.CalcCurrentState(neighbors, prev_state)

	assert.Equal(t, 0, state)
}

func TestCalCurrentStateCond2(t *testing.T) {
	neighbors := 2
	prev_state := 1
	state := utils.CalcCurrentState(neighbors, prev_state)

	assert.Equal(t, 1, state)
}

func TestCalCurrentStateCond3(t *testing.T) {
	neighbors := 4
	prev_state := 1
	state := utils.CalcCurrentState(neighbors, prev_state)

	assert.Equal(t, 0, state)
}

func TestCalCurrentStateCond4(t *testing.T) {
	neighbors := 3
	prev_state := 0
	state := utils.CalcCurrentState(neighbors, prev_state)

	assert.Equal(t, 1, state)
}
