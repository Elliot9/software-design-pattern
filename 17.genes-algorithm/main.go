package main

import (
	"fmt"
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"github/elliot9/class17/pkg/geneticAlgorithm/crossover"
	"github/elliot9/class17/pkg/geneticAlgorithm/mutation"
	"github/elliot9/class17/pkg/geneticAlgorithm/selection"
	"github/elliot9/class17/productionschedule"
	"github/elliot9/class17/shoppingRecommand"
)

const (
	// 最大迭代次數
	MaxNumberOfIterations = 10000
)

func main() {
	fmt.Println("基因演算法套件測試")
	fmt.Println("====================")

	fmt.Println("\n測試情境 1：工廠生產排程...")
	mockProductionSchedule()

	fmt.Println("\n測試情境 2：購物網站推薦")
	mockShoppingRecommand()
}

func mockProductionSchedule() {
	ps := productionschedule.ProductionScheduleInstance

	// 初始化產品
	ps.AddProduct("A", 2)
	ps.AddProduct("B", 4)
	ps.AddProduct("C", 6)

	// 工廠資源
	ps.SetWorkerCount(4)
	ps.SetMachineCount(2)

	// 客戶需求
	ps.AddProductionRequirement("A", 100)
	ps.AddProductionRequirement("B", 200)
	ps.AddProductionRequirement("C", 300)

	// 終止條件
	terminationCondition := productionschedule.NewTerminationCondition()

	// 配適度
	fitness := productionschedule.NewFitness()

	// 個體工廠
	individualFactory := productionschedule.NewIndividualFactory()

	// 基因工廠
	genesFactory := productionschedule.NewGenesFactory()

	ga := geneticalgorithm.GeneticAlgorithm{
		MaxNumberOfIterations: MaxNumberOfIterations,
		TerminationCondition:  terminationCondition,
		Selection:             selection.NewTournament(fitness, 4),
		Crossover:             crossover.NewUniformCrossover(individualFactory),
		Mutation:              mutation.NewRandomReplacement(individualFactory, genesFactory),
		Fitness:               fitness,
	}

	result := ga.Optimize(getMockSchedules()).(*productionschedule.Schedule)

	// 結果顯示
	for _, production := range result.Productions {
		fmt.Printf("product: %v, number: %v, worker: %v, machine: %v, startTime: %v, endTime: %v\n", production.Product.Name, production.Number, production.WorkerID, production.MachineID, production.StartTime, production.EndTime)
	}

	fmt.Printf("total time: %v, fitness: %v, satisfied: %v\n", result.GetTotalTime(), fitness.CalculateFitness(result), terminationCondition.IsSatisfied(result))
}

func getMockSchedules() geneticalgorithm.Population {
	// 初始化族群
	s1 := productionschedule.NewSchedule()
	s1.AppendProduction("A", 100, 1, 1)
	s1.AppendProduction("B", 100, 2, 0)
	s1.AppendProduction("C", 100, 1, 1)
	s1.AppendProduction("C", 100, 2, 0)
	s1.AppendProduction("A", 50, 1, 1)
	s1.AppendProduction("B", 50, 2, 0)

	s2 := productionschedule.NewSchedule()
	s2.AppendProduction("A", 100, 1, 0)
	s2.AppendProduction("B", 100, 2, 1)
	s2.AppendProduction("C", 100, 1, 1)
	s2.AppendProduction("A", 100, 1, 0)
	s2.AppendProduction("B", 100, 2, 1)
	s2.AppendProduction("C", 100, 1, 1)

	s3 := productionschedule.NewSchedule()
	s3.AppendProduction("A", 120, 1, 1)
	s3.AppendProduction("C", 300, 2, 0)
	s3.AppendProduction("C", 600, 1, 1)
	s3.AppendProduction("C", 1, 2, 0)
	s3.AppendProduction("C", 10, 1, 1)
	s3.AppendProduction("C", 30, 2, 0)

	s4 := productionschedule.NewSchedule()
	s4.AppendProduction("A", 50, 1, 1)
	s4.AppendProduction("B", 100, 2, 0)
	s4.AppendProduction("C", 150, 1, 1)
	s4.AppendProduction("A", 50, 1, 1)
	s4.AppendProduction("B", 100, 2, 0)
	s4.AppendProduction("C", 50, 1, 1)

	population := []geneticalgorithm.Individual{
		s1, s2, s3, s4,
	}

	return geneticalgorithm.CreatePopulation(population)
}

func mockShoppingRecommand() {
	sr := shoppingRecommand.ShoppingRecommandInstance

	// 初始化產品
	sr.AddProduct(100, 2000, "A")
	sr.AddProduct(200, 3000, "A")
	sr.AddProduct(150, 5000, "B")
	sr.AddProduct(300, 4000, "B")
	sr.AddProduct(180, 6000, "C")
	sr.AddProduct(250, 7000, "C")

	// 設定客戶偏好
	sr.SetCustomerPreference("A", 0.8)
	sr.SetCustomerPreference("B", 0.6)
	sr.SetCustomerPreference("C", 0.2)

	// 設定客戶預算
	sr.SetCustomerBudget(700)

	// 設定客戶容量
	sr.SetCustomerCapacity(15000)

	// 終止條件
	terminationCondition := shoppingRecommand.NewTerminationCondition()

	// 配適度
	fitness := shoppingRecommand.NewFitness()

	// 個體工廠
	individualFactory := shoppingRecommand.NewIndividualFactory()

	// 基因工廠
	genesFactory := shoppingRecommand.NewGenesFactory()

	ga := geneticalgorithm.GeneticAlgorithm{
		MaxNumberOfIterations: MaxNumberOfIterations,
		TerminationCondition:  terminationCondition,
		Selection:             selection.NewRank(fitness),
		Crossover:             crossover.NewUniformCrossover(individualFactory),
		Mutation:              mutation.NewRandomReplacement(individualFactory, genesFactory),
		Fitness:               fitness,
	}

	result := ga.Optimize(getMockRecommendations()).(*shoppingRecommand.Recommendation)

	// 結果顯示
	for _, item := range result.Items {
		fmt.Printf("product: %v, number: %v\n", item.Product.ID, item.Quantity)
	}

	fmt.Printf("total price: %v, fitness: %v, over limit: %v\n", result.GetTotalPrice(), fitness.CalculateFitness(result), terminationCondition.IsOverLimit(result))
}

func getMockRecommendations() geneticalgorithm.Population {
	// 初始化族群
	r1 := shoppingRecommand.NewRecommendation()
	r1.AddRecommendItem(1, 1)
	r1.AddRecommendItem(4, 1)

	r2 := shoppingRecommand.NewRecommendation()
	r2.AddRecommendItem(2, 1)
	r2.AddRecommendItem(5, 1)

	r3 := shoppingRecommand.NewRecommendation()
	r3.AddRecommendItem(3, 1)
	r3.AddRecommendItem(6, 1)

	r4 := shoppingRecommand.NewRecommendation()
	r4.AddRecommendItem(1, 1)
	r4.AddRecommendItem(2, 1)

	r5 := shoppingRecommand.NewRecommendation()
	r5.AddRecommendItem(5, 1)
	r5.AddRecommendItem(6, 1)

	r6 := shoppingRecommand.NewRecommendation()
	r6.AddRecommendItem(6, 2)

	population := []geneticalgorithm.Individual{
		r1, r2, r3, r4, r5, r6,
	}

	return geneticalgorithm.CreatePopulation(population)
}
