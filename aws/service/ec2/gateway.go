package ec2

import "github.com/aws/aws-sdk-go/service/ec2"

func NewInternetGateway(c *ec2.EC2, vpc *ec2.Vpc) (*ec2.InternetGateway, error) {
	out, err := c.CreateInternetGateway(&ec2.CreateInternetGatewayInput{})
	if err != nil {
		return nil, err
	}
	return out.InternetGateway, attachInternetGatewayToVPC(c, out.InternetGateway, vpc)
}

func RemoveInternetGateway(c *ec2.EC2, internetGatewayID *string) error {
	_, err := c.DeleteInternetGateway(&ec2.DeleteInternetGatewayInput{
		DryRun:            nil,
		InternetGatewayId: internetGatewayID,
	})
	return err
}

func attachInternetGatewayToVPC(c *ec2.EC2, igw *ec2.InternetGateway, vpc *ec2.Vpc) error {
	_, err := c.AttachInternetGateway(&ec2.AttachInternetGatewayInput{
		InternetGatewayId: igw.InternetGatewayId,
		VpcId:             vpc.VpcId,
	})
	return err
}
