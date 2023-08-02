package main

import (
	"fmt"
	"github.com/Superm4n97/aws-operations-poc/aws"
	"k8s.io/klog/v2"
)

/*
instance create order
	* VPC
	* Subnet
	* KeyPairs
	* Instance
*/

func main() {
	c, err := aws.EC2Client()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = aws.VpcWithInternetGateway(c)
	if err != nil {
		klog.Errorf(err.Error())
	}

	return
}
