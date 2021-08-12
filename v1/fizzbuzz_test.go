package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzzLength(t *testing.T) {

	t.Run("Creates array of any size passed", func(t *testing.T) {

		test9 := FizzBuzz(9)

		got := len(test9)
		want := 9

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, test9)
		}
	})

}

func TestFizz(t *testing.T) {

	got := FizzBuzz(4)
	want := []string{"1", "2", "Fizz", "4"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestBuzz(t *testing.T) {

	got := FizzBuzz(6)[4]
	want := "Buzz"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFizzBuzz(t *testing.T) {

	got := FizzBuzz(16)[14]
	want := "FizzBuzz"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
