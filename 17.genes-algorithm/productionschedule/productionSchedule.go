package productionschedule

type ProductionSchedule struct {
	Products               map[string]*Product
	WorkerCount            int
	MachineCount           int
	ProductionRequirements []ProductionRequirement
}

var ProductionScheduleInstance *ProductionSchedule

func init() {
	ProductionScheduleInstance = NewProductionSchedule()
}

func NewProductionSchedule() *ProductionSchedule {
	return &ProductionSchedule{
		Products:     make(map[string]*Product),
		WorkerCount:  0,
		MachineCount: 0,
	}
}

func (p *ProductionSchedule) AddProduct(name string, costTime int) {
	p.Products[name] = NewProduct(name, costTime)
}

func (p *ProductionSchedule) SetWorkerCount(count int) {
	p.WorkerCount = count
}

func (p *ProductionSchedule) SetMachineCount(count int) {
	p.MachineCount = count
}

func (p *ProductionSchedule) AddProductionRequirement(productName string, requiredNumber int) {
	p.ProductionRequirements = append(p.ProductionRequirements, NewProductionRequirement(p.Products[productName], requiredNumber))
}
