package utils

import (
	"context"
	"log"
	"testing"
)

func TestGetCost(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	ctx := context.TODO()
	client, err := NewConfig(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}

	start, end := GetRange()
	total, err := GetCost(ctx, client, "ap-northeast-1", start, end)
	if err != nil {
		t.Errorf("GetCost: 戻り値でエラーが発生しています。")
		t.Logf("err is %s", err)
		return
	}

	if !(0 <= total) {
		t.Errorf("total: 値が正しくありません。")
	}
	t.Logf("total is %v", total)
}