package slices_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/luviz/go-pkgs/slices"
)

func TestFilterNumb(t *testing.T) {
	ints := []int{1, 3, 2, 4, 5, 2, 1, 6, 2, 8, 2, 9, 0, 8, 7}
	filterRes := slices.Filter(ints, func(v, _ int, _ []int) bool {
		return v < 5
	})
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
	for _, v := range filterRes {
		if !strings.Contains(v, "a") {
			t.Fatal(v, "Dose not contains a")
		}
	}
}

func TestZip(t *testing.T) {
	values := []int{1, 3, 2, 4, 5, 2, 1, 6, 2, 8, 2, 9, 0, 8, 7}
	names := []string{
		"Albert",
		"Betty",
		"Casey",
		"Danny",
		"Eddie",
		"Freddy",
	}

	nameVal := slices.Zip(names, values)
	min_len := min(len(values), len(names))
	for ix := 0; ix < min_len; ix++ {
		if nameVal[ix].First != names[ix] || nameVal[ix].Second != values[ix] {
			fmt.Println(nameVal[ix], names[ix], values[ix])
			t.Fail()
		}
	}
}

type TestMapOut struct {
	name, altName string
	ix            int
}

func TestMap(t *testing.T) {
	changeName := func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), "a", "_")
	}

	names := []string{
		"Albert",
		"Betty",
		"Casey",
		"Danny",
		"Eddie",
		"Freddy",
	}
	tmo := slices.Map(names, func(n string, ix int, _ []string) TestMapOut {
		return TestMapOut{name: n, ix: ix, altName: changeName(n)}
	})

	if len(tmo) != len(names) {
		fmt.Println(tmo)
		t.Fail()
	}

	for ix, v := range tmo {
		orgName := names[ix]
		altName := changeName(orgName)

		if v.name != orgName || v.altName == altName && v.ix == ix {
			continue
		}
		fmt.Println(ix, v)
		t.Fail()
	}
}
