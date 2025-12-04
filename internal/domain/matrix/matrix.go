package matrix

import (
	"errors"
	"fmt"
)

// Matrix representa una matriz de números reales.
type Matrix struct {
	Values [][]float64 `json:"values"`
	Rows   int         `json:"rows"`
	Cols   int         `json:"cols"`
}

// NewMatrix valida y crea una nueva Matrix a partir de un slice 2D.
func NewMatrix(values [][]float64) (*Matrix, error) {
	if len(values) == 0 {
		return nil, errors.New("matrix cannot be empty")
	}

	rows := len(values)
	cols := len(values[0])
	if cols == 0 {
		return nil, errors.New("matrix must have at least one column")
	}

	// Validar que todas las filas tengan la misma cantidad de columnas
	for i, row := range values {
		if len(row) != cols {
			return nil, fmt.Errorf("row %d has different length: expected %d, got %d", i, cols, len(row))
		}
	}

	return &Matrix{
		Values: values,
		Rows:   rows,
		Cols:   cols,
	}, nil
}

// QRResult contiene el resultado de la descomposición QR.
type QRResult struct {
	Q [][]float64 `json:"Q"`
	R [][]float64 `json:"R"`
}

// QRDecomposer es la interfaz que implementa cualquier descomposición QR.
type QRDecomposer interface {
	Decompose(m *Matrix) (*QRResult, error)
}
