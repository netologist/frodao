package nullable_test

import (
	"testing"

	"github.com/hasanozgan/frodao/nullable"
)

func Test_Nullable(t *testing.T) {
	var got nullable.Type[string] = nullable.New("lorem")

	if got.TypeValue != "lorem" {
		t.Errorf("nullable.Type[string] = %v; want lorem", got.TypeValue)
	}

	got.Scan("ipsum")

	if value, err := got.Value(); err != nil {
		t.Errorf("Failed %v", err)
	} else if value != "ipsum" {
		t.Errorf("nullable.Type[string] = %v; want ipsum", got.TypeValue)
	}
}

func Test_Nullable_WhenEmpty(t *testing.T) {
	var got nullable.Type[string]

	if value, err := got.Value(); err != nil {
		t.Errorf("Failed %v", err)
	} else if value != nil {
		t.Errorf("nullable.Type[string] it should be null %v", value)
	}
}
