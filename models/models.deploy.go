package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

func GetDeployments(account CloudAccount) []types.DeploymentInfo {
	var keyId string
	var secretKey string
	if account == HntCloud {
		keyId = os.Getenv("AWS_ACCESS_KEY_ID_HNT")
		secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY_HNT")
	} else if account == Fnd {
		keyId = os.Getenv("AWS_ACCESS_KEY_ID_FND")
		secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY_FND")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(keyId, secretKey, "")))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := codedeploy.NewFromConfig(cfg)
	resp, err := client.ListDeployments(context.TODO(), &codedeploy.ListDeploymentsInput{})
	if err != nil {
		log.Fatalf("failed to list deployments, %v", err)
	}

	var deps []types.DeploymentInfo
	fmt.Println("Deployments:")
	for i, deployment := range resp.Deployments {
		if i >= 20 {
			break
		}
		dep := getDeployment(client, deployment)
		deps = append(deps, *dep)
		fmt.Println(i, deployment, *dep.ApplicationName, dep.Status)
	}
	fmt.Println(resp.NextToken)
	return deps
}

func getDeployment(client *codedeploy.Client, id string) *types.DeploymentInfo {
	resp, err := client.GetDeployment(context.TODO(), &codedeploy.GetDeploymentInput{
		DeploymentId: &id,
	})
	if err != nil {
		log.Fatalf("failed to get deployment, %v", err)
	}
	return resp.DeploymentInfo
}
