package matrix

import (
	"errors"
)

type Matrix struct {
	data []float64
	rows int
	cols int
	step int
}

func NewMatrix(data []float64, rows, cols int) *Matrix {
	m := new(Matrix)
	m.rows = rows
	m.cols = cols
	m.step = cols
	m.data = data
	return m
}

func (m *Matrix) Rows() int {
	return m.rows
}

func (m *Matrix) Cols() int {
	return m.cols
}

func (m *Matrix) Step() int {
	return m.step
}

func (m *Matrix) Get(i, j int) (float64, error) {
	index := i*m.Step() + j
	if index >= len(m.data) {
		return 0, errors.New("Index out of range")
	}
	return m.data[index], nil
}

func (m *Matrix) Set(i, j int, v float64) error {
	index := i*m.Step() + j
	if index >= len(m.data) {
		return errors.New("Index out of range")
	}
	m.data[index] = v
	return nil
}
