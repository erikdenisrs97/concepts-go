package main

import (
	"fmt"

	"github.com/erikdenisrs97/tests/addresses"
)

func main() {
	tipoEndereco := addresses.AddressType("Avenue Paulista")
	fmt.Println(tipoEndereco)
}
