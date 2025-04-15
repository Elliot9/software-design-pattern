package main

import (
	"fmt"
	"github/elliot9/class16/core"
	"sync"
)

func main() {
	mockRaceConditionMatrix([]string{"Shrinking", "Reflection", "Scaling"})
}

func mockRaceConditionMatrix(modelNames []string) {
	wg := sync.WaitGroup{}

	for _, modelName := range modelNames {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				models := core.NewModels()

				m1 := models.CreateModel(modelName)
				m2 := core.NewModel(modelName)
				result := m1.GetMatrix() == m2.GetMatrix()
				fmt.Println(result)
			}
		}()
	}

	wg.Wait()
}
