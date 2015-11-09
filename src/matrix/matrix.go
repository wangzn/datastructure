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

func (m *Matrix) AddValue(i, j int, v float64) error {
	value, err := m.Get(i, j)
	if err != nil {
		return err
	}
	return m.Set(i, j, value+v)
}

func (m *Matrix) SubValue(i, j int, v float64) error {
	value, err := m.Get(i, j)
	if err != nil {
		return err
	}
	return m.Set(i, j, value-v)
}

func (m *Matrix) MultiplyValue(i, j int, v float64) error {
	value, err := m.Get(i, j)
	if err != nil {
		return err
	}
	return m.Set(i, j, value*v)
}

func (m *Matrix) Add(n *Matrix) error {
	if m.Rows() != n.Rows() || m.Cols() != n.Cols() {
		return errors.New("Matrix size does not match")
	}
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			v, _ := n.Get(i, j)
			m.AddValue(i, j, v)
		}
	}
	return nil
}

func (m *Matrix) Sub(n *Matrix) error {
	if m.Rows() != n.Rows() || m.Cols() != n.Cols() {
		return errors.New("Two matrix sizes are not same")
	}
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			v, _ := n.Get(i, j)
			m.SubValue(i, j, v)
		}
	}
	return nil
}

func (m *Matrix) Multiply(n *Matrix) error {
	if m.Rows() != n.Rows() || m.Cols() != n.Cols() {
		return errors.New("Two matrix sizes are not same")
	}
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			v, _ := n.Get(i, j)
			m.MultiplyValue(i, j, v)
		}
	}
	return nil
}

func (m *Matrix) DiagonalCopy() []float64 {
	diag := make([]float64, m.Cols())
	for i := 0; i < len(diag); i++ {
		diag[i], _ = m.Get(i, i)
	}
	return diag
}

func (m *Matrix) copy() *Matrix {
	n := new(Matrix)
	n.rows = m.Rows()
	n.cols = m.Cols()
	n.step = m.Step()

	n.data = make([]float64, m.Cols()*m.Rows())

	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			v, _ := m.Get(i, j)
			n.Set(i, j, v)
		}
	}
	return n
}
