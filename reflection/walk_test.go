package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		name          string
		input         interface{}
		expectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string field",
			struct {
				name string
				city string
			}{"Chris", "Paris"},
			[]string{"Chris", "Paris"},
		},
		{
			"struct with a string and int field",
			struct {
				name string
				age  int
			}{"Chris", 22},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			var got []string
			walk(test.input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.expectedCalls) {
				t.Errorf("\nExpected :%v \ngot :%v", test.expectedCalls, got)
			}
		})
	}

}
