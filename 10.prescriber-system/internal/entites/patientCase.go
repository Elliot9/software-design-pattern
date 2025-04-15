package entites

import "time"

type PatientCase struct {
	Symptoms     []string     `json:"symptoms"`
	Prescription Prescription `json:"prescription"`
	CaseTime     time.Time    `json:"case_time"`
}

func NewPatientCase(symptoms []string, prescription Prescription, caseTime time.Time) PatientCase {
	return PatientCase{
		Symptoms:     symptoms,
		Prescription: prescription,
		CaseTime:     caseTime,
	}
}
