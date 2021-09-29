package gtool

import (
	"reflect"
	"testing"
)

func TestWorkManager(t *testing.T) {
	tests := []struct {
		name        string
		max         int // 同時執行上限
		runCount    int // 執行次數
		finishCount int // 執行結束數量
		want        int
	}{
		{
			"執行數量上限10_執行1次_完成1次",
			10,
			1,
			1,
			0,
		},
		{
			"執行數量上限10_執行5次_完成4次",
			10,
			5,
			4,
			1,
		},
		{
			"執行數量上限10_執行10次_完成3次",
			10,
			10,
			3,
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWorkManager(tt.max)
			for i := 0; i < tt.runCount; i++ {
				w.Start()
			}
			for i := 0; i < tt.finishCount; i++ {
				w.End()
			}
			if got := w.Count(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WorkManager.Count() 執行中數量 = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkManager_Wait(t *testing.T) {
	tests := []struct {
		name string
		max  int // 同時執行上限
	}{
		{
			"執行數量上限1",
			1,
		},
		{
			"執行數量上限10",
			10,
		},
		{
			"執行數量上限100",
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWorkManager(tt.max)
			for i := 0; i < tt.max; i++ {
				w.Start()
				w.End()
			}

			w.Wait()
		})
	}
}
