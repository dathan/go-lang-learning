//A recursive solution is suitable for the types of problems that require you to
//support your next result with the outcomes of previous invocations of your function.

package challenges_test

import (
	"fmt"
	"testing"
)

/**
Encrypted Words
You've devised a simple encryption method for alphabetic strings that shuffles the characters in such a way
that the resulting string is hard to quickly read,
but is easy to convert back into the original string.

When you encrypt a string S,
you start with an initially-empty resulting string R and append characters to it as follows:

Append the middle character of S
(
	if S has even length,
		then we define the middle character as the left-most of the two central characters
)

Append the encrypted version of the substring of S that's to the left of the middle character (if non-empty)

Append the encrypted version of the substring of S that's to the right of the middle character (if non-empty)

For example, to encrypt the string "abc", we first take "b",
and then append the encrypted version of "a" (which is just "a")
and the encrypted version of "c" (which is just "c") to get "bac".

If we encrypt "abcxcba" we'll get "xbacbca".
That is, we take "x" and then append the encrypted version "abc" and
then append the encrypted version of "cba".

Input
	S contains only lower-case alphabetic characters

	1 <= |S| <= 10,000

Output
	Return string R, the encrypted version of S.

	Example 1
	S = "abc"
	R = "bac"

	Example 2
	S = "abcd"
	R = "bacd"

	Example 3
	S = "abcxcba"
	R = "xbacbca"

	Example 4
	S = "facebook"
	R = "eafcobok"
*/
func findEncryptedWord(s string) string {

	slen := len(s)
	if slen == 0 {
		return ""
	}

	if slen == 1 {
		return s
	}

	// initally start is an empty result r
	r := ""
	middleLen := (slen - 1) / 2

	// from index middleLen to length of string
	middleChar := string(s[middleLen])

	//0:a 1:b 2:c
	fmt.Printf("middleChar: %s middleLen: %d slen: %d LEFT: %s RIGHT: %s\n", middleChar, middleLen, slen, s[0:middleLen], s[middleLen+1:])

	r = middleChar + findEncryptedWord(s[0:middleLen]) + findEncryptedWord(s[middleLen+1:]) // now you want to modify the info your passing forwardf
	return r
}

/**
Change in a Foreign Currency
You likely know that different currencies have coins and bills of different denominations.
In some currencies, it's actually impossible to receive change for a given amount of money.
For example, Canada has given up the 1-cent penny.
If you're owed 94 cents in Canada, a shopkeeper will graciously supply you with
95 cents instead since there exists a 5-cent coin.

Given a list of the available denominations,
determine if it's possible to receive exact change for an amount of money targetMoney.
Both the denominations and target amount will be given in generic units of that currency.

Signature
boolean canGetExactChange(int targetMoney, int[] denominations)
Input
1 ≤ |denomintaions| ≤ 100
1 ≤ denominations[i] ≤ 10,000
1 ≤ targetMoney ≤ 1,000,000

Output
Return true if it's possible to receive exactly targetMoney given the available denominations, and false if not.

Example 1
denominations = [5, 10, 25, 100, 200]
targetMoney = 94
output = false
Every denomination is a multiple of 5, so you can't receive exactly 94 units of money in this currency.

Example 2
denominations = [4, 17, 29]
targetMoney = 75
output = true
You can make 75 units with the denominations [17, 29, 29]
*/
func canGetExactChange(targetMoney int, denominations []int) bool {

	// need the base case - this is how recurssion stops and really begins to operate
	if targetMoney == 0 {
		return true // by returning a value which doesn’t require any further code execution, you enable the recusion to stop
	}

	if len(denominations) == 0 {
		return false
	}

	if len(denominations) == 1 && targetMoney%denominations[0] == 0 {
		return true
	}

	if len(denominations) == 1 {
		return false
	}

	coin := denominations[0]
	for i := 0; i*coin < targetMoney; i++ {
		// ALWAYS MODIFY THE INFORMATION YOU’RE PASSING FORWARD
		if canGetExactChange(targetMoney-i*coin, denominations[1:]) { // takeing one step forward to the base case

			//NOTE: recusion is not iteration
			return true
		}
	}
	return false
}

func Test_canGetExactChange(t *testing.T) {
	type args struct {
		targetMoney   int
		denominations []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"exactchange", args{75, []int{4, 17, 29}}, true},
		{"notexactchange", args{94, []int{5, 10, 25, 100, 200}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canGetExactChange(tt.args.targetMoney, tt.args.denominations); got != tt.want {
				t.Errorf("canGetExactChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findEncryptedWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"abc", args{"abc"}, "bac"},
		{"abcd", args{"abcd"}, "bacd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEncryptedWord(tt.args.s); got != tt.want {
				t.Errorf("findEncryptedWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
