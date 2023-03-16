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
