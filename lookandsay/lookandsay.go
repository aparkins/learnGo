package lookandsay

import (
    "strconv"
)

func LookAndSay(writer func(int) string, n int) string {
    current := writer(1)
    
    for i := 1; i < n; i++ {
        current = nextLookAndSay(writer, current)
    }

    return current
}

func nextLookAndSay(writer func(int) string, current string) string {
    result := ""

    for chunk, remainder := splitCharRange(current);
        chunk != "";
        chunk, remainder = splitCharRange(remainder) {
	result += writer(len(chunk)) + string(chunk[0])
    }

    return result
}

func splitCharRange(str string) (string, string) {
    if str == "" {
        return "", ""
    }
    curChar := str[0]
    i := 1
    for i < len(str) {
        if str[i] != curChar {
	    break
        }
        i++
    }
    return str[0:i], str[i:]
}

func WriteStandard(x int) string {
    return strconv.Itoa(x)
}

func WriteRoman(x int) string {
    type recfunc func(recfunc, int, string) string
    innerRoman := func (f recfunc, x int, acc string) string {
        if x == 0 {
            if acc == "" {
                return "0"
            } else {
                return acc
            }
        } else if x >= 1000 {
            x -= 1000
            acc += "M"
        } else if x >= 900 {
            x -= 900
            acc += "CM"
        } else if x >= 500 {
            x -= 500
            acc += "D"
        } else if x >= 400 {
            x -= 400
            acc += "CD"
        } else if x >= 100 {
            x -= 100
            acc += "C"
        } else if x >= 90 {
            x -= 90
            acc += "XC"
        } else if x >= 50 {
            x -= 50
            acc += "L"
        } else if x >= 40 {
            x -= 40
            acc += "XL"
        } else if x >= 10 {
            x -= 10
            acc += "X"
        } else if x >= 9 {
            x -= 9
            acc += "IX"
        } else if x >= 5 {
            x -= 5
            acc += "V"
        } else if x >= 4 {
            x -= 4
            acc += "IV"
        } else {
            x -= 1
            acc += "I"
        }
	return f(f, x, acc)
    }
    return innerRoman(innerRoman, x, "")
}
