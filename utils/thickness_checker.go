package utils

func ThicknessChecker(totalPage int) string {
	if totalPage < 100 {
		return "tipis"
	} else {
		return "tebal"
	}
}
