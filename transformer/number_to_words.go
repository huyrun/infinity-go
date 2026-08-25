package transformer

import "strings"

func NumberToWords(num int64) string {
	if num == 0 {
		return "Không"
	}
	return strings.TrimSpace(doNumberToWords(num, true))
}

func doNumberToWords(num int64, rightMost bool) string {
	if num < 10 {
		return belowTen[num]
	}

	if num < 20 {
		return belowTwenty[num-10]
	}

	if num < 100 {
		onesPlaceNum := num % 10
		tensPlaceNum := num / 10
		onesPlaceWords := doNumberToWords(onesPlaceNum, rightMost)

		if onesPlaceNum == 1 && tensPlaceNum > 1 {
			onesPlaceWords = irregular[onesPlaceWords]
		}
		if rightMost && onesPlaceNum == 4 && tensPlaceNum > 1 {
			onesPlaceWords = irregular[onesPlaceWords]
		}
		if onesPlaceNum == 5 && tensPlaceNum > 0 {
			onesPlaceWords = irregular[onesPlaceWords]
		}

		return strings.TrimSpace(belowHundred[tensPlaceNum] + " " + onesPlaceWords)
	}

	if num < 1_000 {
		mod := num % 100
		words := doNumberToWords(mod, rightMost)
		if len(words) > 0 {
			if mod < 10 {
				words = "Linh " + words
			}
		}
		return strings.TrimSpace(doNumberToWords(num/100, false) + " Trăm " + words)
	}

	if num < 1_000_000 {
		mod := num % 1_000
		words := doNumberToWords(mod, rightMost)
		if len(words) > 0 {
			if mod < 10 {
				return strings.TrimSpace(doNumberToWords(num/1_000, false) + " Nghìn Không Trăm Linh " + words)
			}
			if mod < 100 {
				return strings.TrimSpace(doNumberToWords(num/1_000, false) + " Nghìn Không Trăm " + words)
			}
		}
		return strings.TrimSpace(doNumberToWords(num/1_000, false) + " Nghìn " + words)
	}

	if num < 1_000_000_000 {
		mod := num % 1_000_000
		words := doNumberToWords(mod, rightMost)
		if len(words) > 0 {
			if mod < 10 {
				return strings.TrimSpace(doNumberToWords(num/1_000_000, false) + " Triệu Không Trăm Linh " + words)
			}
			if mod < 100 {
				return strings.TrimSpace(doNumberToWords(num/1_000_000, false) + " Triệu Không Trăm " + words)
			}
			if mod < 1_000 {
				return strings.TrimSpace(doNumberToWords(num/1_000_000, false) + " Triệu " + words)
			}
			if mod < 10_000 {
				return strings.TrimSpace(doNumberToWords(num/1_000_000, false) + " Triệu Không Trăm Linh " + words)
			}
			if mod < 100_000 {
				return strings.TrimSpace(doNumberToWords(num/1_000_000, false) + " Triệu Không Trăm " + words)
			}
		}
		return strings.TrimSpace(doNumberToWords(num/1_000_000, false) + " Triệu " + words)
	}

	//num >= 1_000_000_000
	mod := num % 1_000_000_000
	words := doNumberToWords(mod, rightMost)
	if len(words) > 0 {
		if mod < 10 {
			return strings.TrimSpace(doNumberToWords(num/1_000_000_000, false) + " Tỷ Không Trăm Linh " + words)
		}
		if mod < 100 {
			return strings.TrimSpace(doNumberToWords(num/1_000_000_000, false) + " Tỷ Không Trăm " + words)
		}
		if mod < 1_000 {
			return strings.TrimSpace(doNumberToWords(num/1_000_000_000, false) + " Tỷ " + words)
		}
		if mod < 10_000 {
			return strings.TrimSpace(doNumberToWords(num/1_000_000_000, false) + " Tỷ Không Trăm Linh " + words)
		}
		if mod < 100_000 {
			return strings.TrimSpace(doNumberToWords(num/1_000_000_000, false) + " Tỷ Không Trăm " + words)
		}
		if mod < 1_000_000 {
			return strings.TrimSpace(doNumberToWords(num/1_000_000_000, false) + " Tỷ " + words)
		}
		if mod < 10_000_000 {
			return strings.TrimSpace(doNumberToWords(num/1_000_000_000, false) + " Tỷ Không Trăm Linh " + words)
		}
		if mod < 100_000_000 {
			return strings.TrimSpace(doNumberToWords(num/1_000_000_000, false) + " Tỷ Không Trăm " + words)
		}
	}

	return strings.TrimSpace(doNumberToWords(num/1_000_000_000, false) + " Tỷ " + words)
}

var irregular = map[string]string{
	"Một": "Mốt",
	"Bốn": "Tư",
	"Năm": "Lăm",
}

var belowTen = []string{"", "Một", "Hai", "Ba", "Bốn", "Năm", "Sáu", "Bảy", "Tám", "Chín"}

var belowTwenty = []string{"Mười", "Mười Một", "Mười Hai", "Mười Ba", "Mười Bốn", "Mười Lăm", "Mười Sáu", "Mười Bảy", "Mười Tám", "Mười Chín"}

var belowHundred = []string{"", "Mười", "Hai Mươi", "Ba Mươi", "Bốn Mươi", "Năm Mươi", "Sáu Mươi", "Bảy Mươi", "Tám Mươi", "Chín Mươi"}
