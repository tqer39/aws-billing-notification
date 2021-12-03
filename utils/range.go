package utils

import "time"

// 期間を取得
func GetRange() (string, string) {
	now := time.Now()
	nowyyyyMMdd := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	start := nowyyyyMMdd.Format("2006-01-02")
	end := now.Format("2006-01-02")
	return start, end
}