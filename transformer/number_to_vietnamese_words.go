package transformer

import (
	"slices"
	"strings"
)

func NumberToVietnameseWords(number int64) string {
	if number == 0 {
		return "không"
	}

	var sign string
	var num uint64

	if number < 0 {
		num = uint64(-(number + 1)) + 1
		sign = "âm"
	} else {
		num = uint64(number)
	}

	return joinWords(sign, doNumberToVietnameseWords(num, countDigits(num), true))
}

func doNumberToVietnameseWords(number uint64, numDigits int, rightMost bool) string {
	if number < 10 {
		return belowTen[number]
	}

	if number < 20 {
		return belowTwenty[number-10]
	}

	left, right, leftDigits, rightDigits, zeroPadding := decomposeNumber(number, numDigits)

	if number < 100 {
		if (right == 1 && left > 1) || (rightMost && right == 4 && left > 1) || (right == 5 && left > 0) {
			return joinWords(belowHundred[left], belowTenIrregular[right])
		}
		
		return joinWords(belowHundred[left], belowTen[right])
	}

	return joinWords(
		doNumberToVietnameseWords(left, leftDigits, false),
		scaleUnit(numDigits),
		linkingWords(numDigits, rightDigits, zeroPadding),
		doNumberToVietnameseWords(right, rightDigits, rightMost),
	)
}

func decomposeNumber(number uint64, numDigits int) (left uint64, right uint64, leftDigits, rightDigits, zeroPadding int) {
	if number < 100 {
		left = number / 10
		right = number % 10
	} else if number < 1_000 {
		left = number / 100
		right = number % 100
	} else {
		group := (numDigits - 1) / 3
		left = number / thousandPowers[group]
		right = number % thousandPowers[group]
	}

	leftDigits = countDigits(left)

	if right == 0 {
		return left, right, leftDigits, 0, 0
	}

	rightDigits = countDigits(right)
	zeroPadding = ((rightDigits+2)/3)*3 - rightDigits
	return
}

func countDigits(number uint64) int {
	var count int
	for number > 0 {
		number = number / 10
		count++
	}

	return count
}

func scaleUnit(numDigits int) string {
	return units[((numDigits + 2) / 3)]
}

func linkingWords(numDigits, rightDigits, zeroPadding int) string {
	if numDigits <= 3 {
		if rightDigits == 1 {
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
	return strings.Join(slices.DeleteFunc(words, func(s string) bool { return s == "" }), " ")
}

var belowTen = []string{"", "một", "hai", "ba", "bốn", "năm", "sáu", "bảy", "tám", "chín"}

var belowTenIrregular = []string{"", "mốt", "hai", "ba", "tư", "lăm", "sáu", "bảy", "tám", "chín"}

var belowTwenty = []string{"mười", "mười một", "mười hai", "mười ba", "mười bốn", "mười lăm", "mười sáu", "mười bảy", "mười tám", "mười chín"}

var belowHundred = []string{"", "mười", "hai mươi", "ba mươi", "bốn mươi", "năm mươi", "sáu mươi", "bảy mươi", "tám mươi", "chín mươi"}

var units = []string{"", "trăm", "nghìn", "triệu", "tỷ", "nghìn tỷ", "triệu tỷ", "tỷ tỷ"}

var thousandPowers = []uint64{1, 1_000, 1_000_000, 1_000_000_000, 1_000_000_000_000, 1_000_000_000_000_000, 1_000_000_000_000_000_000}
