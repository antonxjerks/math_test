package golearn

import (
	"testing"
)

func TestAbsolute(t *testing.T) { // class yg utk di pakai testing
	t.Run("negative test case", func(t *testing.T) {

	})
	t.Run("positive test case", func(t *testing.T) {

	})
	c := Absolute(-5)

	if c != 5 {
		t.Logf("expect 5, but got %d", c)
		t.Fail() // untuk kita mengagalkan function unit test
	}
}

func TestAbsolute_WithPositive(t *testing.T) {
	c := Absolute(5)
	if c != 5 {
		t.Logf("expect 5, but got %d", c)
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	t.Run("negative case test", func(t *testing.T) {

	})
}
