package earthbucks

type WorkParAlgoName string 

const (
	WorkParAlgoNameNull = "null"
	WorkParAlgoNameAlgo1627 = "algo1627"
)

var workParAlgoNum = map[WorkParAlgoName]int{
	WorkParAlgoNameNull:    0,
	WorkParAlgoNameAlgo1627: 1,
}

var workParAlgoName = map[int]WorkParAlgoName{
	0: WorkParAlgoNameNull,
	1: WorkParAlgoNameAlgo1627,
}

func GetWorkParAlgoName(id int) (WorkParAlgoName, bool) {
	algo, exists := workParAlgoName[id]
	return algo, exists
}

func GetWorkParAlgoNum(name WorkParAlgoName) (int, bool) {
	num, exists := workParAlgoNum[name]
	return num, exists
}