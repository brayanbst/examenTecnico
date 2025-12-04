package matrix

// Matrix representa una matriz rectangular.
type Matrix struct {
	Data [][]float64
}

// NewMatrix valida filas / columnas y crea Matrix.
func NewMatrix(data [][]float64) (*Matrix, error) {
	if len(data) == 0 || len(data[0]) == 0 {
		return nil, ErrEmptyMatrix
	}
	cols := len(data[0])
	for i := range data {
		if len(data[i]) != cols {
			return nil, ErrInvalidShape
		}
	}
	return &Matrix{Data: data}, nil
}
