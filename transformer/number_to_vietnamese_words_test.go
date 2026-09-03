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
		{name: "0", num: 0, word: "không"},
		{name: "1", num: 1, word: "một"},
		{name: "4", num: 4, word: "bốn"},
		{name: "5", num: 5, word: "năm"},
		{name: "9", num: 9, word: "chín"},

		{name: "10", num: 10, word: "mười"},
		{name: "11", num: 11, word: "mười một"},
		{name: "14", num: 14, word: "mười bốn"},
		{name: "15", num: 15, word: "mười lăm"},
		{name: "19", num: 19, word: "mười chín"},

		{name: "20", num: 20, word: "hai mươi"},
		{name: "21", num: 21, word: "hai mươi mốt"},
		{name: "24", num: 24, word: "hai mươi tư"},
		{name: "25", num: 25, word: "hai mươi lăm"},
		{name: "29", num: 29, word: "hai mươi chín"},

		// 100 - 999
		{name: "100", num: 100, word: "một trăm"},
		{name: "101", num: 101, word: "một trăm linh một"},
		{name: "104", num: 104, word: "một trăm linh bốn"},
		{name: "105", num: 105, word: "một trăm linh năm"},
		{name: "110", num: 110, word: "một trăm mười"},
		{name: "111", num: 111, word: "một trăm mười một"},
		{name: "114", num: 114, word: "một trăm mười bốn"},
		{name: "115", num: 115, word: "một trăm mười lăm"},
		{name: "120", num: 120, word: "một trăm hai mươi"},
		{name: "121", num: 121, word: "một trăm hai mươi mốt"},
		{name: "124", num: 124, word: "một trăm hai mươi tư"},
		{name: "125", num: 125, word: "một trăm hai mươi lăm"},
		{name: "140", num: 140, word: "một trăm bốn mươi"},
		{name: "141", num: 141, word: "một trăm bốn mươi mốt"},
		{name: "144", num: 144, word: "một trăm bốn mươi tư"},
		{name: "145", num: 145, word: "một trăm bốn mươi lăm"},
		{name: "150", num: 150, word: "một trăm năm mươi"},
		{name: "151", num: 151, word: "một trăm năm mươi mốt"},
		{name: "154", num: 154, word: "một trăm năm mươi tư"},
		{name: "155", num: 155, word: "một trăm năm mươi lăm"},
		{name: "500", num: 500, word: "năm trăm"},
		{name: "501", num: 501, word: "năm trăm linh một"},
		{name: "504", num: 504, word: "năm trăm linh bốn"},
		{name: "505", num: 505, word: "năm trăm linh năm"},
		{name: "999", num: 999, word: "chín trăm chín mươi chín"},

		// Thousands
		{name: "1_000", num: 1_000, word: "một nghìn"},
		{name: "1_001", num: 1_001, word: "một nghìn không trăm linh một"},
		{name: "1_010", num: 1_010, word: "một nghìn không trăm mười"},
		{name: "1_011", num: 1_011, word: "một nghìn không trăm mười một"},
		{name: "1_100", num: 1_100, word: "một nghìn một trăm"},
		{name: "1_111", num: 1_111, word: "một nghìn một trăm mười một"},

		{name: "2_005", num: 2_005, word: "hai nghìn không trăm linh năm"},
		{name: "2_015", num: 2_015, word: "hai nghìn không trăm mười lăm"},
		{name: "2_021", num: 2_021, word: "hai nghìn không trăm hai mươi mốt"},

		{name: "10_000", num: 10_000, word: "mười nghìn"},
		{name: "10_001", num: 10_001, word: "mười nghìn không trăm linh một"},
		{name: "10_010", num: 10_010, word: "mười nghìn không trăm mười"},
		{name: "10_011", num: 10_011, word: "mười nghìn không trăm mười một"},
		{name: "10_100", num: 10_100, word: "mười nghìn một trăm"},
		{name: "11_000", num: 11_000, word: "mười một nghìn"},

		{name: "20_005", num: 20_005, word: "hai mươi nghìn không trăm linh năm"},
		{name: "20_015", num: 20_015, word: "hai mươi nghìn không trăm mười lăm"},
		{name: "20_105", num: 20_105, word: "hai mươi nghìn một trăm linh năm"},
		{name: "201_005", num: 201_005, word: "hai trăm linh một nghìn không trăm linh năm"},

		{name: "100_000", num: 100_000, word: "một trăm nghìn"},
		{name: "100_005", num: 100_005, word: "một trăm nghìn không trăm linh năm"},
		{name: "100_011", num: 100_011, word: "một trăm nghìn không trăm mười một"},
		{name: "100_100", num: 100_100, word: "một trăm nghìn một trăm"},
		{name: "101_000", num: 101_000, word: "một trăm linh một nghìn"},
		{name: "101_005", num: 101_005, word: "một trăm linh một nghìn không trăm linh năm"},
		{name: "110_000", num: 110_000, word: "một trăm mười nghìn"},
		{name: "999_999", num: 999_999, word: "chín trăm chín mươi chín nghìn chín trăm chín mươi chín"},

		// Millions
		{name: "1_000_000", num: 1_000_000, word: "một triệu"},
		{name: "1_000_001", num: 1_000_001, word: "một triệu không trăm linh một"},
		{name: "1_000_010", num: 1_000_010, word: "một triệu không trăm mười"},
		{name: "1_000_011", num: 1_000_011, word: "một triệu không trăm mười một"},

		{name: "1_000_100", num: 1_000_100, word: "một triệu một trăm"},
		{name: "1_001_000", num: 1_001_000, word: "một triệu không trăm linh một nghìn"},
		{name: "1_010_000", num: 1_010_000, word: "một triệu không trăm mười nghìn"},
		{name: "1_100_000", num: 1_100_000, word: "một triệu một trăm nghìn"},

		{name: "2_000_005", num: 2_000_005, word: "hai triệu không trăm linh năm"},
		{name: "2_001_005", num: 2_001_005, word: "hai triệu không trăm linh một nghìn không trăm linh năm"},
		{name: "20_000_005", num: 20_000_005, word: "hai mươi triệu không trăm linh năm"},
		{name: "20_010_005", num: 20_010_005, word: "hai mươi triệu không trăm mười nghìn không trăm linh năm"},
		{name: "200_000_005", num: 200_000_005, word: "hai trăm triệu không trăm linh năm"},

		{name: "1_234_567", num: 1_234_567, word: "một triệu hai trăm ba mươi bốn nghìn năm trăm sáu mươi bảy"},
		{name: "999_999_999", num: 999_999_999, word: "chín trăm chín mươi chín triệu chín trăm chín mươi chín nghìn chín trăm chín mươi chín"},

		// Billions
		{name: "1_000_000_000", num: 1_000_000_000, word: "một tỷ"},
		{name: "1_000_000_001", num: 1_000_000_001, word: "một tỷ không trăm linh một"},
		{name: "1_000_000_004", num: 1_000_000_004, word: "một tỷ không trăm linh bốn"},
		{name: "1_000_000_005", num: 1_000_000_005, word: "một tỷ không trăm linh năm"},
		{name: "1_000_000_010", num: 1_000_000_010, word: "một tỷ không trăm mười"},
		{name: "1_000_000_011", num: 1_000_000_011, word: "một tỷ không trăm mười một"},
		{name: "1_000_000_014", num: 1_000_000_014, word: "một tỷ không trăm mười bốn"},
		{name: "1_000_000_015", num: 1_000_000_015, word: "một tỷ không trăm mười lăm"},
		{name: "1_000_000_020", num: 1_000_000_020, word: "một tỷ không trăm hai mươi"},
		{name: "1_000_000_021", num: 1_000_000_021, word: "một tỷ không trăm hai mươi mốt"},
		{name: "1_000_000_024", num: 1_000_000_024, word: "một tỷ không trăm hai mươi tư"},
		{name: "1_000_000_025", num: 1_000_000_025, word: "một tỷ không trăm hai mươi lăm"},

		{name: "1_000_000_100", num: 1_000_000_100, word: "một tỷ một trăm"},
		{name: "1_000_000_101", num: 1_000_000_101, word: "một tỷ một trăm linh một"},
		{name: "1_000_000_104", num: 1_000_000_104, word: "một tỷ một trăm linh bốn"},
		{name: "1_000_000_105", num: 1_000_000_105, word: "một tỷ một trăm linh năm"},
		{name: "1_000_000_110", num: 1_000_000_110, word: "một tỷ một trăm mười"},
		{name: "1_000_000_111", num: 1_000_000_111, word: "một tỷ một trăm mười một"},
		{name: "1_000_000_114", num: 1_000_000_114, word: "một tỷ một trăm mười bốn"},
		{name: "1_000_000_115", num: 1_000_000_115, word: "một tỷ một trăm mười lăm"},
		{name: "1_000_000_120", num: 1_000_000_120, word: "một tỷ một trăm hai mươi"},
		{name: "1_000_000_121", num: 1_000_000_121, word: "một tỷ một trăm hai mươi mốt"},
		{name: "1_000_000_124", num: 1_000_000_124, word: "một tỷ một trăm hai mươi tư"},
		{name: "1_000_000_125", num: 1_000_000_125, word: "một tỷ một trăm hai mươi lăm"},

		{name: "1_000_001_000", num: 1_000_001_000, word: "một tỷ không trăm linh một nghìn"},
		{name: "1_000_001_005", num: 1_000_001_005, word: "một tỷ không trăm linh một nghìn không trăm linh năm"},
		{name: "1_000_001_010", num: 1_000_001_010, word: "một tỷ không trăm linh một nghìn không trăm mười"},
		{name: "1_000_001_100", num: 1_000_001_100, word: "một tỷ không trăm linh một nghìn một trăm"},
		{name: "1_000_001_105", num: 1_000_001_105, word: "một tỷ không trăm linh một nghìn một trăm linh năm"},

		{name: "1_000_010_000", num: 1_000_010_000, word: "một tỷ không trăm mười nghìn"},
		{name: "1_000_010_005", num: 1_000_010_005, word: "một tỷ không trăm mười nghìn không trăm linh năm"},
		{name: "1_000_010_100", num: 1_000_010_100, word: "một tỷ không trăm mười nghìn một trăm"},

		{name: "1_000_100_000", num: 1_000_100_000, word: "một tỷ một trăm nghìn"},
		{name: "1_000_100_005", num: 1_000_100_005, word: "một tỷ một trăm nghìn không trăm linh năm"},
		{name: "1_000_100_100", num: 1_000_100_100, word: "một tỷ một trăm nghìn một trăm"},

		{name: "1_001_000_000", num: 1_001_000_000, word: "một tỷ không trăm linh một triệu"},
		{name: "1_001_000_005", num: 1_001_000_005, word: "một tỷ không trăm linh một triệu không trăm linh năm"},
		{name: "1_001_000_100", num: 1_001_000_100, word: "một tỷ không trăm linh một triệu một trăm"},

		{name: "1_001_001_000", num: 1_001_001_000, word: "một tỷ không trăm linh một triệu không trăm linh một nghìn"},
		{name: "1_001_001_005", num: 1_001_001_005, word: "một tỷ không trăm linh một triệu không trăm linh một nghìn không trăm linh năm"},

		{name: "1_010_000_000", num: 1_010_000_000, word: "một tỷ không trăm mười triệu"},
		{name: "1_010_000_005", num: 1_010_000_005, word: "một tỷ không trăm mười triệu không trăm linh năm"},
		{name: "1_010_001_005", num: 1_010_001_005, word: "một tỷ không trăm mười triệu không trăm linh một nghìn không trăm linh năm"},

		{name: "1_100_000_000", num: 1_100_000_000, word: "một tỷ một trăm triệu"},
		{name: "1_100_000_005", num: 1_100_000_005, word: "một tỷ một trăm triệu không trăm linh năm"},
		{name: "1_100_001_005", num: 1_100_001_005, word: "một tỷ một trăm triệu không trăm linh một nghìn không trăm linh năm"},

		{name: "2_000_000_005", num: 2_000_000_005, word: "hai tỷ không trăm linh năm"},
		{name: "2_001_004_005", num: 2_001_004_005, word: "hai tỷ không trăm linh một triệu không trăm linh bốn nghìn không trăm linh năm"},
		{name: "20_000_000_005", num: 20_000_000_005, word: "hai mươi tỷ không trăm linh năm"},
		{name: "200_000_000_005", num: 200_000_000_005, word: "hai trăm tỷ không trăm linh năm"},

		{name: "1_234_561_891", num: 1_234_561_891, word: "một tỷ hai trăm ba mươi bốn triệu năm trăm sáu mươi mốt nghìn tám trăm chín mươi mốt"},
		{name: "12_345_678_915", num: 12_345_678_915, word: "mười hai tỷ ba trăm bốn mươi lăm triệu sáu trăm bảy mươi tám nghìn chín trăm mười lăm"},
		{name: "201_004_005_005", num: 201_004_005_005, word: "hai trăm linh một tỷ không trăm linh bốn triệu không trăm linh năm nghìn không trăm linh năm"},
		{name: "999_999_999_999", num: 999_999_999_999, word: "chín trăm chín mươi chín tỷ chín trăm chín mươi chín triệu chín trăm chín mươi chín nghìn chín trăm chín mươi chín"},

		{name: "10_000_000_000", num: 10_000_000_000, word: "mười tỷ"},
		{name: "100_000_000_000", num: 100_000_000_000, word: "một trăm tỷ"},
		{name: "1_000_000_000_000", num: 1_000_000_000_000, word: "một nghìn tỷ"},
		{name: "10_000_000_000_000", num: 10_000_000_000_000, word: "mười nghìn tỷ"},
		{name: "100_000_000_000_000", num: 100_000_000_000_000, word: "một trăm nghìn tỷ"},
		{name: "1_000_000_000_000_000", num: 1_000_000_000_000_000, word: "một triệu tỷ"},

		{name: "9_223_372_036_854_775_807", num: 9_223_372_036_854_775_807, word: "chín tỷ tỷ hai trăm hai mươi ba triệu tỷ ba trăm bảy mươi hai nghìn tỷ không trăm ba mươi sáu tỷ tám trăm năm mươi bốn triệu bảy trăm bảy mươi lăm nghìn tám trăm linh bảy"},

		{name: "-1", num: -1, word: "âm một"},
		{name: "-9_223_372_036_854_775_808", num: -9_223_372_036_854_775_808, word: "âm chín tỷ tỷ hai trăm hai mươi ba triệu tỷ ba trăm bảy mươi hai nghìn tỷ không trăm ba mươi sáu tỷ tám trăm năm mươi bốn triệu bảy trăm bảy mươi lăm nghìn tám trăm linh tám"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberToVietnameseWords(tt.num)
			require.Equal(t, tt.word, got)
		})
	}
}
