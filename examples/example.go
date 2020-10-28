package examples

import (
	"fmt"
	"github.com/davidlukac/go-pleasant-vault-client/pkg/client"
	"testing"
)

func Test(t *testing.T) {
	c := client.NewClient("https://vault.foobar.com", "some.username", "asdf3ih43")
	s := c.GetSecret("b29031b3-a1dd-q2wq-1as4e-06fc69366559")
	fmt.Println(s)
}
