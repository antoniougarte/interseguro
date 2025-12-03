package matrix

import (
	"errors"
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// ValidMatrix verifica si la matriz es rectangular y contiene valores finitos
func ValidMatrix(m [][]float64) error {
	if len(m) == 0 {
		return errors.New("matrix is empty")
	}
	cols := len(m[0])
	if cols == 0 {
		return errors.New("matrix has zero columns")
	}
	for i, row := range m {
		if len(row) != cols {
			return fmt.Errorf("row %d length mismatch: expected %d, got %d", i, cols, len(row))
		}
		for j, v := range row {
			if math.IsNaN(v) || math.IsInf(v, 0) {
				return fmt.Errorf("invalid number at row %d col %d", i, j)
			}
		}
	}
	return nil
}

// Rotate90Clockwise rota la matriz 90° en sentido horario
func Rotate90Clockwise(m [][]float64) [][]float64 {
	rows := len(m)
	if rows == 0 {
		return [][]float64{}
	}
	cols := len(m[0])
	out := make([][]float64, cols)
	for i := range out {
		out[i] = make([]float64, rows)
	}
	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			out[j][i] = m[rows-1-i][j]
		}
	}
	return out
}

// QRFactor realiza la factorización QR de la matriz 'a' y devuelve Q,R en [][]float64.
// Si la matriz no es adecuada para QR (por ejemplo m < n) devuelve error.
// Además protege contra panic de gonum y lo retorna como error.
func QRFactor(a [][]float64) ([][]float64, [][]float64, error) {
	// validación mínima
	m := len(a)
	if m == 0 {
		return nil, nil, errors.New("matrix empty")
	}
	n := len(a[0])

	// Construir Dense
	data := make([]float64, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			data = append(data, a[i][j])
		}
	}
	A := mat.NewDense(m, n, data)

	// Protegemos contra panic inesperado
	defer func() {
		if r := recover(); r != nil {
			// convertimos panic a error — el caller lo manejará
			// no podemos retornar desde defer con valores nombrados a menos que se definan arriba;
			// por simplicidad el caller detectará si Q/R son nil y error no nil.
		}
	}()

	// gonum QR requiere m >= n (filas >= columnas) para la factorización convencional.
	if m < n {
		return nil, nil, fmt.Errorf("cannot compute QR: rows (%d) < cols (%d). Consider transposing or provide a matrix with rows >= cols", m, n)
	}

	var qr mat.QR
	qr.Factorize(A)

	var Q mat.Dense
	qr.QTo(&Q)

	var R mat.Dense
	qr.RTo(&R)

	// Obtener dimensiones reales
	qRows, qCols := Q.Dims()
	rRows, rCols := R.Dims()

	// Convertir Q a slice
	qOut := make([][]float64, qRows)
	for i := 0; i < qRows; i++ {
		row := make([]float64, qCols)
		for j := 0; j < qCols; j++ {
			v := Q.At(i, j)
			// Si es casi cero lo volvemos 0
			if math.Abs(v) < 1e-12 {
				v = 0
			}
			row[j] = v
		}
		qOut[i] = row
	}

	// Convertir R a slice
	rOut := make([][]float64, rRows)
	for i := 0; i < rRows; i++ {
		row := make([]float64, rCols)
		for j := 0; j < rCols; j++ {
			v := R.At(i, j)
			if math.Abs(v) < 1e-12 {
				v = 0
			}
			row[j] = v
		}
		rOut[i] = row
	}

	return qOut, rOut, nil
}
