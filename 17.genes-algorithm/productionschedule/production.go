package productionschedule

import geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"

type Production struct {
	Product   *Product
	Number    int
	WorkerID  int
	MachineID int
	StartTime int
	EndTime   int
}

func (p *Production) GetValue() any {
	return p
}

var _ geneticalgorithm.Genes = &Production{}

func NewProduction(product *Product, number int, workerID int, machineID int) *Production {
	return &Production{
		Product:   product,
		Number:    number,
		WorkerID:  workerID,
		MachineID: machineID,
	}
}

func (p *Production) SetTime(startTime int) {
	p.StartTime = startTime
	p.EndTime = startTime + p.Number*p.Product.CostTime
}
