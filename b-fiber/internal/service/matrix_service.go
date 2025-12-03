package service

import (
	"github.com/antoniougarte/b-fiber/internal/client"
	"github.com/antoniougarte/b-fiber/internal/matrix"
	"github.com/antoniougarte/b-fiber/internal/model"
)

var nodeClient = client.NewNodeJSClient("http://node-api:5000")

// ProcessMatrix ahora acepta el token
func ProcessMatrix(req model.MatrixRequest, token string) (model.MatrixResponse, error) {
	// Validar la matriz original (rectangular, valores finitos)
	if err := matrix.ValidMatrix(req.Matrix); err != nil {
		return model.MatrixResponse{}, err
	}

	// Rotar primero (seguir mostrando la rotada en la respuesta)
	rotated := matrix.Rotate90Clockwise(req.Matrix)

	// Calcular QR sobre LA MATRIZ ORIGINAL (no la rotada).
	// Esto evita intentar QR en matrices con filas < cols resultantes de la rotación.
	Q, R, err := matrix.QRFactor(req.Matrix)
	if err != nil {
		return model.MatrixResponse{}, err
	}

	statsReq := client.StatsRequest{
		Rotated: rotated,
		Q:       Q,
		R:       R,
	}

	stats, err := nodeClient.GetStatistics(statsReq, token)
	if err != nil {
		return model.MatrixResponse{}, err
	}

	resp := model.MatrixResponse{
		Rotated: rotated,
		Q:       Q,
		R:       R,
		Stats: &model.Statistics{
			Max:     stats.Max,
			Min:     stats.Min,
			Average: stats.Average,
			Sum:     stats.Sum,
			IsDiagonal: struct {
				Rotated bool `json:"rotated"`
				Q       bool `json:"Q"`
				R       bool `json:"R"`
			}{
				Rotated: stats.IsDiagonal.Rotated,
				Q:       stats.IsDiagonal.Q,
				R:       stats.IsDiagonal.R,
			},
		},
	}

	return resp, nil
}
