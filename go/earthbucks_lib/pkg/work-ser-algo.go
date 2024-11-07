package earthbucks

type WorkSerAlgoName string

const (
	WorkSerAlgoNameNull     WorkSerAlgoName = "null"
	WorkSerAlgoNameBlake3   WorkSerAlgoName = "blake3"
	WorkSerAlgoNameBlake3_2 WorkSerAlgoName = "blake3_2"
	WorkSerAlgoNameBlake3_3 WorkSerAlgoName = "blake3_3"
)

var workSerAlgoName = map[int]WorkSerAlgoName{
	0: WorkSerAlgoNameNull,
	1: WorkSerAlgoNameBlake3,
	2: WorkSerAlgoNameBlake3_2,
	3: WorkSerAlgoNameBlake3_3,
}

var workSerAlgoNum = map[WorkSerAlgoName]int{
	WorkSerAlgoNameNull:     0,
	WorkSerAlgoNameBlake3:   1,
	WorkSerAlgoNameBlake3_2: 2,
	WorkSerAlgoNameBlake3_3: 3,
}

// GetWorkSerAlgoName returns the algorithm name for a given ID.
func GetWorkSerAlgoName(id int) (WorkSerAlgoName, bool) {
	algo, exists := workSerAlgoName[id]
	return algo, exists
}

// GetWorkSerAlgoNum returns the algorithm number for a given algorithm name.
func GetWorkSerAlgoNum(name WorkSerAlgoName) (int, bool) {
	num, exists := workSerAlgoNum[name]
	return num, exists
}