package prefix

import "testing"

func TestCommon(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name       string
		args       args
		wantPrefix string
	}{
		{
			name:       "列表为空",
			args:       args{[]string{}},
			wantPrefix: "",
		},
		{
			name:       "包含空字符串",
			args:       args{[]string{"prefix", ""}},
			wantPrefix: "",
		},
		{
			name:       "仅英文：包含",
			args:       args{[]string{"prefix", "people"}},
			wantPrefix: "p",
		},
		{
			name:       "仅英文：不包含",
			args:       args{[]string{"name", "people"}},
			wantPrefix: "",
		},
		{
			name:       "仅中文：包含",
			args:       args{[]string{"我是一个人", "我是一个好人"}},
			wantPrefix: "我是一个",
		},
		{
			name:       "仅中文：不包含",
			args:       args{[]string{"我是一个人", "他是一个好人"}},
			wantPrefix: "",
		},
		{
			name:       "中英文混合：包含",
			args:       args{[]string{"我是me一个人", "我是me"}},
			wantPrefix: "我是me",
		},
		{
			name:       "中英文混合：不包含",
			args:       args{[]string{"me我是一个人", "他是一个好人"}},
			wantPrefix: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix := Common(tt.args.list)
			if gotPrefix != tt.wantPrefix {
				t.Errorf("Common() = %v, want %v", gotPrefix, tt.wantPrefix)
			}
			t.Log(gotPrefix)
		})
	}
}
