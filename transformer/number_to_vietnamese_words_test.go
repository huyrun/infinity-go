package transformer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		name string
		num  int64
		word string
	}{
		// 0 - 99
		{name: "0", num: 0, word: "Không"},
		{name: "1", num: 1, word: "Một"},
		{name: "4", num: 4, word: "Bốn"},
		{name: "5", num: 5, word: "Năm"},
		{name: "9", num: 9, word: "Chín"},

		{name: "10", num: 10, word: "Mười"},
		{name: "11", num: 11, word: "Mười Một"},
		{name: "14", num: 14, word: "Mười Bốn"},
		{name: "15", num: 15, word: "Mười Lăm"},
		{name: "19", num: 19, word: "Mười Chín"},

		{name: "20", num: 20, word: "Hai Mươi"},
		{name: "21", num: 21, word: "Hai Mươi Mốt"},
		{name: "24", num: 24, word: "Hai Mươi Tư"},
		{name: "25", num: 25, word: "Hai Mươi Lăm"},
		{name: "29", num: 29, word: "Hai Mươi Chín"},

		// 100 - 999
		{name: "100", num: 100, word: "Một Trăm"},
		{name: "101", num: 101, word: "Một Trăm Linh Một"},
		{name: "104", num: 104, word: "Một Trăm Linh Bốn"},
		{name: "105", num: 105, word: "Một Trăm Linh Năm"},
		{name: "110", num: 110, word: "Một Trăm Mười"},
		{name: "111", num: 111, word: "Một Trăm Mười Một"},
		{name: "114", num: 114, word: "Một Trăm Mười Bốn"},
		{name: "115", num: 115, word: "Một Trăm Mười Lăm"},
		{name: "120", num: 120, word: "Một Trăm Hai Mươi"},
		{name: "121", num: 121, word: "Một Trăm Hai Mươi Mốt"},
		{name: "124", num: 124, word: "Một Trăm Hai Mươi Tư"},
		{name: "125", num: 125, word: "Một Trăm Hai Mươi Lăm"},
		{name: "140", num: 140, word: "Một Trăm Bốn Mươi"},
		{name: "141", num: 141, word: "Một Trăm Bốn Mươi Mốt"},
		{name: "144", num: 144, word: "Một Trăm Bốn Mươi Tư"},
		{name: "145", num: 145, word: "Một Trăm Bốn Mươi Lăm"},
		{name: "150", num: 150, word: "Một Trăm Năm Mươi"},
		{name: "151", num: 151, word: "Một Trăm Năm Mươi Mốt"},
		{name: "154", num: 154, word: "Một Trăm Năm Mươi Tư"},
		{name: "155", num: 155, word: "Một Trăm Năm Mươi Lăm"},
		{name: "500", num: 500, word: "Năm Trăm"},
		{name: "501", num: 501, word: "Năm Trăm Linh Một"},
		{name: "504", num: 504, word: "Năm Trăm Linh Bốn"},
		{name: "505", num: 505, word: "Năm Trăm Linh Năm"},
		{name: "999", num: 999, word: "Chín Trăm Chín Mươi Chín"},

		// Thousands
		{name: "1_000", num: 1_000, word: "Một Nghìn"},
		{name: "1_001", num: 1_001, word: "Một Nghìn Không Trăm Linh Một"},
		{name: "1_010", num: 1_010, word: "Một Nghìn Không Trăm Mười"},
		{name: "1_011", num: 1_011, word: "Một Nghìn Không Trăm Mười Một"},
		{name: "1_100", num: 1_100, word: "Một Nghìn Một Trăm"},
		{name: "1_111", num: 1_111, word: "Một Nghìn Một Trăm Mười Một"},

		{name: "2_005", num: 2_005, word: "Hai Nghìn Không Trăm Linh Năm"},
		{name: "2_015", num: 2_015, word: "Hai Nghìn Không Trăm Mười Lăm"},
		{name: "2_021", num: 2_021, word: "Hai Nghìn Không Trăm Hai Mươi Mốt"},

		{name: "10_000", num: 10_000, word: "Mười Nghìn"},
		{name: "10_001", num: 10_001, word: "Mười Nghìn Không Trăm Linh Một"},
		{name: "10_010", num: 10_010, word: "Mười Nghìn Không Trăm Mười"},
		{name: "10_011", num: 10_011, word: "Mười Nghìn Không Trăm Mười Một"},
		{name: "10_100", num: 10_100, word: "Mười Nghìn Một Trăm"},
		{name: "11_000", num: 11_000, word: "Mười Một Nghìn"},

		{name: "20_005", num: 20_005, word: "Hai Mươi Nghìn Không Trăm Linh Năm"},
		{name: "20_015", num: 20_015, word: "Hai Mươi Nghìn Không Trăm Mười Lăm"},
		{name: "20_105", num: 20_105, word: "Hai Mươi Nghìn Một Trăm Linh Năm"},
		{name: "201_005", num: 201_005, word: "Hai Trăm Linh Một Nghìn Không Trăm Linh Năm"},

		{name: "100_000", num: 100_000, word: "Một Trăm Nghìn"},
		{name: "100_005", num: 100_005, word: "Một Trăm Nghìn Không Trăm Linh Năm"},
		{name: "100_011", num: 100_011, word: "Một Trăm Nghìn Không Trăm Mười Một"},
		{name: "100_100", num: 100_100, word: "Một Trăm Nghìn Một Trăm"},
		{name: "101_000", num: 101_000, word: "Một Trăm Linh Một Nghìn"},
		{name: "101_005", num: 101_005, word: "Một Trăm Linh Một Nghìn Không Trăm Linh Năm"},
		{name: "110_000", num: 110_000, word: "Một Trăm Mười Nghìn"},
		{name: "999_999", num: 999_999, word: "Chín Trăm Chín Mươi Chín Nghìn Chín Trăm Chín Mươi Chín"},

		// Millions
		{name: "1_000_000", num: 1_000_000, word: "Một Triệu"},
		{name: "1_000_001", num: 1_000_001, word: "Một Triệu Không Trăm Linh Một"},
		{name: "1_000_010", num: 1_000_010, word: "Một Triệu Không Trăm Mười"},
		{name: "1_000_011", num: 1_000_011, word: "Một Triệu Không Trăm Mười Một"},

		{name: "1_000_100", num: 1_000_100, word: "Một Triệu Một Trăm"},
		{name: "1_001_000", num: 1_001_000, word: "Một Triệu Không Trăm Linh Một Nghìn"},
		{name: "1_010_000", num: 1_010_000, word: "Một Triệu Không Trăm Mười Nghìn"},
		{name: "1_100_000", num: 1_100_000, word: "Một Triệu Một Trăm Nghìn"},

		{name: "2_000_005", num: 2_000_005, word: "Hai Triệu Không Trăm Linh Năm"},
		{name: "2_001_005", num: 2_001_005, word: "Hai Triệu Không Trăm Linh Một Nghìn Không Trăm Linh Năm"},
		{name: "20_000_005", num: 20_000_005, word: "Hai Mươi Triệu Không Trăm Linh Năm"},
		{name: "20_010_005", num: 20_010_005, word: "Hai Mươi Triệu Không Trăm Mười Nghìn Không Trăm Linh Năm"},
		{name: "200_000_005", num: 200_000_005, word: "Hai Trăm Triệu Không Trăm Linh Năm"},

		{name: "1_234_567", num: 1_234_567, word: "Một Triệu Hai Trăm Ba Mươi Bốn Nghìn Năm Trăm Sáu Mươi Bảy"},
		{name: "999_999_999", num: 999_999_999, word: "Chín Trăm Chín Mươi Chín Triệu Chín Trăm Chín Mươi Chín Nghìn Chín Trăm Chín Mươi Chín"},

		// Billions
		{name: "1_000_000_000", num: 1_000_000_000, word: "Một Tỷ"},
		{name: "1_000_000_001", num: 1_000_000_001, word: "Một Tỷ Không Trăm Linh Một"},
		{name: "1_000_000_004", num: 1_000_000_004, word: "Một Tỷ Không Trăm Linh Bốn"},
		{name: "1_000_000_005", num: 1_000_000_005, word: "Một Tỷ Không Trăm Linh Năm"},
		{name: "1_000_000_010", num: 1_000_000_010, word: "Một Tỷ Không Trăm Mười"},
		{name: "1_000_000_011", num: 1_000_000_011, word: "Một Tỷ Không Trăm Mười Một"},
		{name: "1_000_000_014", num: 1_000_000_014, word: "Một Tỷ Không Trăm Mười Bốn"},
		{name: "1_000_000_015", num: 1_000_000_015, word: "Một Tỷ Không Trăm Mười Lăm"},
		{name: "1_000_000_020", num: 1_000_000_020, word: "Một Tỷ Không Trăm Hai Mươi"},
		{name: "1_000_000_021", num: 1_000_000_021, word: "Một Tỷ Không Trăm Hai Mươi Mốt"},
		{name: "1_000_000_024", num: 1_000_000_024, word: "Một Tỷ Không Trăm Hai Mươi Tư"},
		{name: "1_000_000_025", num: 1_000_000_025, word: "Một Tỷ Không Trăm Hai Mươi Lăm"},

		{name: "1_000_000_100", num: 1_000_000_100, word: "Một Tỷ Một Trăm"},
		{name: "1_000_000_101", num: 1_000_000_101, word: "Một Tỷ Một Trăm Linh Một"},
		{name: "1_000_000_104", num: 1_000_000_104, word: "Một Tỷ Một Trăm Linh Bốn"},
		{name: "1_000_000_105", num: 1_000_000_105, word: "Một Tỷ Một Trăm Linh Năm"},
		{name: "1_000_000_110", num: 1_000_000_110, word: "Một Tỷ Một Trăm Mười"},
		{name: "1_000_000_111", num: 1_000_000_111, word: "Một Tỷ Một Trăm Mười Một"},
		{name: "1_000_000_114", num: 1_000_000_114, word: "Một Tỷ Một Trăm Mười Bốn"},
		{name: "1_000_000_115", num: 1_000_000_115, word: "Một Tỷ Một Trăm Mười Lăm"},
		{name: "1_000_000_120", num: 1_000_000_120, word: "Một Tỷ Một Trăm Hai Mươi"},
		{name: "1_000_000_121", num: 1_000_000_121, word: "Một Tỷ Một Trăm Hai Mươi Mốt"},
		{name: "1_000_000_124", num: 1_000_000_124, word: "Một Tỷ Một Trăm Hai Mươi Tư"},
		{name: "1_000_000_125", num: 1_000_000_125, word: "Một Tỷ Một Trăm Hai Mươi Lăm"},

		{name: "1_000_001_000", num: 1_000_001_000, word: "Một Tỷ Không Trăm Linh Một Nghìn"},
		{name: "1_000_001_005", num: 1_000_001_005, word: "Một Tỷ Không Trăm Linh Một Nghìn Không Trăm Linh Năm"},
		{name: "1_000_001_010", num: 1_000_001_010, word: "Một Tỷ Không Trăm Linh Một Nghìn Không Trăm Mười"},
		{name: "1_000_001_100", num: 1_000_001_100, word: "Một Tỷ Không Trăm Linh Một Nghìn Một Trăm"},
		{name: "1_000_001_105", num: 1_000_001_105, word: "Một Tỷ Không Trăm Linh Một Nghìn Một Trăm Linh Năm"},

		{name: "1_000_010_000", num: 1_000_010_000, word: "Một Tỷ Không Trăm Mười Nghìn"},
		{name: "1_000_010_005", num: 1_000_010_005, word: "Một Tỷ Không Trăm Mười Nghìn Không Trăm Linh Năm"},
		{name: "1_000_010_100", num: 1_000_010_100, word: "Một Tỷ Không Trăm Mười Nghìn Một Trăm"},

		{name: "1_000_100_000", num: 1_000_100_000, word: "Một Tỷ Một Trăm Nghìn"},
		{name: "1_000_100_005", num: 1_000_100_005, word: "Một Tỷ Một Trăm Nghìn Không Trăm Linh Năm"},
		{name: "1_000_100_100", num: 1_000_100_100, word: "Một Tỷ Một Trăm Nghìn Một Trăm"},

		{name: "1_001_000_000", num: 1_001_000_000, word: "Một Tỷ Không Trăm Linh Một Triệu"},
		{name: "1_001_000_005", num: 1_001_000_005, word: "Một Tỷ Không Trăm Linh Một Triệu Không Trăm Linh Năm"},
		{name: "1_001_000_100", num: 1_001_000_100, word: "Một Tỷ Không Trăm Linh Một Triệu Một Trăm"},

		{name: "1_001_001_000", num: 1_001_001_000, word: "Một Tỷ Không Trăm Linh Một Triệu Không Trăm Linh Một Nghìn"},
		{name: "1_001_001_005", num: 1_001_001_005, word: "Một Tỷ Không Trăm Linh Một Triệu Không Trăm Linh Một Nghìn Không Trăm Linh Năm"},

		{name: "1_010_000_000", num: 1_010_000_000, word: "Một Tỷ Không Trăm Mười Triệu"},
		{name: "1_010_000_005", num: 1_010_000_005, word: "Một Tỷ Không Trăm Mười Triệu Không Trăm Linh Năm"},
		{name: "1_010_001_005", num: 1_010_001_005, word: "Một Tỷ Không Trăm Mười Triệu Không Trăm Linh Một Nghìn Không Trăm Linh Năm"},

		{name: "1_100_000_000", num: 1_100_000_000, word: "Một Tỷ Một Trăm Triệu"},
		{name: "1_100_000_005", num: 1_100_000_005, word: "Một Tỷ Một Trăm Triệu Không Trăm Linh Năm"},
		{name: "1_100_001_005", num: 1_100_001_005, word: "Một Tỷ Một Trăm Triệu Không Trăm Linh Một Nghìn Không Trăm Linh Năm"},

		{name: "2_000_000_005", num: 2_000_000_005, word: "Hai Tỷ Không Trăm Linh Năm"},
		{name: "2_001_004_005", num: 2_001_004_005, word: "Hai Tỷ Không Trăm Linh Một Triệu Không Trăm Linh Bốn Nghìn Không Trăm Linh Năm"},
		{name: "20_000_000_005", num: 20_000_000_005, word: "Hai Mươi Tỷ Không Trăm Linh Năm"},
		{name: "200_000_000_005", num: 200_000_000_005, word: "Hai Trăm Tỷ Không Trăm Linh Năm"},

		{name: "1_234_561_891", num: 1_234_561_891, word: "Một Tỷ Hai Trăm Ba Mươi Bốn Triệu Năm Trăm Sáu Mươi Mốt Nghìn Tám Trăm Chín Mươi Mốt"},
		{name: "12_345_678_915", num: 12_345_678_915, word: "Mười Hai Tỷ Ba Trăm Bốn Mươi Lăm Triệu Sáu Trăm Bảy Mươi Tám Nghìn Chín Trăm Mười Lăm"},
		{name: "201_004_005_005", num: 201_004_005_005, word: "Hai Trăm Linh Một Tỷ Không Trăm Linh Bốn Triệu Không Trăm Linh Năm Nghìn Không Trăm Linh Năm"},
		{name: "999_999_999_999", num: 999_999_999_999, word: "Chín Trăm Chín Mươi Chín Tỷ Chín Trăm Chín Mươi Chín Triệu Chín Trăm Chín Mươi Chín Nghìn Chín Trăm Chín Mươi Chín"},

		// Trillion
		{name: "1_000_000_000_000", num: 1_000_000_000_000, word: "Một Nghìn Tỷ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberToWords(tt.num)
			require.Equal(t, tt.word, got)
		})
	}
}
