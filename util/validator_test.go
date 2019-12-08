package util

import "testing"

func TestIsStringExist(t *testing.T) {

	var dataSource = []string{
		"hola",
		"amigo",
		"hey",
		"hai",
	}

	type args struct {
		source     string
		dataSource []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Success",
			args: args{
				source:     "hai",
				dataSource: dataSource,
			},
			want: true,
		},
		{
			name: "Test Failed",
			args: args{
				source:     "hoy",
				dataSource: dataSource,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringExist(tt.args.source, tt.args.dataSource); got != tt.want {
				t.Errorf("IsStringExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateInput(t *testing.T) {
	var languageSource = []string{
		"id", "en", "ja", "kr",
	}

	type args struct {
		sourceLanguage string
		destLanguage   string
		dataSource     []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Success",
			args: args{
				sourceLanguage: "id",
				destLanguage:   "en",
				dataSource:     languageSource,
			},
			want: true,
		},
		{
			name: "Test Failed - Source Language Not Found",
			args: args{
				sourceLanguage: "zn",
				destLanguage:   "en",
				dataSource:     languageSource,
			},
			want: false,
		},
		{
			name: "Test Failed - Destination Language Not Found",
			args: args{
				sourceLanguage: "id",
				destLanguage:   "woi",
				dataSource:     languageSource,
			},
			want: false,
		},
		{
			name: "Test Failed - Invalid Input",
			args: args{
				sourceLanguage: "hey",
				destLanguage:   "woi",
				dataSource:     languageSource,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateInput(tt.args.sourceLanguage, tt.args.destLanguage, tt.args.dataSource); got != tt.want {
				t.Errorf("ValidateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
