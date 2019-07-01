package dictionary

import "testing"

func TestSearch(t *testing.T) {

	d := Dictionary{"test": "this is just a test"}

	searchTests := []struct {
		name        string
		key         string
		want        string
		expectedErr error
	}{
		{name: "Simple key value retrieval", key: "test", want: "this is just a test"},
		{name: "No entry", key: "NOKEY", want: "", expectedErr: ErrWordNotFound},
	}

	for _, test := range searchTests {
		t.Run(test.name, func(t *testing.T) {
			got, err := d.Search(test.key)
			if test.expectedErr != nil {
				assertError(t, err, test.expectedErr)
			} else {
				assertNoError(t, err)
				assertStrings(t, got, test.want)
			}
		})
	}

}

func TestUpdate(t *testing.T) {
	t.Run("Simple definition update test", func(t *testing.T) {
		word := "test"
		definition := "small definiton"
		dict := Dictionary{word: definition}
		newDefinition := "newer and longer defintion"

		err := dict.Update(word, newDefinition)

		assertNoError(t, err)
		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("Updating unexisting word test", func(t *testing.T) {
		dict := Dictionary{}
		newDefinition := "newer and longer defintion"

		err := dict.Update("NO WORD", newDefinition)

		assertError(t, err, ErrUpdateFailed)
	})
}

func TestAdd(t *testing.T) {

	t.Run("Simple add test", func(t *testing.T) {
		d := Dictionary{}
		want := "addedvalue"
		d.Add("addedkey", want)

		got, err := d.Search("addedkey")

		assertNoError(t, err)

		if got != want {
			t.Errorf("The entry hasen't been added to the map")
		}
	})

	t.Run("Duplicate add test", func(t *testing.T) {
		d := Dictionary{"addedkey": "addedvalue"}
		err := d.Add("addedkey", "addedvalue")

		assertError(t, err, ErrWordExists)

	})
}

func TestDelete(t *testing.T) {

	word := "test"
	dict := Dictionary{word: "test defintion"}

	dict.Delete(word)

	_, err := dict.Search(word)
	if err != ErrWordNotFound {
		t.Errorf("Expected '%s' to be deleted", word)
	}
}

func assertDefinition(t *testing.T, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)

	assertNoError(t, err)

	if got != definition {
		t.Errorf("\nIncorrect defintion\ngot: \"%s\" \nwant: \"%s\"", got, definition)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot: \"%s\" \nwant: \"%s\"", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatalf("Unexpected error occured \n%s", got.Error())
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatalf("Expected an error but didn't get one\nExpected error: \"%s\"", want.Error())
	}

	if got != want {
		t.Errorf("Error type is incorrect \ngot: \"%s\" \nwant: \"%s\"", got.Error(), want.Error())
	}
}
