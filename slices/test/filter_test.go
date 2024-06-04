package test_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/luviz/go-pkgs/slices"
)

func TestFilterNumbers(t *testing.T) {
	ints := []int{1, 3, 2, 4, 5, 2, 1, 6, 2, 8, 2, 9, 0, 8, 7}
	filterRes := slices.Filter(ints, func(v, _ int, _ []int) bool {
		return v < 5
	})
	fmt.Printf("res: %#v", filterRes)
	for _, v := range filterRes {
		if v > 5 {
			t.Fatal(v, "larger then 5")
		}
	}
}

func TestFilterString(t *testing.T) {
	names := []string{
		"Albert",
		"Betty",
		"Casey",
		"Danny",
		"Eddie",
		"Freddy",
	}
	filterRes := slices.Filter(names, func(v string, _ int, _ []string) bool {
		return strings.Contains(v, "a")
	})
	fmt.Printf("res: %#v", filterRes)
	for _, v := range filterRes {
		if !strings.Contains(v, "a") {
			t.Fatal(v, "Dose not contains a")
		}
	}
}
