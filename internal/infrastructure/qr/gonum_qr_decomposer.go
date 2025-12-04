package qr

import (
	"fmt"

	"github.com/brayanbst/matrix-service-go/internal/domain/matrix"
	"gonum.org/v1/gonum/mat"
)

// GonumQRDecomposer implementa QRDecomposer usando gonum.
type GonumQRDecomposer struct{}

// NewGonumQRDecomposer crea una nueva instancia de descomposición QR con gonum.
func NewGonumQRDecomposer() matrix.QRDecomposer {
	return &GonumQRDecomposer{}
}

// Decompose realiza la factorización QR de la matriz usando gonum.
func (d *GonumQRDecomposer) Decompose(m *matrix.Matrix) (*matrix.QRResult, error) {
	if m == nil {
		return nil, fmt.Errorf("matrix is nil")
	}

	r, c := m.Rows, m.Cols
	if r == 0 || c == 0 {
		return nil, fmt.Errorf("matrix must have at least one row and one column")
	}

	// Convertimos [][]float64 a un slice plano para gonum.
	data := make([]float64, 0, r*c)
	for _, row := range m.Values {
		data = append(data, row...)
	}

	a := mat.NewDense(r, c, data)

	var qr mat.QR
	qr.Factorize(a)

	var qMat, rMat mat.Dense

	// QTo y RTo no devuelven error, solo llenan las matrices destino.
	qr.QTo(&qMat)
	qr.RTo(&rMat)

	// Convertimos Q de gonum a [][]float64
	qRows, qCols := qMat.Dims()
	Q := make([][]float64, qRows)
	for i := 0; i < qRows; i++ {
		Q[i] = make([]float64, qCols)
		for j := 0; j < qCols; j++ {
			Q[i][j] = qMat.At(i, j)
		}
	}

	// Convertimos R de gonum a [][]float64
	rRows, rCols := rMat.Dims()
	R := make([][]float64, rRows)
	for i := 0; i < rRows; i++ {
		R[i] = make([]float64, rCols)
		for j := 0; j < rCols; j++ {
			R[i][j] = rMat.At(i, j)
		}
	}

	return &matrix.QRResult{
		Q: Q,
		R: R,
	}, nil
}
