package main

import (
	"fmt"
	"github.com/RHEcosystemAppEng/sbo-go-library/pkg/binding/convert"
)

func main() {
	connString, err := convert.GetMongoDBConnectionString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(connString)
	connString, err = convert.GetPostgreSQLConnectionString()
	if err != nil {
		panic(err)
	}
	fmt.Println(connString)
}
