package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	client := eks.NewFromConfig(cfg)

	input := &eks.CreateClusterInput{
		Name:    aws.String("caas-cluster"),
		RoleArn: aws.String("arn:aws:iam::123456789012:role/EKSClusterRole"),
		ResourcesVpcConfig: &types.VpcConfigRequest{
			SubnetIds:        []string{"subnet-abc123", "subnet-def456"},
			SecurityGroupIds: []string{"sg-01234"},
		},
		Version: aws.String("1.27"),
	}

	resp, err := client.CreateCluster(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cluster creation initiated. Status:", resp.Cluster.Status)
}
