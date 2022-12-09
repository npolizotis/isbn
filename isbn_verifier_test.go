package main

import (
	"testing"
)

func TestIsValidISBN(t *testing.T) {
	for _, test := range testCases {
		observed := IsValidISBN(test.isbn)
		if observed == test.expected {
			t.Logf("PASS: %s", test.description)
		} else {
			t.Errorf("FAIL: %s\nIsValidISBN(%q)\nExpected: %t, Actual: %t",
				test.description, test.isbn, test.expected, observed)
		}
	}
}

func BenchmarkIsValidISBN(t *testing.B) {
	for _, test := range testCases {
		t.Run(test.isbn, func(t2 *testing.B) {
			for i := 0; i < t2.N; i++ {
				IsValidISBN(test.isbn)
			}
		})
	}
}

func BenchmarkIsValidISBNSingle(t *testing.B) {
	for i := 0; i < t.N; i++ {
		IsValidISBN("3-598-21507-X")
	}
}
