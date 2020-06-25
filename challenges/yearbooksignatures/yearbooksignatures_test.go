package challenges_test

import (
	"fmt"
	"reflect"
	"testing"
)

/*
There are n students,
	numbered from 1 to n, each with their own yearbook.
They would like to pass their yearbooks around and get them signed by other students.
You're given a list of n integers arr[1..n],
	which is guaranteed to be a permutation of 1..n
	(in other words, it includes the integers from 1 to n exactly once each, in some order).

The meaning of this list is described below.
Initially, each student is holding their own yearbook.
The students will then repeat the following two steps each minute:
	Each student i will first sign the yearbook that they're currently holding
	(which may either belong to themselves or to another student),
	and then they'll pass it to student arr[i].
	It's possible that arr[i] = i for any given i, in which case student i will pass their yearbook back to themselves.

Once a student has received their own yearbook back, they will hold on to it and no longer participate in the passing process.

It's guaranteed that, for any possible valid input, each student will eventually receive their own yearbook back and will never end up holding more than one yearbook at a time.
You must compute a list of n integers output,
whose ith element is equal to the number of signatures that will be present in student i's
yearbook once they receive it back.*

Signature
int[] findSignatureCounts(int[] arr)

Input
n is in the range [1, 100,000].
Each value arr[i] is in the range [1, n], and all values in arr[i] are distinct.

Output
Return a list of n integers output, as described above.

Example 1
n = 2
arr = [2, 1]
output = [2, 2]
The first student will sign their own yearbook and pass it to the second,
who will also sign it and pass it back to the first student, resulting in 2 signatures.
Meanwhile, the second student's yearbook will similarly be signed both by themselves and then by the first student.

Example 2
n = 2
arr = [1, 2]
output = [1, 1]
Each student will simply pass their yearbook back to themselves, resulting in 1 signature each.
*/

func findSignatures(students []int) []int {
	// how many passes per student position
	var ans []int = make([]int, len(students))

	for pos, _ := range students {
		csp := pos
		ans[pos] = 1 // you can pass it to yourself so default 1
		thisStudent := pos + 1
		fmt.Printf("POS: %d STUDENT: %d PATH: %d\n", pos, thisStudent, students[csp])
		// keep passing until this student gets the book back
		for students[csp] != thisStudent { // if you get your yearbook back
			ans[pos]++              // increment the passes when looped
			csp = students[csp] - 1 // array is zero based, they are passing it to student 1 based
			fmt.Printf("\t J: %d  passed to student: %d\n", csp, students[csp])
		}
	}

	fmt.Printf("ANSWER: %v\n", ans)
	return ans
}

func Test_findSignatures(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{"testok", args{[]int{1, 2}}, []int{1, 1}},
		{"testok", args{[]int{2, 1}}, []int{2, 2}},
		{"testok", args{[]int{4, 3, 2, 5, 1}}, []int{3, 2, 2, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSignatures(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSignatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
