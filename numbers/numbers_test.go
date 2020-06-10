package numbers

import (
	"testing"
)

// Add is bla
func TestAdd(t *testing.T) {
	data := []struct {
		title    string
		init     MyInt
		input    int32
		expected MyInt
	}{
		{"add", 1, 1, 2},
	}
	for _, d := range data {
		output := d.init.Add(d.input)
		if output != d.expected {
			t.Error("for", d.title, "got", output, "should got", d.expected)
		}
	}
}

func TestSub(t *testing.T) {
	data := []struct {
		title    string
		init     MyInt
		input    int32
		expected MyInt
	}{
		{"sub", 10, 5, 5},
	}
	for _, d := range data {
		output := d.init.Sub(d.input)
		if output != d.expected {
			t.Error("for", d.title, "got", output, "should got", d.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	data := []struct {
		title    string
		init     MyInt
		input    int32
		expected MyInt
	}{
		{"multiply", 10, 5, 50},
	}
	for _, d := range data {
		output := d.init.Multiply(d.input)
		if output != d.expected {
			t.Error("for", d.title, "got", output, "should got", d.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	data := []struct {
		title    string
		init     MyInt
		input    int32
		expected MyInt
	}{
		{"a", 100, 2, 50},
		{"b", 0, 10, 0},
		{"c", 10, 0, 0},
	}
	for _, d := range data {
		output := d.init.Divide(d.input)
		if output != d.expected {
			t.Error("for", d.title, "got", output, "should got", d.expected)
		}
	}
}
