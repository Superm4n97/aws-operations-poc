package ec2

import (
	"errors"
	"github.com/Superm4n97/aws-operations-poc/utils"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	allIPs = "0.0.0.0/0"
)

type DestinationOptions struct {
	CarrierGatewayId            *string
	EgressOnlyInternetGatewayId *string
	GatewayId                   *string
	InstanceId                  *string
	LocalGatewayId              *string
	NatGatewayId                *string
	NetworkInterfaceId          *string
	RouteTableId                *string
	TransitGatewayId            *string
	VpcPeeringConnection        *string
	VpcEndpointId               *string
}

func NewRoute(c *ec2.EC2, desOps *DestinationOptions, routeTableId *string) error {
	_, err := c.CreateRoute(&ec2.CreateRouteInput{
		CarrierGatewayId:            desOps.CarrierGatewayId,
		DestinationCidrBlock:        utils.StringP(allIPs),
		EgressOnlyInternetGatewayId: desOps.EgressOnlyInternetGatewayId,
		GatewayId:                   desOps.GatewayId,
		InstanceId:                  desOps.InstanceId,
		LocalGatewayId:              desOps.LocalGatewayId,
		NatGatewayId:                desOps.NatGatewayId,
		NetworkInterfaceId:          desOps.NetworkInterfaceId,
		RouteTableId:                routeTableId,
		TransitGatewayId:            desOps.TransitGatewayId,
		VpcEndpointId:               desOps.VpcEndpointId,
		VpcPeeringConnectionId:      desOps.VpcPeeringConnection,
	})
	return err
}

func RouteTable(c *ec2.EC2, vpcId *string) (*ec2.RouteTable, error) {
	out, err := c.DescribeRouteTables(&ec2.DescribeRouteTablesInput{})
	if err != nil {
		return nil, err
	}
	for _, rt := range out.RouteTables {
		if rt.VpcId == vpcId {
			return rt, nil
		}
	}
	return nil, errors.New("no route table found")
}
