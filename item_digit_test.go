package base64Captcha

import (
	"bytes"
	"testing"
)

func TestNewItemDigit(t *testing.T) {
	type args struct {
		width    int
		height   int
		dotCount int
		maxSkew  float64
	}
	tests := []struct {
		name string
		args args
		want *ItemDigit
	}{
		{"one", args{240, 80, 6, 0.8}, nil},
		{"one", args{240, 80, 5, 0.8}, nil},
		{"one", args{240, 80, 6, 0.8}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItemDigit(tt.args.width, tt.args.height, tt.args.dotCount, tt.args.maxSkew); got == nil {
				t.Errorf("NewItemDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemDigit_EncodeBinary(t *testing.T) {

	idd := NewItemDigit(80, 300, 20, 0.25)

	tests := []struct {
		name string
		m    *ItemDigit
		want []byte
	}{
		{"one", idd, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.EncodeBinary(); len(got) == 0 {
				t.Errorf("ItemDigit.EncodeBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemDigit_WriteTo(t *testing.T) {
	tests := []struct {
		name    string
		m       *ItemDigit
		want    int64
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			got, err := tt.m.WriteTo(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemDigit.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ItemDigit.WriteTo() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ItemDigit.WriteTo() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestItemDigit_EncodeB64string(t *testing.T) {
	idd := NewItemDigit(80, 300, 20, 0.25)

	tests := []struct {
		name string
		m    *ItemDigit
		want string
	}{
		{"", idd, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.EncodeB64string(); got == tt.want {
				t.Errorf("ItemDigit.EncodeB64string() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemDigit_DotCountZero(t *testing.T) {
	_ = NewItemDigit(80, 300, 0, 0.25)
}
