package utils

func ReleaseYearValidator(year int) string {
	if year < 1980 || year > 2024 {
		return "You can only add books released in between 1980 - 2024"
	} else {
		return ""
	}
}
