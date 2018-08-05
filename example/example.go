package main

import (
	"fmt"

	"github.com/kzyn/azure-sas"
)

func main() {
	uri := "<uri>"
	key := "<primary key>"
	expiry := 3600

	token, err := sas.Generate(uri, key, int64(expiry))
	if err != nil {
		panic(err)
	}

	fmt.Printf("SAS Token+ %s \n", token)
}
