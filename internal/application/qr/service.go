package qr

import (
	"context"
	"errors"

	"github.com/brayanbst/matrix-service-go/internal/domain/matrix"
)

// Stats representa la respuesta de estadísticas del servicio Node.
type Stats struct {
	MaxValue  float64 `json:"maxValue"`
	MinValue  float64 `json:"minValue"`
	Average   float64 `json:"average"`
	TotalSum  float64 `json:"totalSum"`
	Diagonals []bool  `json:"diagonals"`
}

// StatsPort es el puerto para cualquier cliente de estadísticas (Node).
type StatsPort interface {
	ComputeStats(ctx context.Context, matrices [][][]float64, authHeader string) (*Stats, error)
}

// Service es el caso de uso principal para QR.
type Service struct {
	qrDecomposer matrix.QRDecomposer
	statsPort    StatsPort
}

var ErrStatsPortNotConfigured = errors.New("stats port is not configured")

// NewService crea un nuevo Service.
// statsPort es opcional; si no se pasa, las funciones *AndStats fallarán con ErrStatsPortNotConfigured.
func NewService(qrDecomposer matrix.QRDecomposer, statsPort ...StatsPort) *Service {
	s := &Service{qrDecomposer: qrDecomposer}
	if len(statsPort) > 0 {
		s.statsPort = statsPort[0]
	}
	return s
}

// ComputeQR realiza la factorización QR de una matriz.
func (s *Service) ComputeQR(ctx context.Context, m *matrix.Matrix) (*matrix.QRResult, error) {
	return s.qrDecomposer.Decompose(m)
}

// ComputeQRAndStats realiza la factorización QR y luego pide estadísticas al servicio Node.
func (s *Service) ComputeQRAndStats(ctx context.Context, m *matrix.Matrix, authHeader string) (*matrix.QRResult, *Stats, error) {
	if s.statsPort == nil {
		return nil, nil, ErrStatsPortNotConfigured
	}

	qrResult, err := s.qrDecomposer.Decompose(m)
	if err != nil {
		return nil, nil, err
	}

	// Preparamos las matrices Q y R para enviarlas a Node.
	matrices := [][][]float64{qrResult.Q, qrResult.R}

	stats, err := s.statsPort.ComputeStats(ctx, matrices, authHeader)
	if err != nil {
		return nil, nil, err
	}

	return qrResult, stats, nil
}
