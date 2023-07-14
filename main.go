package main

import (
	"fmt"
	"github.com/Superm4n97/aws-operations-poc/aws/service/ec2"
)

func main() {
	err := ec2.NewInstance()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("instance created")
	return
}
