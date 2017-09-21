package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/tim15/protodemo/api/shopping"
)

func main() {
	fmt.Println("Starting app!")

	// Initialize some data
	pb := shopping.Item{Name: "juice", Type: shopping.Item_DAIRY}
	fmt.Printf("%+v\n", pb)

	// Create marshaler
	m := jsonpb.Marshaler{}
	// m.EmitDefaults = true

	// Marshall from proto-struct to JSON string
	str, err := m.MarshalToString(&pb)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(str)

	// All done!
}
