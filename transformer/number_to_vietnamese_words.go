package transformer

import (
	"slices"
	"strings"
)

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

	return joinWords(negative, doNumberToVietnameseWords(n, countDigits(n), true))
}

func doNumberToVietnameseWords(num uint64, countNumDigits int, rightMost bool) string {
	if num < 10 {
		return belowTen[num]
	}

	if num < 20 {
		return belowTwenty[num-10]
	}

	leftPart, rightPart, countLeftPartDigits, countRightPartDigits, zeroPadding := decomposeNumber(num, countNumDigits)

	if num < 100 {
		onesPlaceWords := doNumberToVietnameseWords(rightPart, countRightPartDigits, rightMost)

		if rightPart == 1 && leftPart > 1 {
			onesPlaceWords = belowTenIrregular[onesPlaceWords]
		}
		if rightMost && rightPart == 4 && leftPart > 1 {
			onesPlaceWords = belowTenIrregular[onesPlaceWords]
		}
		if rightPart == 5 && leftPart > 0 {
			onesPlaceWords = belowTenIrregular[onesPlaceWords]
		}

		return joinWords(belowHundred[leftPart], onesPlaceWords)
	}

	return joinWords(
		doNumberToVietnameseWords(leftPart, countLeftPartDigits, false),
		scaleUnit(countNumDigits),
		linkingWords(countNumDigits, countRightPartDigits, zeroPadding),
		doNumberToVietnameseWords(rightPart, countRightPartDigits, rightMost),
	)
}

func decomposeNumber(num uint64, countNumDigits int) (leftPart uint64, rightPart uint64, countLeftPartDigits, countRightPartDigits, zeroPadding int) {
	if num < 100 {
		leftPart = num / 10
		rightPart = num % 10
	} else if num < 1_000 {
		leftPart = num / 100
		rightPart = num % 100
	} else {
		group := countNumDigits / 3
		if countNumDigits%3 == 0 {
			group--
		}

		leftPart = num / thousandPowers[group]
		rightPart = num % thousandPowers[group]
	}

	countLeftPartDigits = countDigits(leftPart)

	if rightPart == 0 {
		return leftPart, rightPart, countLeftPartDigits, 0, 0
	}

	countRightPartDigits = countDigits(rightPart)
	zeroPadding = ((countRightPartDigits+2)/3)*3 - countRightPartDigits
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
	return units[((countNumDigits + 2) / 3)]
}

func linkingWords(countNumDigits, countRemainderDigits, zeroPadding int) string {
	if countNumDigits <= 3 {
		if countRemainderDigits == 1 {
			return "lẻ"
		}
		return ""
	}

	switch zeroPadding {
	case 1:
		return "không trăm"
	case 2:
		return "không trăm lẻ"
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

var belowTenIrregular = map[string]string{
	"một": "mốt",
	"bốn": "tư",
	"năm": "lăm",
}

var belowTen = []string{"", "một", "hai", "ba", "bốn", "năm", "sáu", "bảy", "tám", "chín"}

var belowTwenty = []string{"mười", "mười một", "mười hai", "mười ba", "mười bốn", "mười lăm", "mười sáu", "mười bảy", "mười tám", "mười chín"}

var belowHundred = []string{"", "mười", "hai mươi", "ba mươi", "bốn mươi", "năm mươi", "sáu mươi", "bảy mươi", "tám mươi", "chín mươi"}

var units = []string{"", "trăm", "nghìn", "triệu", "tỷ", "nghìn tỷ", "triệu tỷ", "tỷ tỷ"}

var thousandPowers = []uint64{1, 1_000, 1_000_000, 1_000_000_000, 1_000_000_000_000, 1_000_000_000_000_000, 1_000_000_000_000_000_000}
