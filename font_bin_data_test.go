package base64Captcha

import (
	"reflect"
	"testing"
)

func TestAsset(t *testing.T) {

	for _, value := range _bindata {

		_, err := value()
		if err != nil {
			t.Error(err)
		}

	}

}

func Test_bindata_read(t *testing.T) {
	type args struct {
		data []byte
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bindata_read(tt.args.data, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("bindata_read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bindata_read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_3dumb_ttf(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_3dumb_ttf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_3dumb_ttf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_3dumb_ttf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_apothecaryfont_ttf(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_apothecaryfont_ttf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_apothecaryfont_ttf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_apothecaryfont_ttf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_comismsh_ttf(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_comismsh_ttf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_comismsh_ttf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_comismsh_ttf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_dennethree_dee_ttf(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_dennethree_dee_ttf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_dennethree_dee_ttf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_dennethree_dee_ttf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_deborahfancydress_ttf(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_deborahfancydress_ttf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_deborahfancydress_ttf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_deborahfancydress_ttf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_flim_flam_ttf(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_flim_flam_ttf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_flim_flam_ttf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_flim_flam_ttf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_ritasmith_ttf(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_ritasmith_ttf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_ritasmith_ttf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_ritasmith_ttf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_actionj_ttf(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_actionj_ttf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_actionj_ttf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_actionj_ttf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_chromohv_ttf(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_chromohv_ttf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_chromohv_ttf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_chromohv_ttf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_readme_md(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_readme_md()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_readme_md() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_readme_md() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts_wqy_microhei_ttc(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts_wqy_microhei_ttc()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts_wqy_microhei_ttc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts_wqy_microhei_ttc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssetNames(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssetNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssetDir(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssetDir(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssetDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
