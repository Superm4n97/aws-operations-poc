package ec2

import (
	"github.com/Superm4n97/aws-operations-poc/utils/convert"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func NewSubnet(c *ec2.EC2, vpcID *string) (*ec2.Subnet, error) {
	out, err := c.CreateSubnet(&ec2.CreateSubnetInput{
		CidrBlock: convert.StringP(vpcCIDR),
		VpcId:     vpcID,
	})
	if err != nil {
		return nil, err
	}
	return out.Subnet, nil
}

func RemoveSubnet(c *ec2.EC2, subnetID *string) error {
	_, err := c.DeleteSubnet(&ec2.DeleteSubnetInput{
		SubnetId: subnetID,
	})
	return err
}
