package base64Captcha

import (
	"image/color"
	"reflect"
	"testing"
)

func TestParseHexColorFast(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantC   color.RGBA
		wantErr bool
	}{
		{name: "Test # error", args: args{"25565"}, wantErr: true, wantC: color.RGBA{}},
		{name: "Test color white", args: args{"#fff"}, wantErr: false, wantC: color.RGBA{255, 255, 255, 255}},
		{name: "Test color red", args: args{"#a02e2e"}, wantErr: false, wantC: color.RGBA{160, 46, 46, 255}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := parseHexColor(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseHexColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("parseHexColor() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
