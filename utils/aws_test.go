package utils

import (
	"context"
	"testing"
)

// AWS Settings
func TestNewConfig(t *testing.T) {
	ctx := context.TODO()
	_, err := NewConfig(ctx)
	if err != nil {
		t.Errorf("NewConfig: 戻り値でエラーが発生しています。")
	}
	t.Logf("err is %s", err)
}