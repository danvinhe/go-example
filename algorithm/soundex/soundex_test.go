package soundex

import "testing"

func TestValue(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		{
			name:       "中国",
			args:       args{"中国"},
			wantResult: "",
			wantErr:    true,
		},
		{
			name:       "helindan",
			args:       args{"helindan"},
			wantResult: "H453",
			wantErr:    false,
		},
		{
			name:       "HeLinDan",
			args:       args{"HeLinDan"},
			wantResult: "H453",
			wantErr:    false,
		},
		{
			name:       "HELINDAN",
			args:       args{"HELINDAN"},
			wantResult: "H453",
			wantErr:    false,
		},
		{
			name:       "Tynczak",
			args:       args{"Tynczak"},
			wantResult: "T522",
			wantErr:    false,
		},
		{
			name:       "Pfister",
			args:       args{"Pfister"},
			wantResult: "P236",
			wantErr:    false,
		},
		{
			name:       "Knuth",
			args:       args{"Knuth"},
			wantResult: "K530",
			wantErr:    false,
		},
		{
			name:       "Kant",
			args:       args{"Kant"},
			wantResult: "K530",
			wantErr:    false,
		},
		{
			name:       "Jarovski",
			args:       args{"Jarovski"},
			wantResult: "J612",
			wantErr:    false,
		},
		{
			name:       "Resnik",
			args:       args{"Resnik"},
			wantResult: "R252",
			wantErr:    false,
		},
		{
			name:       "Reznick",
			args:       args{"Reznick"},
			wantResult: "R252",
			wantErr:    false,
		},
		{
			name:       "Euler",
			args:       args{"Euler"},
			wantResult: "E460",
			wantErr:    false,
		},
		{
			name:       "Peterson",
			args:       args{"Peterson"},
			wantResult: "P362",
			wantErr:    false,
		},
		{
			name:       "Jefferson",
			args:       args{"Jefferson"},
			wantResult: "J162",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Value(tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("Value() = %v, want %v", gotResult, tt.wantResult)
			}
			t.Log(gotResult)
		})
	}
}
