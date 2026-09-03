package transformer

import (
	"strings"
)

func NumberToVietnameseWords(num int64) string {
	if num == 0 {
		return "không"
	}

	s := strings.Builder{}
	var n uint64

	if num < 0 {
		n = uint64(-(num + 1)) + 1
		s.WriteString("âm ")
	} else {
		n = uint64(num)
	}

	s.WriteString(doNumberToVietnameseWords(n, true))

	return strings.TrimSpace(s.String())
}

func doNumberToVietnameseWords(num uint64, rightMost bool) string {
	if num < 10 {
		return belowTen[num]
	}

	if num < 20 {
		return belowTwenty[num-10]
	}

	if num < 100 {
		onesPlaceNum := num % 10
		tensPlaceNum := num / 10
		onesPlaceWords := doNumberToVietnameseWords(onesPlaceNum, rightMost)

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
		words := doNumberToVietnameseWords(mod, rightMost)
		if len(words) > 0 {
			if mod < 10 {
				words = "linh " + words
			}
		}
		return strings.TrimSpace(doNumberToVietnameseWords(num/100, false) + " trăm " + words)
	}

	if num < 1_000_000 {
		mod := num % 1_000
		words := doNumberToVietnameseWords(mod, rightMost)
		if len(words) > 0 {
			if mod < 10 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000, false) + " nghìn không trăm linh " + words)
			}
			if mod < 100 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000, false) + " nghìn không trăm " + words)
			}
		}
		return strings.TrimSpace(doNumberToVietnameseWords(num/1_000, false) + " nghìn " + words)
	}

	if num < 1_000_000_000 {
		mod := num % 1_000_000
		words := doNumberToVietnameseWords(mod, rightMost)
		if len(words) > 0 {
			if mod < 10 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000, false) + " triệu không trăm linh " + words)
			}
			if mod < 100 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000, false) + " triệu không trăm " + words)
			}
			if mod < 1_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000, false) + " triệu " + words)
			}
			if mod < 10_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000, false) + " triệu không trăm linh " + words)
			}
			if mod < 100_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000, false) + " triệu không trăm " + words)
			}
		}
		return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000, false) + " triệu " + words)
	}

	if num < 1_000_000_000_000 {
		mod := num % 1_000_000_000
		words := doNumberToVietnameseWords(mod, rightMost)
		if len(words) > 0 {
			if mod < 10 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000, false) + " tỷ không trăm linh " + words)
			}
			if mod < 100 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000, false) + " tỷ không trăm " + words)
			}
			if mod < 1_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000, false) + " tỷ " + words)
			}
			if mod < 10_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000, false) + " tỷ không trăm linh " + words)
			}
			if mod < 100_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000, false) + " tỷ không trăm " + words)
			}
			if mod < 1_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000, false) + " tỷ " + words)
			}
			if mod < 10_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000, false) + " tỷ không trăm linh " + words)
			}
			if mod < 100_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000, false) + " tỷ không trăm " + words)
			}
		}

		return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000, false) + " tỷ " + words)
	}

	if num < 1_000_000_000_000_000 {
		mod := num % 1_000_000_000_000
		words := doNumberToVietnameseWords(mod, rightMost)
		if len(words) > 0 {
			if mod < 10 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ không trăm linh " + words)
			}
			if mod < 100 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ không trăm " + words)
			}
			if mod < 1_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ " + words)
			}
			if mod < 10_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ không trăm linh " + words)
			}
			if mod < 100_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ không trăm " + words)
			}
			if mod < 1_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ " + words)
			}
			if mod < 10_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ không trăm linh " + words)
			}
			if mod < 100_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ không trăm " + words)
			}
			if mod < 1_000_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ " + words)
			}
			if mod < 10_000_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ không trăm linh " + words)
			}
			if mod < 100_000_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ không trăm " + words)
			}
		}

		return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000, false) + " nghìn tỷ " + words)
	}

	if num < 1_000_000_000_000_000_000 {
		mod := num % 1_000_000_000_000_000
		words := doNumberToVietnameseWords(mod, rightMost)
		if len(words) > 0 {
			if mod < 10 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ không trăm linh " + words)
			}
			if mod < 100 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ không trăm " + words)
			}
			if mod < 1_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ " + words)
			}
			if mod < 10_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ không trăm linh " + words)
			}
			if mod < 100_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ không trăm " + words)
			}
			if mod < 1_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ " + words)
			}
			if mod < 10_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ không trăm linh " + words)
			}
			if mod < 100_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ không trăm " + words)
			}
			if mod < 1_000_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ " + words)
			}
			if mod < 10_000_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ không trăm linh " + words)
			}
			if mod < 100_000_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ không trăm " + words)
			}
			if mod < 1_000_000_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ " + words)
			}
			if mod < 10_000_000_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ không trăm linh " + words)
			}
			if mod < 100_000_000_000_000 {
				return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ không trăm " + words)
			}
		}

		return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " triệu tỷ " + words)
	}

	// num >= 1_000_000_000_000_000_000_000
	mod := num % 1_000_000_000_000_000_000
	words := doNumberToVietnameseWords(mod, rightMost)
	if len(words) > 0 {
		if mod < 10 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm linh " + words)
		}
		if mod < 100 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm " + words)
		}
		if mod < 1_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ " + words)
		}
		if mod < 10_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm linh " + words)
		}
		if mod < 100_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm " + words)
		}
		if mod < 1_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ " + words)
		}
		if mod < 10_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm linh " + words)
		}
		if mod < 100_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000, false) + " tỷ tỷ không trăm " + words)
		}
		if mod < 1_000_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ " + words)
		}
		if mod < 10_000_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm linh " + words)
		}
		if mod < 100_000_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm " + words)
		}
		if mod < 1_000_000_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ " + words)
		}
		if mod < 10_000_000_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm linh " + words)
		}
		if mod < 100_000_000_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm " + words)
		}
		if mod < 1_000_000_000_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ " + words)
		}
		if mod < 10_000_000_000_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm linh " + words)
		}
		if mod < 100_000_000_000_000_000 {
			return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ không trăm " + words)
		}
	}

	return strings.TrimSpace(doNumberToVietnameseWords(num/1_000_000_000_000_000_000, false) + " tỷ tỷ " + words)
}

var irregular = map[string]string{
	"một": "mốt",
	"bốn": "tư",
	"năm": "lăm",
}

var belowTen = []string{"", "một", "hai", "ba", "bốn", "năm", "sáu", "bảy", "tám", "chín"}

var belowTwenty = []string{"mười", "mười một", "mười hai", "mười ba", "mười bốn", "mười lăm", "mười sáu", "mười bảy", "mười tám", "mười chín"}

var belowHundred = []string{"", "mười", "hai mươi", "ba mươi", "bốn mươi", "năm mươi", "sáu mươi", "bảy mươi", "tám mươi", "chín mươi"}
