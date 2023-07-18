package main

import (
	"fmt"
	"github.com/Superm4n97/aws-operations-poc/aws"
	self_ec2 "github.com/Superm4n97/aws-operations-poc/aws/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type info struct {
	vpc      *ec2.Vpc
	subnet   *ec2.Subnet
	keypair  *ec2.KeyPairInfo
	instance *ec2.Instance
}

var temp info

func createVM(c *ec2.EC2) error {
	vpc, err := self_ec2.NewVPC(c)
	if err != nil {
		return err
	}
	fmt.Println("VPC created with id: ", *vpc.VpcId)

	subnet, err := self_ec2.NewSubnet(c, *vpc.VpcId)
	if err != nil {
		return err
	}
	fmt.Println("subnet created with id: ", *subnet.SubnetId)

	keypair, err := self_ec2.NewKeyPair(c, "sdk-test")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("keypair created with name: ", *keypair.KeyName)

	ins, err := self_ec2.NewInstance(c, *keypair.KeyName, *subnet.SubnetId)
	if err != nil {
		return err
	}
	fmt.Println("instance created with id:", *ins.InstanceId)
	return nil
}

func deleteVM(c *ec2.EC2) error {
	err := self_ec2.RemoveInstances(c, *temp.instance.InstanceId)
	if err != nil {
		return err
	}
	fmt.Println("instance deleted")

	err = self_ec2.RemoveKeyPair(c, *temp.keypair.KeyName)
	if err != nil {
		return err
	}
	fmt.Println("keypair deleted")

	err = self_ec2.RemoveSubnet(c, *temp.subnet.SubnetId)
	if err != nil {
		return err
	}
	fmt.Println("Subnet deleted")

	err = self_ec2.RemoveVPC(c, *temp.vpc.VpcId)
	if err != nil {
		return err
	}
	fmt.Println("VPC deleted")
	return nil
}

func main() {
	c, err := aws.EC2Client()
	if err != nil {
		fmt.Println(err.Error())
	}

	for true {
		var s string
		_, err = fmt.Scan(&s)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if s == "exit" {
			break
		}

		if s == "create machine" {
			err = createVM(c)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
		}

		if s == "delete machine" {
			err = deleteVM(c)
		}
	}

	return
}
