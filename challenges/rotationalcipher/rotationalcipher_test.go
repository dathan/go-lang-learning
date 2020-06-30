package challenges_test

import (
	"testing"
	"unicode"
)

/*
Rotational Cipher
One simple way to encrypt a string is to "rotate" every alphanumeric character by a certain amount. Rotating a character means replacing it with another character that is a certain number of steps away in normal alphabetic or numerical order.
For example, if the string "Zebra-493?" is rotated 3 places, the resulting string is "Cheud-726?". Every alphabetic character is replaced with the character 3 letters higher (wrapping around from Z to A), and every numeric character replaced with the character 3 digits higher (wrapping around from 9 to 0). Note that the non-alphanumeric characters remain unchanged.
Given a string and a rotation factor, return an encrypted string.
Signature
string rotationalCipher(string input, int rotationFactor)
Input
1 <= |input| <= 1,000,000
0 <= rotationFactor <= 1,000,000
Output
Return the result of rotating input a number of times equal to rotationFactor.
Example 1
input = Zebra-493?
rotationFactor = 3
output = Cheud-726?
Example 2
input = abcdefghijklmNOPQRSTUVWXYZ0123456789
rotationFactor = 39
output = nopqrstuvwxyzABCDEFGHIJKLM9012345678
*/

func rotationalCipher(input string, rotationFactor int) string {
	// going through a rune list
	ret := ""
	for _, r := range input {
		rF := rotationFactor
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if (r >= 'a' && r <= 'z') || r >= 'A' && r <= 'Z' {
				if rotationFactor > 26 {
					rF = (rotationFactor) % 26
				}

			}

			if r >= '0' && r <= '9' {

				if rotationFactor > 9 {
					rF = rotationFactor % 10
				}

				// if the sum wraps handle the wrap
				if (int(r) + rF) > '9' {
					form := (int(r) - 10) + rF
					ret += string(form)
					//fmt.Printf("NUM Replacing: %c (%d) with %s (%d) rotationalFactor: %d\n", r, int(r), string(form), form, rF)
					continue
				}

			}

			// if the sum wraps handle the wrap
			if r >= 'n' && r <= 'z' {
				if (int(r) + rF) > 'z' {
					form := (int(r) - 26) + rF
					ret += string(rune(form))
					//fmt.Printf("LOWER Replacing: %c (%d) with %s (%d) rotationalFactor: %d == %s\n", r, int(r), string(rune(form)), form, rF, ret)
					continue
				}
			}

			if r >= 'N' && r <= 'Z' {
				if (int(r) + rF) > 'Z' {
					form := (int(r) - 26) + rF
					ret += string(rune(form))
					//fmt.Printf("UPPER Replacing: %c (%d) with %s (%d) rotationalFactor: %d == %s\n", r, int(r), string(rune(form)), form, rF, ret)
					continue
				}
			}

			//fmt.Printf("Replacing: %c (%d) with %s (%d) rotationalFactor: %d\n", r, int(r), string(int(r)+rF), (int(r) + rF), rF)
			ret += string(int(r) + rF)
			continue
		}
		ret += string(r)

	}
	//fmt.Printf("DEPLOY: %s\n", ret)
	return ret
}

func Test_rotationalCipher(t *testing.T) {
	type args struct {
		input          string
		rotationFactor int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestOK", args{"Zebra-493?", 3}, "Cheud-726?"},
		{"TestOK", args{"nopqrstuvwxyz0123456789", 39}, "abcdefghijklm9012345678"},
		{"TestOK", args{"abcdefghijklmNOPQRSTUVWXYZ0123456789", 39}, "nopqrstuvwxyzABCDEFGHIJKLM9012345678"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotationalCipher(tt.args.input, tt.args.rotationFactor); got != tt.want {
				t.Errorf("rotationalCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
