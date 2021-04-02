package util

type StrReplaceStruct struct {
	CapitalLetter    int `json:"capital_letter"`
	LowercaseLetters int `json:"lowercase_letters"`
	Number           int `json:"number"`
	OtherString      int `json:"other_string"`
}

func StrReplaceAllString(s2 string) (strReplace StrReplaceStruct) {
	for i := strReplace.OtherString; i < len(s2); i++ {
		switch {
		case 64 < s2[i] && s2[i] < 91:
			strReplace.CapitalLetter += 1
		case 96 < s2[i] && s2[i] < 123:
			strReplace.LowercaseLetters += 1
		case 47 < s2[i] && s2[i] < 58:
			strReplace.Number += 1
		default:
			strReplace.OtherString += 1
		}
	}
	return strReplace
}
