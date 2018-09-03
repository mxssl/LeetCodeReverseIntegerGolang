package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := 123456

	expected := 654321

	output := reverse(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}

func TestCase2(t *testing.T) {
	input := -123456

	expected := -654321

	output := reverse(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}

func TestCase3(t *testing.T) {
	input := 2147483648

	expected := 0

	output := reverse(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}

func TestCase4(t *testing.T) {
	input := -2147483649

	expected := 0

	output := reverse(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}
