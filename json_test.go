package gtool

import (
	"reflect"
	"testing"
	"time"
)

func TestJSONTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		t      time.Time
		args   args
		hasErr bool
	}{
		{
			"正常時間",
			time.Date(2020, time.October, 14, 12, 33, 10, 0, time.UTC),
			args{[]byte(`"2020-10-14 12:33:10"`)},
			false,
		},
		{
			"空字串",
			time.Time{},
			args{[]byte(`""`)},
			false,
		},
		{
			"錯誤時間",
			time.Time{},
			args{[]byte(`"test"`)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := JSONTime{}
			if err := test.UnmarshalJSON(tt.args.data); (err != nil) != tt.hasErr {
				t.Errorf("JSONTime.UnmarshalJSON() error = %v, hasErr %v", err, tt.hasErr)
				return
			}
			if !reflect.DeepEqual(test.Time(), tt.t) {
				t.Errorf("JSONTime.UnmarshalJSON() JSONTime = %v, wantJSONTime %v", test.Time(), tt.t)
			}
		})
	}
}

func TestJSONTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		t      time.Time
		want   string
		hasErr bool
	}{
		{
			"正常時間",
			time.Date(2020, time.October, 14, 12, 33, 10, 0, time.UTC),
			`"2020-10-14 12:33:10"`,
			false,
		},
		{
			"空字串",
			time.Time{},
			`""`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := JSONTime(tt.t)
			got, err := test.MarshalJSON()
			if (err != nil) != tt.hasErr {
				t.Errorf("JSONTime.MarshalJSON() error = %v, hasErr %v", err, tt.hasErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("JSONTime.MarshalJSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestJsonToIntList(t *testing.T) {
	type args struct {
		jsonStr string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 bool
	}{
		{
			"json decode []int",
			args{`[1,2,3]`},
			[]int{1, 2, 3},
			true,
		},
		{
			"json decode []int fail",
			args{`[1,"b",3]`},
			[]int{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := JsonToIntList(tt.args.jsonStr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToIntList() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("JsonToIntList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestJsonToStringList(t *testing.T) {
	type args struct {
		jsonStr string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 bool
	}{
		{
			"json decode []string",
			args{`["a","b","c"]`},
			[]string{"a", "b", "c"},
			true,
		},
		{
			"json decode []string fail",
			args{`{"1":"abc"}`},
			[]string{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := JsonToStringList(tt.args.jsonStr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToStringList() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("JsonToStringList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
