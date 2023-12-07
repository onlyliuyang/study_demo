package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/testProject/protobufMarshal/protobuf"
	"log"
)

func main() {
	p1 := pb.Person{
		Id:   *proto.Int32(1),
		Name: *proto.String("liuyang"),
	}

	p2 := pb.Person{
		Id:   2,
		Name: "gopher",
	}

	all_p := pb.AllPerson{
		Per: []*pb.Person{&p1, &p2},
	}

	//对数据序列化
	data, err := proto.Marshal(&all_p)
	if err != nil {
		log.Fatalln("Mashal data error:", err)
	}

	//对已结序列化的数据反序列化
	var target pb.AllPerson
	err = proto.Unmarshal(data, &target)
	if err != nil {
		log.Fatalln("UnMashal data error:", err)
	}
	fmt.Println(target.Per[0])
}
