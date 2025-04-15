package core

import (
	"bytes"
	"encoding/json"
	"github/elliot9/class10/infra"
	"github/elliot9/class10/internal/entites"

	"github.com/yukithm/json2csv"
)

type OutPutPrescriptionStrategy interface {
	Execute(prescription entites.Prescription, filePath string)
}

type JsonPersistence struct{}

type CsvPersistence struct{}

func (h *JsonPersistence) Execute(prescription entites.Prescription, filePath string) {
	jsonStr, err := json.Marshal(prescription)
	if err != nil {
		return
	}
	persistence := infra.NewFilePersistence(filePath)
	persistence.Save(string(jsonStr))
}

func (h *CsvPersistence) Execute(prescription entites.Prescription, filePath string) {
	b := &bytes.Buffer{}
	wr := json2csv.NewCSVWriter(b)

	jsonStr, err := json.Marshal(prescription)
	if err != nil {
		return
	}

	obj, err := json2obj(string(jsonStr))
	if err != nil {
		return
	}

	csvStr, err := json2csv.JSON2CSV(obj)
	if err != nil {
		return
	}

	err = wr.WriteCSV(csvStr)
	if err != nil {
		return
	}

	wr.Flush()
	persistence := infra.NewFilePersistence(filePath)
	persistence.Save(b.String())
}

func NewJsonPersistence() *JsonPersistence {
	return &JsonPersistence{}
}

func NewCsvPersistence() *CsvPersistence {
	return &CsvPersistence{}
}

func json2obj(jsonstr string) (interface{}, error) {
	r := bytes.NewReader([]byte(jsonstr))
	d := json.NewDecoder(r)
	d.UseNumber()
	var obj interface{}
	if err := d.Decode(&obj); err != nil {
		return nil, err
	}
	return obj, nil
}

var _ OutPutPrescriptionStrategy = &JsonPersistence{}
var _ OutPutPrescriptionStrategy = &CsvPersistence{}
