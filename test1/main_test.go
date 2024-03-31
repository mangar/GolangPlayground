package main

import (
	"testing"
)

//
// go test -v
//

func TestRepeatCount(t *testing.T) {

	type testRepeatCount struct {
		Entry string
		Expected []string
		ExpectedCount []int
	}

	var tests = []testRepeatCount {
		testRepeatCount{"abcadefbghicjkldm", []string{"a", "b", "c", "d"}, []int{2, 2, 2, 2}},
		testRepeatCount{"aaaa", []string{"a"}, []int{4}},
		testRepeatCount{"aabaab", []string{"a", "b"}, []int{4, 2}},
	}


	for _, test := range tests {
		out := repeatCount(test.Entry)
		
		for i, k := range test.Expected {

			if out[k] != test.ExpectedCount[i] {
				t.Errorf("Output %q != to expected %q", out[k], test.ExpectedCount[i])	
			}

		}
		
	}

}




func TestReplaceDuplicatedWith(t *testing.T) {

	type testReplaceDuplicatedWith struct {
		Entry string
		ReplaceChat string
		Expected string
	}

	var tests = []testReplaceDuplicatedWith {
		testReplaceDuplicatedWith{"abcadefbghicjkldm", "_", "abc_def_ghi_jkl_m"},
		testReplaceDuplicatedWith{"aaaa", "_", "a___"},
		testReplaceDuplicatedWith{"aabaab", "_", "a_b___"},
	}


	for _, test := range tests {
		if out := replaceDuplicatedWith( test.ReplaceChat, test.Entry ); out != test.Expected {
			t.Errorf("Output %q != to expected %q", out, test.Expected)
		}
	}

}



// 
// go test -bench=.
// 

func BenchmarkRepeatCount(b *testing.B) {
	b.ReportAllocs()
	repeatCount("abcdefghijklmnopqrstuvwxyzzywqqasjfkwlwowvmsklsdcmcnasljajijo")
}

func BenchmarkRepeatCountIFExist(b *testing.B) {
	b.ReportAllocs()
	repeatCountIFExist("abcdefghijklmnopqrstuvwxyzzywqqasjfkwlwowvmsklsdcmcnasljajijo")
}

// func BenchmarkReplaceDuplicatedWith(b *testing.B) {
// 	b.ReportAllocs()
// 	replaceDuplicatedWith( "_", "abcdefghijklmnopqrstuvwxyzzywqqasjfkwlwowvmsklsdcmcnasljajijo" )
// }