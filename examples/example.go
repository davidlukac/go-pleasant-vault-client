package examples

import (
	"fmt"
	"github.com/davidlukac/go-pleasant-vault-client/pkg/client"
	"testing"
)

func Test(t *testing.T) {
	c := client.NewClient("https://vault.foobar.com", "some.username", "asdf3ih43")
	s := c.GetSecret("b29031b3-3951-41c9-b0bf-06fc69366559")
	fmt.Println(s)
}
