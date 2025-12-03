package model

type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type MatrixResponse struct {
	Rotated [][]float64 `json:"rotated"`
	Q       [][]float64 `json:"Q"`
	R       [][]float64 `json:"R"`
	Stats   *Statistics `json:"stats,omitempty"`
}

type Statistics struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Average    float64 `json:"average"`
	Sum        float64 `json:"sum"`
	IsDiagonal struct {
		Rotated bool `json:"rotated"`
		Q       bool `json:"Q"`
		R       bool `json:"R"`
	} `json:"isDiagonal"`
}
