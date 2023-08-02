package main

import "github.com/patyumi/go-expert/8manipeventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World - From Producer")
}
