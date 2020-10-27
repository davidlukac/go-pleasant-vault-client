package test

import (
	. "github.com/davidlukac/go-pleasant-vault-client/pkg/client"
	"testing"
)

func TestObfuscatePassword(t *testing.T) {
	cases := map[string]string{
		"foobarpass": "f*****s",
		"":           "*****",
		"a":          "a*****a",
		"ab":         "a*****b",
	}

	for k, v := range cases {
		t.Logf("Testing for '%s'", k)

		got := ObfuscatePassword(k)

		if got != v {
			t.Errorf("For '%s' got  '%s'; want 'f*****s'!", k, got)
		}

		t.Logf("Got %s", got)
	}
}
