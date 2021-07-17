package examples

import (
	"fmt"
	"github.com/davidlukac/go-pleasant-vault-client/pkg/client"
	"testing"
)

// Test example Password Server client.
func Test(t *testing.T) {
	c := client.Vault{URL: "https://vault.foobar.com", Username: "some.username", Password: "1sdf3ih43"}
	s := c.GetSecret("b29031b3-a1dd-q2wq-1as4e-06fc69366559")
	fmt.Println(s)
}
