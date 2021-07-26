package main

import (
	"encoding/json"
	"log"

	"google.golang.org/protobuf/proto"
)

func encode() []byte {
	var hello int32
	hello = 22
	isMarried := false

	p := Person{
		Name:          "Buford",
		Age: &hello,
		IsMarried: &isMarried,
	}

	out, err := proto.Marshal(&p)
	if err != nil {
		log.Fatalln("failed to encode: ", err)
	}

	return out
}

func decode(protobuf []byte) string {

	p := Person{}
	if err := proto.Unmarshal(protobuf, &p); err != nil {
		log.Fatalln("failed to parse", err)
	}
	
	out, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(out)
}
