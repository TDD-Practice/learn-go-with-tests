package reflection

import (
	"reflect"
	"testing"
)

/*
func TestWalk(t *testing.T) {

	expected := "Matheus"

	var received []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(s string) {
		received = append(received, s)
	})

	if len(received) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(received), 1)
	}

	if received[0] != expected {
		t.Errorf("expected %q, but got %q", expected, received[0])
	}
}
*/

func TestWalk(t *testing.T) {
	testCases := []struct {
		desc          string
		input         interface{}
		expectedCalls []string
	}{
		{
			desc:          "single field that is 'string' type",
			input:         struct{ Name string }{"Matheus"},
			expectedCalls: []string{"Matheus"},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var received []string
			walk(tC.input, func(s string) {
				received = append(received, s)
			})
			if !reflect.DeepEqual(received, tC.expectedCalls) {
				t.Errorf("received %v, but was expecting %v", received, tC.expectedCalls)
			}
		})
	}
}
