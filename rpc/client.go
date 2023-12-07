package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
)

func main()  {
	client, err := rpc.Dial("tcp", "localhost:13133")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		var replay bool
		err = client.Call("Listener.GetLine", line, &replay)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(replay)
	}
}
