package aws

import (
	_ec2 "github.com/Superm4n97/aws-operations-poc/aws/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2"
	"k8s.io/klog/v2"
)

// VpcWithInternetGateway creates resources in the following order
//   - VPC
//   - Subnet
//   - Internet Gateway
//   - Attach Internet Gateway to VPC
//   - Update Route entry
func VpcWithInternetGateway(c *ec2.EC2) error {
	//vpc
	vpc, err := _ec2.NewVPC(c)
	if err != nil {
		return err
	}
	klog.Infof("vpc: %s is created", *vpc.VpcId)

	//subnet
	subnet, err := _ec2.NewSubnet(c, vpc.VpcId)
	if err != nil {
		return err
	}
	klog.Infof("subnet: %s is created", *subnet.SubnetId)

	//get route table
	rt, err := _ec2.RouteTable(c, vpc.VpcId)
	if err != nil {
		return err
	}

	//internetGateway
	igw, err := _ec2.NewInternetGateway(c, vpc)
	if err != nil {
		return err
	}
	klog.Infof("internet gateway: %s is created", *igw.InternetGatewayId)

	//route connection
	err = _ec2.NewRoute(c, &_ec2.DestinationOptions{
		GatewayId: igw.InternetGatewayId,
	}, rt.RouteTableId)
	if err != nil {
		return err
	}
	klog.Infof("route inserted to route table %s", *rt.RouteTableId)

	return nil
}
