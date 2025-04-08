package generator

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	t.Run("Test Generator", func(t *testing.T) {
		got, err := Generate()
		AssertGenerator(t, got, err)
	})
}

func AssertGenerator(t testing.TB, got string, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
	if got == "" {
		t.Error()
	}
}
