package main

import (
	"context"
	"log"

	"github.com/tqer39/aws-billing-go/utils"
)

func main() {
	// AWSの設定
	ctx := context.TODO()
	client, err := utils.NewConfig(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// 取得する期間
	start, end := utils.GetRange()
	log.Printf("start=%s, end=%s", start, end)

	// 現在の使用料金を取得
	total, err := utils.GetCost(ctx, client, "ap-northeast-1", start, end)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("total=%v", total)
}