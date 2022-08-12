package anytool

import (
	"reflect"
	"testing"
)

func Test_resolvePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				path: "users/0/id",
			},
			want:    []string{"users", "0", "id"},
			wantErr: false,
		},
		{
			name: "error path 1",
			args: args{
				path: "/users/0/id",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error path 2",
			args: args{
				path: "users/0/id/",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolvePath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("resolvePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolvePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		a    any
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "success interface case 1",
			args: args{
				a: map[string]interface{}{
					"users": []map[string]interface{}{
						{
							"id": "12",
						},
					},
				},
				path: "users/0/id",
			},
			want:    "12",
			wantErr: false,
		},

		{
			name: "success interface case 1",
			args: args{
				a: map[string]interface{}{
					"users": []interface{}{
						map[string]interface{}{
							"id": "12",
						},
					},
				},
				path: "users/0/id",
			},
			want:    "12",
			wantErr: false,
		},
		{
			name: "success interface case 2",
			args: args{
				a: map[string]interface{}{
					"users": []interface{}{
						map[string]interface{}{
							"id": "12",
						},
						map[string]interface{}{
							"id": "23",
						},
					},
				},
				path: "users/1/id",
			},
			want:    "23",
			wantErr: false,
		},
		{
			name: "success interface case 3",
			args: args{
				a: []interface{}{
					map[string]interface{}{
						"id": 12,
					},
					map[string]interface{}{
						"id": 23,
					},
				},
				path: "1/id",
			},
			want:    23,
			wantErr: false,
		},
		{
			name: "success interface case 4",
			args: args{
				a: []interface{}{
					12, 33,
				},
				path: "1",
			},
			want:    33,
			wantErr: false,
		},

		{
			name: "success case 1",
			args: args{
				a: map[string][]map[string]string{
					"users": {
						{
							"id": "12",
						},
					},
				},
				path: "users/0/id",
			},
			want:    "12",
			wantErr: false,
		},
		{
			name: "success case 2",
			args: args{
				a: map[string][]map[string]string{
					"users": {
						{
							"id": "12",
						},
						{
							"id": "23",
						},
					},
				},
				path: "users/1/id",
			},
			want:    "23",
			wantErr: false,
		},
		{
			name: "success case 3",
			args: args{
				a: []map[string]int{
					{
						"id": 12,
					},
					{
						"id": 23,
					},
				},
				path: "1/id",
			},
			want:    23,
			wantErr: false,
		},
		{
			name: "success case 4",
			args: args{
				a: []int{
					12, 33,
				},
				path: "1",
			},
			want:    33,
			wantErr: false,
		},

		{
			name: "fail case 1",
			args: args{
				a: []int{
					12, 33,
				},
				path: "-1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail case 2",
			args: args{
				a: []int{
					12, 33,
				},
				path: "id/0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail case 3",
			args: args{
				a: []map[string]int{
					{
						"id": 12,
					},
					{
						"id": 23,
					},
				},
				path: "1/user",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail case 4",
			args: args{
				a: []map[string]int{
					{
						"id": 12,
					},
					{
						"id": 23,
					},
				},
				path: "user/1/id",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.a, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Get(map[string]interface{}{
			"users": []interface{}{
				map[string]interface{}{
					"id": "12",
				},
			},
		}, "users/0/id")
	}
}

func BenchmarkGetSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Get(map[string][]map[string]string{
			"users": {
				{
					"id": "12",
				},
			},
		}, "users/0/id")
	}
}

func BenchmarkGetTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Get(map[string]interface{}{
			"users": []map[string]interface{}{
				{
					"id": "12",
				},
			},
		}, "users/0/id")
	}
}

func BenchmarkGetBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Get(map[string]interface{}{
			"users": []interface{}{
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users01": []interface{}{
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users02": []interface{}{
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users03": []interface{}{
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users04": []interface{}{
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users05": []interface{}{
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				map[string]interface{}{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
		}, "users/0/id")
	}
}

func BenchmarkGetBigSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Get(map[string][]map[string]interface{}{
			"users": {
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users1": {
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users2": {
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users3": {
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users4": {
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users5": {
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
			"users6": {
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
				{
					"id":   "12",
					"name": "xx",
					"age":  30,
				},
			},
		}, "users/0/id")
	}
}

func TestGetString(t *testing.T) {
	type args struct {
		a    any
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				a: map[string]interface{}{
					"A": "s",
				},
				path: "A",
			},
			want:    "s",
			wantErr: false,
		},
		{
			name: "case 2",
			args: args{
				a: map[string]interface{}{
					"A": 1,
				},
				path: "A",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "case 3",
			args: args{
				a: interface{}([]interface{}{
					"s",
				}),
				path: "0",
			},
			want:    "s",
			wantErr: false,
		},
		{
			name: "case 4",
			args: args{
				a: []interface{}{
					1,
				},
				path: "A",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetString(tt.args.a, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	type args struct {
		a    any
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				a: map[string]interface{}{
					"A": 1,
				},
				path: "A",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "case 2",
			args: args{
				a: map[string]interface{}{
					"A": "1",
				},
				path: "A",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "case 3",
			args: args{
				a: []interface{}{
					1,
				},
				path: "0",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "case 4",
			args: args{
				a: []interface{}{
					"A",
				},
				path: "0",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInt(tt.args.a, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
