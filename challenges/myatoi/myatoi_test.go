package challenges_test

import (
	"strconv"
	"strings"
	"testing"
	"unicode"
)

/*Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first
non-whitespace character is found. Then, starting from this character, takes an optional
initial plus or minus sign followed by as many numerical digits as possible, and interprets
them as a numerical value.

The string can contain additional characters after those that form the integral number,
 which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number,
or if no such sequence exists because either str is empty or it contains only whitespace characters,
no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store
integers within the 32-bit signed integer range: [−231,  231 − 1].
If the numerical value is out of the range of representable values,
INT_MAX (231 − 1) or INT_MIN (−231) is returned.

Example 1:

Input: "42"
Output: 42
Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
			 Thefore INT_MIN (−231) is returned.
*/
func myAtoi(str string) int {

	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return 0
	}

	multiply := 1
	if str[0] == '+' || str[0] == '-' {

		if str[0] == '-' {
			multiply = -1
		}

		str = str[1:]
	}
	toCast := ""
	for p, r := range str {

		if p == 0 && unicode.IsNumber(r) == false {
			return 0
		}

		if unicode.IsNumber(r) {
			//fmt.Printf("R: %s is a number\n", string(r))
			toCast += string(r)
			continue
		}
		break
	}

	if toCast != "" {
		n, err := strconv.ParseInt(toCast, 10, 32)
		if err != nil {
			n = 2147483647
		}
		j := int(n)
		j *= multiply
		return j

	}
	return 0
}

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"OK", args{"   -42"}, -42},
		{"OK", args{"4193 with word"}, 4193},
		{"OK", args{"words and 987"}, 0},
		{"OK", args{"-91283472332"}, -2147483648},
		{"OK", args{"3.1458"}, 3},
		{"OK", args{""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
