package core

import (
	"encoding/json"
	"fmt"
	"github/elliot9/class10/infra"
	"github/elliot9/class10/internal/entites"
)

type PatientDatabase struct {
	Patients    map[string]*entites.Patient `json:"patients"`
	persistence infra.Persistence
}

func NewPatientDatabase(persistence infra.Persistence) *PatientDatabase {
	jsonStr := persistence.Load()
	db := &PatientDatabase{
		Patients:    map[string]*entites.Patient{},
		persistence: persistence,
	}

	err := json.Unmarshal([]byte(jsonStr), db)
	if err != nil {
		fmt.Println("解析 JSON 失敗:", err)
		return nil
	}

	return db
}

func (p *PatientDatabase) GetPatient(id string) *entites.Patient {
	return p.Patients[id]
}

func (p *PatientDatabase) AddPatientCase(patient *entites.Patient, patientCase entites.PatientCase) {
	patient.Cases = append(patient.Cases, patientCase)
	jsonStr, _ := json.Marshal(p)
	p.persistence.Replace(string(jsonStr))
}
