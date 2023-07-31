package ec2

import "github.com/aws/aws-sdk-go/service/ec2"

func NewInternetGateway(c *ec2.EC2) (*ec2.InternetGateway, error) {
	out, err := c.CreateInternetGateway(&ec2.CreateInternetGatewayInput{})
	if err != nil {
		return nil, err
	}
	return out.InternetGateway, nil
}

func RemoveInternetGateway(c ec2.EC2, internetGatewayID *string) error {
	_, err := c.DeleteInternetGateway(&ec2.DeleteInternetGatewayInput{
		DryRun:            nil,
		InternetGatewayId: internetGatewayID,
	})
	return err
}
