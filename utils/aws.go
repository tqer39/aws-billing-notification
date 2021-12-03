package utils

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

// AWS Settings
func NewConfig(ctx context.Context) (*costexplorer.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("ap-northeast-1"),
		config.WithSharedConfigProfile("private-aws"),
	)
	if err != nil {
		log.Printf("%v\n", err)
		cfg, err = config.LoadDefaultConfig(ctx,
			// Hard coded credentials.
			config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
				Value: aws.Credentials{
					AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
					SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
					Source:          "Profile Based Credentials",
				},
			}),
		)
		if err != nil {
			log.Printf("%v\n", err)
			return nil, err
		}
	}
	client := costexplorer.NewFromConfig(cfg)
	if err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}
	return client, nil
}