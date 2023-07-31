package main

import (
	"fmt"
	"github.com/Superm4n97/aws-operations-poc/aws"
	_ec2 "github.com/Superm4n97/aws-operations-poc/aws/service/ec2"
)

/*
instance create order
	* VPC
	* Subnet
	* KeyPairs
	* Instance
*/

/*
vpc with internet gateway with order
  - VPC
  - Subnet
  - Internet Gateway
  - Route
*/
func main() {
	c, err := aws.EC2Client()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//vpc
	_, err = _ec2.NewVPC(c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	/*
		//subnet
		_, err = _ec2.NewSubnet(c, vpc.VpcId)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//get route table
		rt, err := _ec2.RouteTable(c, vpc.VpcId)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//internetGateway
		igw, err := _ec2.NewInternetGateway(c)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//route connection
		err = _ec2.NewRoute(c, &_ec2.DestinationOptions{
			GatewayId: igw.InternetGatewayId,
		}, rt.RouteTableId)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	*/
	return
}
