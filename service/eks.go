package service

import (
	"context"
	"fmt"
	"os/exec"

	"caas-eks-api-go/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

var client *eks.Client

func init() {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client = eks.NewFromConfig(cfg)
}

func CreateCluster(req models.CreateClusterRequest) (*types.Cluster, error) {
	out, err := client.CreateCluster(context.TODO(), &eks.CreateClusterInput{
		Name:    aws.String(req.Name),
		RoleArn: aws.String(req.RoleArn),
		ResourcesVpcConfig: &types.VpcConfigRequest{
			SubnetIds:        req.SubnetIds,
			SecurityGroupIds: req.SecurityGs,
		},
		Version: aws.String(req.Version),
	})
	if err != nil {
		return nil, err
	}
	return out.Cluster, nil
}

func ListClusters() ([]string, error) {
	resp, err := client.ListClusters(context.TODO(), &eks.ListClustersInput{})
	if err != nil {
		return nil, err
	}
	return resp.Clusters, nil
}

func GetCluster(name string) (*types.Cluster, error) {
	resp, err := client.DescribeCluster(context.TODO(), &eks.DescribeClusterInput{
		Name: aws.String(name),
	})
	if err != nil {
		return nil, err
	}
	return resp.Cluster, nil
}

func DeleteCluster(name string) error {
	_, err := client.DeleteCluster(context.TODO(), &eks.DeleteClusterInput{
		Name: aws.String(name),
	})
	return err
}

func DeployApp(clusterName string) error {
	kubeconfig := fmt.Sprintf("/tmp/%s-kubeconfig.yaml", clusterName)
	updateCmd := exec.Command("aws", "eks", "update-kubeconfig", "--name", clusterName, "--kubeconfig", kubeconfig)
	if out, err := updateCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("kubeconfig failed: %v\n%s", err, out)
	}
	apply := exec.Command("kubectl", "apply", "-f", "deployment.yaml", "--kubeconfig", kubeconfig)
	if out, err := apply.CombinedOutput(); err != nil {
		return fmt.Errorf("kubectl apply failed: %v\n%s", err, out)
	}
	return nil
}
