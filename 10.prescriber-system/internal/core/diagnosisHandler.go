package core

import "github/elliot9/class10/internal/entites"

type DiagnosisHandler interface {
	Match(patient *entites.Patient, symptoms []string) bool
	Handle(patient *entites.Patient, symptoms []string) entites.Prescription
}

type Covid19DiagnosisHandler struct{}
type AttractiveDiagnosisHandler struct{}
type SleepApneaSyndromeDiagnosisHandler struct{}

func NewCovid19DiagnosisHandler() *Covid19DiagnosisHandler {
	return &Covid19DiagnosisHandler{}
}

func NewAttractiveDiagnosisHandler() *AttractiveDiagnosisHandler {
	return &AttractiveDiagnosisHandler{}
}

func NewSleepApneaSyndromeDiagnosisHandler() *SleepApneaSyndromeDiagnosisHandler {
	return &SleepApneaSyndromeDiagnosisHandler{}
}

func (d *Covid19DiagnosisHandler) Match(patient *entites.Patient, symptoms []string) bool {
	prescriptionSymptoms := map[string]bool{"Sneeze": true, "Headache": true, "Cough": true}
	for _, symptom := range symptoms {
		if prescriptionSymptoms[symptom] {
			delete(prescriptionSymptoms, symptom)
		}
	}
	return len(prescriptionSymptoms) == 0
}

func (d *Covid19DiagnosisHandler) Handle(patient *entites.Patient, symptoms []string) entites.Prescription {
	return entites.NewPrescription("清冠一號", "新冠肺炎（專業學名：COVID-19）", []string{"清冠一號"}, "將相關藥材裝入茶包裡，使用500 mL 溫、熱水沖泡悶煮1~3 分鐘後即可飲用。")
}

func (d *AttractiveDiagnosisHandler) Match(patient *entites.Patient, symptoms []string) bool {
	prescriptionSymptoms := map[string]bool{"Sneeze": true}
	for _, symptom := range symptoms {
		if prescriptionSymptoms[symptom] {
			delete(prescriptionSymptoms, symptom)
		}
	}
	return len(prescriptionSymptoms) == 0 && patient.Gender == entites.Female && patient.GetAge() == 18
}

func (d *AttractiveDiagnosisHandler) Handle(patient *entites.Patient, symptoms []string) entites.Prescription {
	return entites.NewPrescription("青春抑制劑", "有人想你了 (專業學名：Attractive)", []string{"假鬢角", "臭味"}, "把假鬢角黏在臉的兩側，讓自己異性緣差一點，自然就不會有人想妳了。")
}

func (d *SleepApneaSyndromeDiagnosisHandler) Match(patient *entites.Patient, symptoms []string) bool {
	prescriptionSymptoms := map[string]bool{"Snore": true}
	for _, symptom := range symptoms {
		if prescriptionSymptoms[symptom] {
			delete(prescriptionSymptoms, symptom)
		}
	}
	return len(prescriptionSymptoms) == 0 && patient.GetBMI() > 26.0
}

func (d *SleepApneaSyndromeDiagnosisHandler) Handle(patient *entites.Patient, symptoms []string) entites.Prescription {
	return entites.NewPrescription("打呼抑制劑", "睡眠呼吸中止症（專業學名：SleepApneaSyndrome）", []string{"一捲膠帶"}, "睡覺時，撕下兩塊膠帶，將兩塊膠帶交錯黏在關閉的嘴巴上，就不會打呼了。")
}

var _ DiagnosisHandler = &Covid19DiagnosisHandler{}
var _ DiagnosisHandler = &AttractiveDiagnosisHandler{}
var _ DiagnosisHandler = &SleepApneaSyndromeDiagnosisHandler{}
