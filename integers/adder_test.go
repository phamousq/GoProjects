// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers

package integers

import "testing"

func TestAdder(t *testing.T){
	sum := Add(2,2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}