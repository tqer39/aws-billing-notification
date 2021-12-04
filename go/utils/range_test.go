package utils

import (
	"regexp"
	"testing"
)

func TestGetRange(t *testing.T) {
	start, end := GetRange()
	targets := []struct{
		Key   string
		Value string
	}{
		{Key: "start", Value: start},
		{Key: "end", Value: end},
	}

	re := regexp.MustCompile(`^[0-9]{4}\-[0-9]{2}\-[0-9]{2}`)
	for _, target := range targets {
		if !re.MatchString(target.Value) {
			t.Errorf("%s: 期間のフォーマットが正しくありません。", target.Key)
		}
		t.Logf("%s is %s", target.Key, target.Value)
	}
}