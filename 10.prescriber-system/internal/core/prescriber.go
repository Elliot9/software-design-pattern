package core

import (
	"github/elliot9/class10/internal/entites"
	"strings"
	"time"
)

const (
	DiagnosisRequestQueueSize = 100
)

// 診斷請求
type DiagnosisRequest struct {
	PatientID string
	Symptoms  []string
	Result    chan entites.Prescription
}

type Prescriber struct {
	db                *PatientDatabase
	diagnosisHandlers []DiagnosisHandler
	demand            chan DiagnosisRequest
	command           OutPutPrescriptionStrategy
}

func NewPrescriber() *Prescriber {
	p := &Prescriber{
		db:                nil,
		diagnosisHandlers: nil,
		demand:            make(chan DiagnosisRequest, DiagnosisRequestQueueSize),
		command:           nil,
	}

	go p.processQueue()
	return p
}

func (p *Prescriber) SetPatientDatabase(db *PatientDatabase) {
	p.db = db
}

func (p *Prescriber) SetDiagnosisHandlers(handlers []DiagnosisHandler) {
	p.diagnosisHandlers = handlers
}

func (p *Prescriber) processQueue() {
	for req := range p.demand {
		go func(req DiagnosisRequest) {
			for _, handler := range p.diagnosisHandlers {
				patient := p.db.GetPatient(req.PatientID)
				if patient == nil {
					req.Result <- entites.Prescription{}
					continue
				}

				if handler.Match(patient, req.Symptoms) {
					prescription := handler.Handle(patient, req.Symptoms)
					// 診斷完畢之後，用戶必須透過處方診斷系統將此診斷結果儲存成該病患的一筆病例
					newPatientCase := entites.PatientCase{
						Symptoms:     req.Symptoms,
						Prescription: prescription,
						CaseTime:     time.Now(),
					}
					p.db.AddPatientCase(patient, newPatientCase)
					req.Result <- prescription
				}
			}

			time.Sleep(time.Second * 3)
			close(req.Result)
		}(req)
	}
}

func (p *Prescriber) AddDemand(patientId, symptoms string) chan entites.Prescription {
	resultChan := make(chan entites.Prescription, 1)
	request := DiagnosisRequest{
		PatientID: patientId,
		Symptoms:  strings.Split(symptoms, ","),
		Result:    resultChan,
	}
	p.demand <- request
	return resultChan
}

func (p *Prescriber) PrescriptionHandle(prescription entites.Prescription, outPut OutPutPrescriptionStrategy, filePath string) {
	outPut.Execute(prescription, filePath)
}
