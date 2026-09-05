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

var numberUnit = []string{
	"", "", "",
	"trăm",
	"nghìn", "nghìn", "nghìn",
	"triệu", "triệu", "triệu",
	"tỷ", "tỷ", "tỷ",
	"nghìn tỷ", "nghìn tỷ", "nghìn tỷ",
	"triệu tỷ", "triệu tỷ", "triệu tỷ",
	"tỷ tỷ", "tỷ tỷ", "tỷ tỷ",
}

var linkingWords = []string{"", "không trăm linh", "không trăm"}

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

	quotient, remainder := divideNum(num)

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

	return joinWords(doNumberToVietnameseWords(quotient, false), middleWords(num, remainder), doNumberToVietnameseWords(remainder, rightMost))
}

func middleWords(num, remainder uint64) string {
	if num < 1_000 {
		var words string
		if remainder < 10 && remainder > 0 {
			words = "linh"
		}
		return joinWords("trăm", words)
	}

	var countNumDigits int
	for num > 0 {
		num = num / 10
		countNumDigits++
	}

	var countRemainderDigits int
	for remainder > 0 {
		remainder = remainder / 10
		countRemainderDigits++
	}

	return joinWords(numberUnit[countNumDigits], linkingWords[countRemainderDigits%3])
}

func divideNum(num uint64) (uint64, uint64) {
	if num < 100 {
		return num / 10, num % 10
	}

	if num < 1_000 {
		return num / 100, num % 100
	}

	var countNumDigits int
	temp := num
	for temp > 0 {
		temp = temp / 10
		countNumDigits++
	}

	group := countNumDigits / 3
	if countNumDigits%3 == 0 {
		group--
	}

	return num / thousandPowers[group], num % thousandPowers[group]
}

func joinWords(words ...string) string {
	return strings.TrimSpace(strings.Join(
		slices.DeleteFunc(words, func(s string) bool {
			return s == ""
		}),
		" ",
	))
}
