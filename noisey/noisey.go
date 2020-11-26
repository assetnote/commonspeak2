package noisey

import "regexp"

func IsNotNoisey(inputString string) bool {
	isMD5, _ := regexp.MatchString("^[0-9a-fA-F]{32}$", inputString)
	isSHA1, _ := regexp.MatchString("^[a-fA-F0-9]{40}$", inputString)
	isSHA256, _ := regexp.MatchString("^[A-Fa-f0-9]{64}$", inputString)
	isSHA512, _ := regexp.MatchString("^[A-Fa-f0-9]{128}$", inputString)
	isNoisy, _ := regexp.MatchString("[\\!(,%]", inputString)
	is100, _ := regexp.MatchString(".{100,}", inputString)
	isId, _ := regexp.MatchString("[0-9]{4,}", inputString)
	isIdLast3, _ := regexp.MatchString("[0-9]{3,}$", inputString)
	isNumberNoisy, _ := regexp.MatchString("[0-9]+[A-Z0-9]{5,}", inputString)
	isDirsDeep, _ := regexp.MatchString("\\/.*\\/.*\\/.*\\/.*\\/.*\\/.*\\/", inputString)
	isUUID, _ := regexp.MatchString("\\w{8}-\\w{4}-\\w{4}-\\w{4}-\\w{12}", inputString)
	isMultiNumLetter, _ := regexp.MatchString("[0-9]+[a-zA-Z]+[0-9]+[a-zA-Z]+[0-9]+", inputString)
	isLowValueFileType, _ := regexp.MatchString("\\.(png|jpg|jpeg|gif|svg|bmp|ttf|avif|wav|mp4|aac|ajax|css|all|)$", inputString)
	if !isMD5 && !isSHA1 && !isSHA256 && !isSHA512 && !isNoisy && !is100 &&
		!isId && !isIdLast3 && !isNumberNoisy && !isDirsDeep && !isUUID && !isMultiNumLetter && !isLowValueFileType {
		return true
	} else {
		return false
	}
}
