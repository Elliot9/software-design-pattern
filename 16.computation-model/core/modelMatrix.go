package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var matrixCache sync.Map

type ModelMatrix struct {
	matrices [1000][1000]float64
}

func GetInstance(modelName string) *ModelMatrix {
	if matrix, ok := matrixCache.Load(modelName); ok {
		return matrix.(*ModelMatrix)
	}

	matrix := newModelMatrix(modelName)
	actual, _ := matrixCache.LoadOrStore(modelName, matrix)
	return actual.(*ModelMatrix)
}

func newModelMatrix(modelName string) *ModelMatrix {
	m := &ModelMatrix{
		matrices: [1000][1000]float64{},
	}

	file, err := os.Open(fmt.Sprintf("assets/%s.mat", modelName))
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		values := strings.Fields(line)

		for j, val := range values {
			m.matrices[i][j], err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil
			}
		}
	}

	return m
}

func (m *ModelMatrix) GetMatrices() [1000][1000]float64 {
	return m.matrices
}
