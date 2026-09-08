package transformer

import (
	"slices"
	"strings"
)

var belowTenIrregular = map[string]string{
	"một": "mốt",
	"bốn": "tư",
	"năm": "lăm",
}

var belowTen = []string{"", "một", "hai", "ba", "bốn", "năm", "sáu", "bảy", "tám", "chín"}

var belowTwenty = []string{"mười", "mười một", "mười hai", "mười ba", "mười bốn", "mười lăm", "mười sáu", "mười bảy", "mười tám", "mười chín"}

var belowHundred = []string{"", "mười", "hai mươi", "ba mươi", "bốn mươi", "năm mươi", "sáu mươi", "bảy mươi", "tám mươi", "chín mươi"}

var units = []string{
	"", "", "",
	"trăm",
	"nghìn", "nghìn", "nghìn",
	"triệu", "triệu", "triệu",
	"tỷ", "tỷ", "tỷ",
	"nghìn tỷ", "nghìn tỷ", "nghìn tỷ",
	"triệu tỷ", "triệu tỷ", "triệu tỷ",
	"tỷ tỷ", "tỷ tỷ", "tỷ tỷ",
}

var thousandPowers = []uint64{
	1,
	1_000,
	1_000_000,
	1_000_000_000,
	1_000_000_000_000,
	1_000_000_000_000_000,
	1_000_000_000_000_000_000,
}

func NumberToVietnameseWords(num int64) string {
	if num == 0 {
		return "không"
	}

	var negative string
	var n uint64

	if num < 0 {
		n = uint64(-(num + 1)) + 1
		negative = "âm"
	} else {
		n = uint64(num)
	}

	return joinWords(negative, doNumberToVietnameseWords(n, true))
}

func doNumberToVietnameseWords(num uint64, rightMost bool) string {
	if num < 10 {
		return belowTen[num]
	}

	if num < 20 {
		return belowTwenty[num-10]
	}

	quotient, remainder, countNumDigits, countRemainderDigits, zeroPadding := divideNum(num)

	if num < 100 {
		onesPlaceWords := doNumberToVietnameseWords(remainder, rightMost)

		if remainder == 1 && quotient > 1 {
			onesPlaceWords = belowTenIrregular[onesPlaceWords]
		}
		if rightMost && remainder == 4 && quotient > 1 {
			onesPlaceWords = belowTenIrregular[onesPlaceWords]
		}
		if remainder == 5 && quotient > 0 {
			onesPlaceWords = belowTenIrregular[onesPlaceWords]
		}

		return joinWords(belowHundred[quotient], onesPlaceWords)
	}

	return joinWords(doNumberToVietnameseWords(quotient, false), scaleUnit(countNumDigits), linkingWords(countNumDigits, countRemainderDigits, zeroPadding), doNumberToVietnameseWords(remainder, rightMost))
}

func divideNum(num uint64) (quotient uint64, remainder uint64, countNumDigits int, countRemainderDigits, zeroPadding int) {
	countNumDigits = countDigits(num)

	if num < 100 {
		quotient = num / 10
		remainder = num % 10
	} else if num < 1_000 {
		quotient = num / 100
		remainder = num % 100
	} else {
		group := countNumDigits / 3
		if countNumDigits%3 == 0 {
			group--
		}

		quotient = num / thousandPowers[group]
		remainder = num % thousandPowers[group]
	}

	if remainder == 0 {
		return quotient, remainder, countNumDigits, 0, 0
	}

	countRemainderDigits = countDigits(remainder)
	zeroPadding = ((countRemainderDigits+2)/3)*3 - countRemainderDigits
	return
}

func countDigits(num uint64) int {
	var count int
	for num > 0 {
		num = num / 10
		count++
	}

	return count
}

func scaleUnit(countNumDigits int) string {
	return units[countNumDigits]
}

func linkingWords(countNumDigits, countRemainderDigits, zeroPadding int) string {
	if countNumDigits <= 3 {
		if countRemainderDigits == 1 {
			return "linh"
		}
		return ""
	}

	switch zeroPadding {
	case 1:
		return "không trăm"
	case 2:
		return "không trăm linh"
	default:
		return ""
	}
}

func joinWords(words ...string) string {
	return strings.TrimSpace(strings.Join(
		slices.DeleteFunc(words, func(s string) bool {
			return s == ""
		}),
		" ",
	))
}
