package entites

import "time"

type Patient struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Gender Gender        `json:"gender"`
	Birth  string        `json:"birth"`
	Height float64       `json:"height"`
	Weight float64       `json:"weight"`
	Cases  []PatientCase `json:"cases"`
}

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (p *Patient) GetAge() int {
	birth, err := time.Parse("2006-01-02", p.Birth)
	if err != nil {
		return 0
	}
	return time.Now().Year() - birth.Year()
}

func (p *Patient) GetBMI() float64 {
	return p.Weight / (p.Height * p.Height)
}
