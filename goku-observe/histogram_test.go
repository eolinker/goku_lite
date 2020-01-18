package goku_observe

import (
	"math"
	"reflect"
	"testing"
)

func TestHistogram_Observe(t *testing.T) {
	type fields struct {
		Buckets []float64
		Count   []int64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want []int64
	}{
		{
			name:   "min",
			fields: fields{
				Buckets: []float64{0.1,0.5,1.0,math.MaxFloat64},
				Count:   []int64{0,0,0,0},
			},
			args:   args{
				value: 0.01,
			},
			want:[]int64{1,1,1,1},
		},
		{
			name:   "lavel",
			fields: fields{
				Buckets: []float64{0.1,0.5,1.0,math.MaxFloat64},
				Count:   []int64{0,0,0,0},
			},

			args:   args{
				value: 0.6,
			},
			want:[]int64{0,0,1,1},
		},
		{
			name:   "over",
			fields:fields{
				Buckets: []float64{0.1,0.5,1.0,math.MaxFloat64},
				Count:   []int64{0,0,0,0},
			},
			args:   args{
				value: 2,
			},
			want:[]int64{0,0,0,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			h := &HistogramObserve{
				Buckets: tt.fields.Buckets,
				Count:   tt.fields.Count,
			}
			h.Observe(tt.args.value)
			got:=h.Collapse()
			if!reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHistogram() = %v, want %v", got, tt.want)
			}
			wantOrg:= []int64{0,0,0,0}
			if!reflect.DeepEqual(h.Count, wantOrg) {
				t.Errorf("NewHistogram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHistogram(t *testing.T) {
	type args struct {
		buckets []float64
	}
	tests := []struct {
		name string
		args args
		want *HistogramObserve
	}{
		// TODO: RegisterDao test cases.
		{
			name: "",
			args: args{
				buckets: []float64{0.1,0.5,1.0},
			},
			want: &HistogramObserver{

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHistogramObserve(len(tt.args.buckets)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHistogram() = %v, want %v", got, tt.want)
			}
		})
	}
}