package main

import (
	"reflect"
	"testing"
)

func TestTlv(t *testing.T) {
	for _, tlvTestCase := range tlvTestCases {
		steps, err := Procesamiento(tlvTestCase.input)
		if tlvTestCase.expectError {
			if err == nil {
				t.Fatalf("FAIL: %s\n\tCollatzConjecture(%v) expected an error, got %v", tlvTestCase.description, tlvTestCase.input, steps)
			} else {
				t.Logf("PASS: %s", tlvTestCase.description)
			}
		} else {
			if err != nil {
				t.Fatalf("FAIL: %s\n\tCollatzConjecture(%v) returns unexpected error %s", tlvTestCase.description, tlvTestCase.input, err.Error())
			}
			if !reflect.DeepEqual(steps, tlvTestCase.expected) {
				t.Fatalf("FAIL: %s\n\tCollatzConjecture(%v) expected %v, got %v", tlvTestCase.description, tlvTestCase.input, tlvTestCase.expected, steps)
			}
			t.Logf("PASS: %s", tlvTestCase.description)
		}

	}
}

func BenchmarkTlv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tlvTestCase := range tlvTestCases {
			Procesamiento(tlvTestCase.input)
		}
	}
}
