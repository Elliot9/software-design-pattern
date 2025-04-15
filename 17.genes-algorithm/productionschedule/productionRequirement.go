package productionschedule

type ProductionRequirement struct {
	Product        *Product
	RequiredNumber int
}

func NewProductionRequirement(product *Product, requiredNumber int) ProductionRequirement {
	return ProductionRequirement{
		Product:        product,
		RequiredNumber: requiredNumber,
	}
}
