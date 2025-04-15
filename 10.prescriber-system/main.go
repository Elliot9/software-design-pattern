package main

import (
	"github/elliot9/class10/facades"
)

var prescriberSystem *facades.PrescriberSystem

func init() {
	prescriberSystem = facades.NewPrescriberSystem()
}

func main() {
	// 裝著病患資料的 JSON 檔案名稱
	// 裝著支援潛在疾病診斷的純文字檔案名稱
	prescriberSystem.Load("tests/mockDB2.json", "tests/diagnosis.txt")

	// 欲診斷病患的名稱和他的多項症狀。
	prescription := prescriberSystem.AskDiagnosis("A123456789", "Sneeze,Headache,Cough")

	// 維護者希望能夠在外部決定在收到診斷結果時要觸發什麼行為，而不用修改模組內部程式。
	// 在完成診斷之後，我能向模組要求把此次診斷結果存到哪個檔案，並且也能選擇要存成 JSON 格式還是 CSV 格式。
	prescriberSystem.OutputPrescription(<-prescription, facades.OutPutPrescriptionStrategy_JsonPersistence, "tests/output.json")
}
