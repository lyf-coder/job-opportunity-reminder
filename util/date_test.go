package util

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	nowTime := time.Now()
	actualResult := GetTimeFormat(nowTime, DATE)
	if !strings.Contains(actualResult, strconv.Itoa(nowTime.Year())) {
		t.Errorf("实际结果：%v", actualResult)
	}

}

func TestGetCurrent(t *testing.T) {
	nowTime := time.Now()
	actualResult := GetCurrent(DATE)
	if !strings.Contains(actualResult, strconv.Itoa(nowTime.Year())) {
		t.Errorf("实际结果：%v", actualResult)
	}
}

func TestGetTime(t *testing.T) {
	result, err := GetTime("2016-11-11", DATE)
	if err != nil {
		t.Error(err)
	}
	expectResult := 2016
	if result.Year() != expectResult {
		t.Errorf("实际结果: %v 期望结果：%v ", result.Year(), expectResult)
	}
}

func TestTimeStrInDuration(t *testing.T) {
	type args struct {
		d       time.Duration
		timeStr string
		format  FORMAT
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "五分钟内",
			args: args{
				d:       5 * time.Minute,
				timeStr: GetTimeFormat(time.Now().In(CstZone).Add(-3*time.Minute), DATETIME),
				format:  DATETIME,
			},
			want: true,
		},
		{
			name: "超过五分钟",
			args: args{
				d:       5 * time.Minute,
				timeStr: GetTimeFormat(time.Now().In(CstZone).Add(-6*time.Minute), DATETIME),
				format:  DATETIME,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeStrInDuration(tt.args.d, tt.args.timeStr, tt.args.format); got != tt.want {
				t.Errorf("TimeStrInDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
