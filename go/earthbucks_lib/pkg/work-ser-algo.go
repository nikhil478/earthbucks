package earthbucks

type WorkSerAlgoName string

const (
	Null     WorkSerAlgoName = "null"
	Blake3   WorkSerAlgoName = "blake3"
	Blake3_2 WorkSerAlgoName = "blake3_2"
	Blake3_3 WorkSerAlgoName = "blake3_3"
)

var workSerAlgoName = map[int]WorkSerAlgoName{
	0: Null,
	1: Blake3,
	2: Blake3_2,
	3: Blake3_3,
}

var workSerAlgoNum = map[WorkSerAlgoName]int{
	Null:     0,
	Blake3:   1,
	Blake3_2: 2,
	Blake3_3: 3,
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