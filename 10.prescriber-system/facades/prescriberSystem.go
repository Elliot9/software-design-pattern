package facades

import (
	"github/elliot9/class10/infra"
	"github/elliot9/class10/internal/core"
	"github/elliot9/class10/internal/entites"
	"github/elliot9/class10/internal/factory"
	"strings"
)

type OutPutPrescriptionStrategy string
type Diagnosis string

const (
	OutPutPrescriptionStrategy_JsonPersistence OutPutPrescriptionStrategy = "JsonPersistence"
	OutPutPrescriptionStrategy_CsvPersistence  OutPutPrescriptionStrategy = "CsvPersistence"
	Diagnosis_COVID19                          Diagnosis                  = "COVID-19"
	Diagnosis_Attractive                       Diagnosis                  = "Attractive"
	Diagnosis_SleepApneaSyndrome               Diagnosis                  = "SleepApneaSyndrome"
)

type PrescriberSystem struct {
	prescriber              *core.Prescriber
	diagnosisHandlerFactory *factory.DiagnosisHandlerFactory
	outputStrategyFactory   *factory.OutPutPrescriptionStrategyFactory
}

// 初始化門面, 工廠如果從外部初始化, 更符合 OCP, 不過會讓 PrescriberSystem 的建構子變得相對複雜
func NewPrescriberSystem() *PrescriberSystem {
	diagnosisHandlerFactory := factory.NewDiagnosisHandlerFactory()
	diagnosisHandlerFactory.Register(string(Diagnosis_COVID19), func() core.DiagnosisHandler {
		return core.NewCovid19DiagnosisHandler()
	})
	diagnosisHandlerFactory.Register(string(Diagnosis_Attractive), func() core.DiagnosisHandler {
		return core.NewAttractiveDiagnosisHandler()
	})
	diagnosisHandlerFactory.Register(string(Diagnosis_SleepApneaSyndrome), func() core.DiagnosisHandler {
		return core.NewSleepApneaSyndromeDiagnosisHandler()
	})

	outputStrategyFactory := factory.NewOutPutPrescriptionStrategyFactory()
	outputStrategyFactory.Register(string(OutPutPrescriptionStrategy_JsonPersistence), func() core.OutPutPrescriptionStrategy {
		return core.NewJsonPersistence()
	})
	outputStrategyFactory.Register(string(OutPutPrescriptionStrategy_CsvPersistence), func() core.OutPutPrescriptionStrategy {
		return core.NewCsvPersistence()
	})

	return &PrescriberSystem{
		prescriber:              core.NewPrescriber(),
		diagnosisHandlerFactory: diagnosisHandlerFactory,
		outputStrategyFactory:   outputStrategyFactory,
	}
}

func (p *PrescriberSystem) Load(dbPath, diagnosisPath string) {
	p.LoadPatientDatabase(dbPath)
	p.LoadDiagnosis(diagnosisPath)
}

func (p *PrescriberSystem) LoadPatientDatabase(filePath string) {
	persistence := infra.NewFilePersistence(filePath)
	p.prescriber.SetPatientDatabase(core.NewPatientDatabase(persistence))
}

func (p *PrescriberSystem) LoadDiagnosis(filePath string) {
	// reset diagnosis handlers
	p.prescriber.SetDiagnosisHandlers([]core.DiagnosisHandler{})

	diagnosisHandlers := strings.Split(infra.NewFilePersistence(filePath).Load(), "\n")
	handlers := []core.DiagnosisHandler{}
	for _, diagnosisHandler := range diagnosisHandlers {
		handlers = append(handlers, p.diagnosisHandlerFactory.Create(diagnosisHandler))
	}
	p.prescriber.SetDiagnosisHandlers(handlers)
}

func (p *PrescriberSystem) AskDiagnosis(id, symptoms string) chan entites.Prescription {
	return p.prescriber.AddDemand(id, symptoms)
}

func (p *PrescriberSystem) OutputPrescription(prescription entites.Prescription, outputStrategy OutPutPrescriptionStrategy, filePath string) {
	strategy := p.outputStrategyFactory.Create(string(outputStrategy))
	p.prescriber.PrescriptionHandle(prescription, strategy, filePath)
}
