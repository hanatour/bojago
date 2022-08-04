package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

func GetDeployments(account CloudAccount) []types.DeploymentInfo {
	var keyId string
	var secretKey string

	fmt.Println(account)
	keyId = os.Getenv("AWS_ACCESS_KEY_ID_" + account.String())
	secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY_" + account.String())

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
	ch := make(chan types.DeploymentInfo)
	fmt.Println("Deployments:")
	for i, deployment := range resp.Deployments {
		if i >= 20 {
			break
		}
		go getDeployment(client, deployment, ch)
	}
	for i, _ := range resp.Deployments {
		if i >= 20 {
			break
		}
		dep := <-ch
		//fmt.Println(*dep.ApplicationName, dep.Status)
		deps = append(deps, dep)
		sort.Slice(deps, func(i, j int) bool {
			return deps[i].CreateTime.After(*deps[j].CreateTime)
		})
	}
	return deps
}

func getDeployment(client *codedeploy.Client, id string, ch chan types.DeploymentInfo) {
	resp, err := client.GetDeployment(context.TODO(), &codedeploy.GetDeploymentInput{
		DeploymentId: &id,
	})
	if err != nil {
		log.Fatalf("failed to get deployment, %v", err)
	}
	ch <- *resp.DeploymentInfo
}
