package reflection

import (
	"reflect"
	"testing"
)

func TestTest(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Nik"},
			[]string{"Nik"},
		},
		{
			"Strict with two string fields",
			struct {
				Name string
				City string
			}{"Nik", "Tokyo"},
			[]string{"Nik", "Tokyo"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Nik", 26},
			[]string{"Nik"},
		},
		{
			"Nested fields",
			Person{
				"Nik",
				Profile{26, "Tokyo"},
			},
			[]string{"Nik", "Tokyo"},
		},
		{
			"Pointers to things",
			&Person{
				"Nik",
				Profile{26, "Tokyo"},
			},
			[]string{"Nik", "Tokyo"},
		},
		{
			"Slices",
			[]Profile{
				{26, "Tokyo"},
				{25, "Brno"},
			},
			[]string{"Tokyo", "Brno"},
		},
		{
			"Arrays",
			[2]Profile{
				{26, "Tokyo"},
				{25, "Brno"},
			},
			[]string{"Tokyo", "Brno"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("With maps", func(t *testing.T) {
		aMap := map[string]string{
			"foo": "bar",
			"baz": "qox",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "qox")
	})

	t.Run("With channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{25, "Tokyo"}
			aChannel <- Profile{24, "Brno"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Tokyo", "Brno"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("With function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{25, "Tokyo"}, Profile{24, "Brno"}
		}

		var got []string
		want := []string{"Tokyo", "Brno"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
