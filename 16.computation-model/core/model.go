package core

type Model interface {
	GetName() string
	LinearTransformation(vector []float64) []float64
	GetMatrix() *ModelMatrix
}

func NewModel(name string) Model {
	return &BaseModel{
		Name: name,
	}
}

type BaseModel struct {
	Name string
}

func (m *BaseModel) GetName() string {
	return m.Name
}

func (m *BaseModel) LinearTransformation(vector []float64) []float64 {
	size := len(vector)
	result := make([]float64, size)

	for i := 0; i < size; i++ {
		var sum float64
		for n := 0; n < size; n++ {
			sum += vector[n] * m.GetMatrix().GetMatrices()[n][i]
		}
		result[i] = sum
	}

	return result
}

func (m *BaseModel) GetMatrix() *ModelMatrix {
	return GetInstance(m.Name)
}

var _ Model = &BaseModel{}
