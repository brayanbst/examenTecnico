package matrix

// QRResult representa el resultado de la factorización QR.
type QRResult struct {
	Q [][]float64
	R [][]float64
}

// QRDecomposer es el puerto de dominio para cualquier implementación de QR.
type QRDecomposer interface {
	Decompose(m *Matrix) (*QRResult, error)
}
