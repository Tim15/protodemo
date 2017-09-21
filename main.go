package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/tim15/protodemo/api/shopping"
)

func main() {
	fmt.Println("Starting app!")
	pb := shopping.Item{Name: "juice", Type: shopping.Item_DAIRY}
	fmt.Printf("%+v\n", pb)
	m := jsonpb.Marshaler{}
	// m.EmitDefaults = true
	str, err := m.MarshalToString(&pb)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(str)
}
