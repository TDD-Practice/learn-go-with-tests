/* package main

import "testing"

func TestMain(m *testing.M) {

}
*/

package concurrency

import (
	"reflect"
	"testing"
)

func stubUrlChecker(url string) (result bool) {
	switch url {
	case "/true":
		result = true
	case "/flase":
		result = false
	}
	return result
}

func TestCheckUrl(t *testing.T) {
	got := CheckUrl(stubUrlChecker, []string{"/true", "/false"})
	want := map[string]bool{"/true": true, "/false": false}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
