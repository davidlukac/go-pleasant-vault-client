package test

import (
	. "github.com/davidlukac/go-pleasant-vault-client/pkg/client"
	. "github.com/onsi/gomega"
	"testing"
)

func TestParseUUID(t *testing.T) {
	g := NewGomegaWithT(t)

	cases := map[string]string{
		"foobar": "foobar",
		"https://vault.ferratum.com:12000/WebClient/Main?itemId=bazqux": "bazqux",
		"htt\\ps:/vault.srambled.url/Main?itemId=bazqux":                "htt\\ps:/vault.srambled.url/Main?itemId=bazqux",
	}

	for k, v := range cases {
		t.Logf("Testing for '%s'", k)

		got := ParseUUID(k, "")

		if got != v {
			t.Errorf("For '%s' got '%s'; want '%s'!", k, got, v)
		}

		t.Logf("Got %s", got)
	}

	g.Expect(func() { ParseUUID("https://vault.ferratum.com:12000/WebClient/Main?foo=bazqux", "") }).To(Panic())

	g.Expect(ParseUUID("https://vault.ferratum.com:12000/WebClient/Main?foo=bazqux", "foo")).To(Equal("bazqux"))
}
