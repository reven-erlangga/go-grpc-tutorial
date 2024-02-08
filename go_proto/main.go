package main

import (
	"fmt"
	pb "go_proto/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Nike Black T-Shirt",
				Price: 10000.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    1,
				Name:  "Nike Air Hordan",
				Price: 10000.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   2,
					Name: "Shoe",
				},
			},
		},
	}

	data, err := proto.Marshal(products)

	if err != nil {
		log.Fatal("Marshall error", err)
	}

	fmt.Println(data) // Encode with compact binary wire format

	decodeProducts := &pb.Products{}
	if err = proto.Unmarshal(data, decodeProducts); err != nil {
		log.Fatal("Unmarshall error", err)
	}

	fmt.Println(decodeProducts) // Decode from compact binary wire format

	for _, product := range decodeProducts.GetData() {
		fmt.Println(product.GetName())
		fmt.Println(product.GetCategory().GetName())
	}
}
