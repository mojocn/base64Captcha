package base64Captcha

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func Test_parseDigitsToString(t *testing.T) {
	for i := 1; i < 10; i++ {
		digti := randomDigits(i)
		s := parseDigitsToString(digti)
		if len(s) != i {
			t.Error("failed")
		}

	}

}

func Test_stringToFakeByte(t *testing.T) {
	for i := 1; i < 10; i++ {
		digti := randomDigits(i)
		s := parseDigitsToString(digti)
		if len(s) != i {
			t.Error("failed")
		}
		fb := stringToFakeByte(s)
		if !reflect.DeepEqual(fb, digti) {
			t.Error("failed")
		}
	}
}

func Test_randomDigits(t *testing.T) {
	for i := 1; i < 10; i++ {
		digti := randomDigits(i)
		if len(digti) != i {
			t.Error("failed")
		}

	}
}

func Test_randomBytes(t *testing.T) {
	for i := 1; i < 10; i++ {
		digti := randomBytes(i)
		if len(digti) != i {
			t.Error("failed")
		}
	}
}

func Test_randomBytesMod(t *testing.T) {
	for i := 1; i < 10; i++ {
		digti := randomBytesMod(i, 'c')
		if len(digti) != i {
			t.Error("failed")
		}
	}
}

func Test_itemWriteFile(t *testing.T) {
	//todo:::
}

func Test_pathExists(t *testing.T) {
	td := os.TempDir()
	defer os.RemoveAll(td)
	p := filepath.Join(td, RandomId())
	if pathExists(p) {
		t.Error("failed")
	}
	_ = os.MkdirAll(p, os.ModePerm)

	if !pathExists(p) {
		t.Error("failed")
	}
}
