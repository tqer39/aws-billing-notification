package utils

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

// AWSの利用料金を取得する関数
func GetCost(ctx context.Context, client *costexplorer.Client, region string, start string, end string) (float64, error) {
	params := &costexplorer.GetCostAndUsageInput{
		Granularity: types.GranularityMonthly,
		// Metrics:     []string{"UNBLENDED_COST"},
		Metrics:     []string{"AMORTIZED_COST"},
		TimePeriod:  &types.DateInterval{Start: aws.String(start), End: aws.String(end)},
		GroupBy: []types.GroupDefinition{{
			Key:  aws.String("SERVICE"),
			Type: types.GroupDefinitionTypeDimension,
		},
		},
	}

	resp, err := client.GetCostAndUsage(ctx, params)
	if err != nil {
		log.Fatalf("%v", err)
		return 0, err
	}

	var amount float64
	for _, val := range resp.ResultsByTime {
		for _, g := range val.Groups {
			for _, m := range g.Metrics {
				f, err := strconv.ParseFloat(*m.Amount, 64)
				if err != nil {
					return 0, fmt.Errorf("AWS CostExplorer returns illegal amount -> %w", err)
				}
				amount += f
			}
		}
	}
	return math.Ceil(amount*100) / 100, nil
}