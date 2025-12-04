package qr

import (
	"gonum.org/v1/gonum/mat"

	"github.com/brayanbst/matrix-service-go/internal/domain/matrix"
)

// GonumQRDecomposer implementa QRDecomposer usando la librería Gonum.
type GonumQRDecomposer struct{}

func NewGonumQRDecomposer() *GonumQRDecomposer {
	return &GonumQRDecomposer{}
}

func (g *GonumQRDecomposer) Decompose(m *matrix.Matrix) (*matrix.QRResult, error) {
	rows := len(m.Data)
	cols := len(m.Data[0])

	flat := make([]float64, 0, rows*cols)
	for i := 0; i < rows; i++ {
		flat = append(flat, m.Data[i]...)
	}

	A := mat.NewDense(rows, cols, flat)

	var qr mat.QR
	qr.Factorize(A)

	var Qmat, Rmat mat.Dense
	qr.QTo(&Qmat)
	qr.RTo(&Rmat)

	return &matrix.QRResult{
		Q: denseToSlices(&Qmat),
		R: denseToSlices(&Rmat),
	}, nil
}

func denseToSlices(m *mat.Dense) [][]float64 {
	r, c := m.Dims()
	out := make([][]float64, r)
	for i := 0; i < r; i++ {
		out[i] = make([]float64, c)
		for j := 0; j < c; j++ {
			out[i][j] = m.At(i, j)
		}
	}
	return out
}
