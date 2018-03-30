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
	"fmt"
)

var providers = make(map[string]Store)

// Store An object implementing Store interface can be registered with SetCustomStore
// function to handle storage and retrieval of captcha ids and solutions for
// them, replacing the default memory store.
//
// It is the responsibility of an object to delete expired and used captchas
// when necessary (for example, the default memory store collects them in Set
// method after the certain amount of captchas has been stored.)
type Store interface {
	// Set sets the digits for the captcha id.
	Set(id string, value string)

	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) string

	InitStore(sc storeConfig) error
}

type storeConfig struct {
	expire int64    	// 单位：秒
	limitNumber int		// 最大个数
	extraConfig string  //额外配置信息，如redis连接等
}

func NewGlobalStore(storeType string, config storeConfig) (Store, error){
	provider, ok := providers[storeType]
	if  !ok{
		return nil, fmt.Errorf("store: unknown provide %q (forgotten import?)", storeType)
	}
	error := provider.InitStore(config)
	return provider, error
}

func Register(name string, store Store){
	if store == nil {
		panic("store: Register store is nil")
	}
	if _, dup := providers[name]; dup {
		panic("store: Register called twice for store " + name)
	}
	providers[name] = store
}

