# go-akeneo

> NOTE: This is a work in progress and must not be used in production

## Usage

```
package main

import (
	"fmt"
	"io/ioutil"

	goakeneo "github.com/go-akeneo/go-akeneo"
)

func main() {
	cb := goakeneo.NewClientBuilder("https://acme.cloud.akeneo.com/")
	pimClient := cb.BuildAuthenticatedByPassword("clientId", "secret", "username", "password")

	res, err := pimClient.GetProductApi().Get("SKU123")
	if err != nil {
		panic(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resBody))
}
```
