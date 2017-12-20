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
	"testing"
)

func TestSetGet(t *testing.T) {
	s := NewMemoryStore(GCLimitNumber, Expiration)
	id := "captcha id"
	d := randText(10, TxtAlphabet)
	s.Set(id, d)
	d2 := s.Get(id, false)
	if d2 == "" || d2 != d {
		t.Errorf("saved %v, getDigits returned got %v", d, d2)
	}
}

func TestGetClear(t *testing.T) {
	s := NewMemoryStore(GCLimitNumber, Expiration)
	id := "captcha id"
	d := randText(10, TxtAlphabet+TxtNumbers)
	s.Set(id, d)
	d2 := s.Get(id, true)
	if d2 == "" || d != d2 {
		t.Errorf("saved %v, getDigitsClear returned got %v", d, d2)
	}
	d2 = s.Get(id, false)
	if d2 != "" {
		t.Errorf("getDigitClear didn't clear (%q=%v)", id, d2)
	}
}

func TestCollect(t *testing.T) {
	//TODO(dchest): can't test automatic collection when saving, because
	//it's currently launched in a different goroutine.
	s := NewMemoryStore(10, -1)
	// create 10 ids
	ids := make([]string, 10)
	d := randText(10, TxtAlphabet+TxtNumbers)
	for i := range ids {
		ids[i] = randomId()
		s.Set(ids[i], d)
	}
	s.(*memoryStore).collect()
	// Must be already collected
	nc := 0
	for i := range ids {
		d2 := s.Get(ids[i], false)
		if d2 != "" {
			t.Errorf("%d: not collected", i)
			nc++
		}
	}
	if nc > 0 {
		t.Errorf("= not collected %d out of %d captchas", nc, len(ids))
	}
}

func BenchmarkSetCollect(b *testing.B) {
	b.StopTimer()
	d := randText(10, TxtAlphabet+TxtNumbers)
	s := NewMemoryStore(9999, -1)
	ids := make([]string, 1000)
	for i := range ids {
		ids[i] = randomId()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			s.Set(ids[j], d)
		}
		s.(*memoryStore).collect()
	}
}
