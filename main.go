package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/vavarodrigues/test-aws-sdk-go-v2/provider"
)

func main() {
	a := provider.AWS{EC2: ec2Client(configClient()), RDS: rdsClient(configClient())}
	result, _ := a.FetchSubnet()
	fmt.Println(*result.DBSubnetGroups[0].DBSubnetGroupName)
}

func configClient() aws.Config {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	return cfg
}

func ec2Client(config aws.Config) *ec2.EC2 {
	return ec2.New(config)
}

func rdsClient(config aws.Config) *rds.RDS {
	return rds.New(config)
}
