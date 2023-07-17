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

	//new instance
	resrv, err := ec2.NewInstance(c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("available instances are:")
	for _, ins := range resrv.Instances {
		fmt.Println(*ins.InstanceId)
	}

	//delete instance
	//err = ec2.DeleteInstances(c, "i-03a93be9c0813dc15")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println("instance successfully deleted")

	return
}
