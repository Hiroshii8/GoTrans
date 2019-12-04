package Util

// TODO - IMPROVE THIS | stil stupid tho
// IsStringExist will check the entire array with Linear Search
func IsStringExist(source string, dataSource []string) bool {
	for _, v := range dataSource {
		if source == v {
			return true
		}
	}
	return false
}

func ValidateInput(sourceLanguage, destLanguage string, dataSource []string) bool {
	if IsStringExist(sourceLanguage, dataSource) && IsStringExist(destLanguage, dataSource) {
		return true
	}
	return false
}
