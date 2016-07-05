package lookandsay_test

import (
	"fmt"
	"github.com/aparkins/learnGo/lookandsay"
	"testing"
)

func TestWriteRoman(t *testing.T) {
	params := []struct {
		x        int
		expected string
	}{
		{1, "I"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{49, "XLIX"},
		{50, "L"},
		{73, "LXXIII"},
		{96, "XCVI"},
		{142, "CXLII"},
		{499, "CDXCIX"},
		{1998, "MCMXCVIII"},
	}

	for _, testCase := range params {
		actual := lookandsay.WriteRoman(testCase.x)
		if actual != testCase.expected {
			fmt.Printf("BAD ROMAN NUMERAL -- expected: %v | actual: %v\n", testCase.expected, actual)
			t.Fail()
		}
	}
}

func TestLookAndSay(t *testing.T) {
	params := []struct {
		writer   func(int) string
		n        int
		expected string
	}{
		{lookandsay.WriteStandard, 1, "1"},
		{lookandsay.WriteStandard, 2, "11"},
		{lookandsay.WriteStandard, 3, "21"},
		{lookandsay.WriteStandard, 4, "1211"},
		{lookandsay.WriteStandard, 5, "111221"},
		{lookandsay.WriteStandard, 6, "312211"},
		{lookandsay.WriteStandard, 7, "13112221"},
		{lookandsay.WriteStandard, 8, "1113213211"},

		{lookandsay.WriteRoman, 1, "I"},
		{lookandsay.WriteRoman, 2, "II"},
		{lookandsay.WriteRoman, 3, "III"},
		{lookandsay.WriteRoman, 4, "IIII"},
		{lookandsay.WriteRoman, 5, "IVI"},
		{lookandsay.WriteRoman, 6, "IIIVII"},
		{lookandsay.WriteRoman, 7, "IIIIIVIII"},
		{lookandsay.WriteRoman, 8, "VIIVIIII"},
		{lookandsay.WriteRoman, 9, "IVIIIIVIVI"},
		{lookandsay.WriteRoman, 10, "IIIVIVIIVIIIVII"},
		{lookandsay.WriteRoman, 11, "IIIIIVIIIVIIIIVIIIIIVIII"},
		{lookandsay.WriteRoman, 12, "VIIVIIIIIVIVIIVVIIVIIII"},
	}

	for _, testCase := range params {
		actual := lookandsay.LookAndSay(testCase.writer, testCase.n)
		if actual != testCase.expected {
			fmt.Printf("BAD LOOK AND SAY -- expected: %v | actual : %v\n", testCase.expected, actual)
			t.Fail()
		}
	}
}

func TestLookAndSayWriter(t *testing.T) {
	params := []struct {
		writer   func(int) string
		expected []string
	}{
		{
			writer:   lookandsay.WriteStandard,
			expected: []string{"1", "11", "21", "1211", "111221", "312211", "13112221", "1113213211"},
		},

		{
			writer: lookandsay.WriteRoman,
			expected: []string{
				"I", "II", "III", "IIII", "IVI", "IIIVII", "IIIIIVIII", "VIIVIIII", "IVIIIIVIVI",
				"IIIVIVIIVIIIVII", "IIIIIVIIIVIIIIVIIIIIVIII", "VIIVIIIIIVIVIIVVIIVIIII",
			},
		},
	}

	for _, testCase := range params {
		channels := lookandsay.LookAndSayWriter(testCase.writer)
		for _, expected := range testCase.expected {
			actual := <-channels.WriteChan
			if actual != expected {
				fmt.Printf("BAD LOOK AND SAY WRITER -- expected: %v | actual : %v\n", testCase.expected, actual)
				t.Fail()
			}
		}
	}
}
