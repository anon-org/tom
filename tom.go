package main

import (
	"fmt"
	tom "github.com/anon-org/tom/querybuilder"
)

func main() {
	str := tom.SelectAll().Field("field1").From("table_1").Build()

	fmt.Println(str)

}
