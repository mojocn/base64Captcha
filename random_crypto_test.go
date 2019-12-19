// Copyright 2017 Eric Zhou. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base64Captcha

import (
	"reflect"
	"testing"
)

func Test_deriveSeed(t *testing.T) {
	type args struct {
		purpose byte
		id      string
		digits  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantOut [16]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := deriveSeed(tt.args.purpose, tt.args.id, tt.args.digits); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("deriveSeed() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestRandomId(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomId(); got != tt.want {
				t.Errorf("RandomId() = %v, want %v", got, tt.want)
			}
		})
	}
}
