package qr

import (
	"context"

	"github.com/brayanbst/matrix-service-go/internal/domain/matrix"
)

// Service es el caso de uso principal para QR.
type Service struct {
	qrDecomposer matrix.QRDecomposer
}

// NewService crea un nuevo Service.
func NewService(qrDecomposer matrix.QRDecomposer) *Service {
	return &Service{qrDecomposer: qrDecomposer}
}

// ComputeQR realiza la factorización QR de una matriz.
func (s *Service) ComputeQR(ctx context.Context, m *matrix.Matrix) (*matrix.QRResult, error) {
	return s.qrDecomposer.Decompose(m)
}
