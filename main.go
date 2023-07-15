package main

import (
	"fmt"
	"github.com/Superm4n97/aws-operations-poc/aws"
	"github.com/Superm4n97/aws-operations-poc/aws/service/ec2"
)

func main() {
	c, err := aws.EC2Client()
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ec2.NewInstance(c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
