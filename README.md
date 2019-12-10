# Base64captcha is a customize-friendly captcha package.
[![Go Report Card](https://goreportcard.com/badge/github.com/mojocn/base64Captcha)](https://goreportcard.com/report/github.com/mojocn/base64Captcha)
[![GoDoc](https://godoc.org/github.com/mojocn/base64Captcha?status.svg)](https://godoc.org/github.com/mojocn/base64Captcha)
[![Build Status](https://travis-ci.org/mojocn/base64Captcha.svg?branch=master)](https://travis-ci.org/mojocn/base64Captcha)
[![codecov](https://codecov.io/gh/mojocn/base64Captcha/branch/master/graph/badge.svg)](https://codecov.io/gh/mojocn/base64Captcha)
![stability-stable](https://img.shields.io/badge/stability-stable-brightgreen.svg)
[![codebeat badge](https://codebeat.co/badges/650029a5-fcea-4416-925e-277e2f178e96)](https://codebeat.co/projects/github-com-mojocn-base64captcha-master)
[![Foundation](https://img.shields.io/badge/Golang-Foundation-green.svg)](http://golangfoundation.org)

Base64captcha supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.
Base64Captcha is used for fast development of RESTful APIs, web apps and backend services in Go.
give a string identifier to the package and it returns with a base64-encoding-png-string

#### Why Base64 for RESTful Application
      Data URIs are now supported by all major browsers. IE supports embedding images since version 8 as well.
      RESTful Application returns small base64 image is more convenient.

#### Documentation

* [English](https://godoc.org/github.com/mojocn/base64Captcha)
* [中文文档](https://github.com/mojocn/base64Captcha/blob/master/README_zh.md)
* [Playground](https://captcha.mojotv.cn)


## 2. Quick Start

### 2.1 Download package
    go get -u github.com/mojocn/base64Captcha
For Gopher from mainland China without VPN `go get golang.org/x/image` failure solution:
- go version > 1.11
- set env `GOPROXY=https://goproxy.io`

### 2.2 How to code with base64Captcha

#### 2.2.1 Impelement [Store interface](/base64Captcha/blob/master/interface_store.go) 
```go
type Store interface {
	// Set sets the digits for the captcha id.
	Set(id string, value string)

	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) string
	//
	Verify(id, answer string, clear bool) bool
}

```

#### 2.2.2 Impelement [Driver interface](/base64Captcha/blob/master/interface_driver.go)
```go
// Driver captcha interface for captcha engine to to write staff
type Driver interface {
	// EncodeBinary covert to bytes
	GenerateItem(content string) (item Item, err error)
	GenerateQuestionAnswer() (q, a string)
}
```

#### 2.2.3 New [Captcha instance]((/base64Captcha/blob/master/captcha.go))
```go

func init() {
	//init rand seed
	rand.Seed(time.Now().UnixNano())
}

// Captcha captcha basic information.
type Captcha struct {
	Driver Driver
	Store  Store
}

func NewCaptcha(driver Driver, store Store) *Captcha {
	return &Captcha{Driver: driver, Store: store}
}

func (c *Captcha) Generate() (id, b64s string, err error) {
	id = randomId()
	content, answer := c.Driver.GenerateQuestionAnswer()
	item, err := c.Driver.GenerateItem(content)
	if err != nil {
		return "", "", err
	}
	c.Store.Set(id, answer)
	b64s = item.EncodeB64string()
	return
}
//if you has multiple captcha instances which shared a same store. You may want to use `store.Verify` method instead.
//Verify by given id key and remove the captcha value in store, return boolean value.
func (c *Captcha) Verify(id, answer string, clear bool) (match bool) {
	match = c.Store.Get(id, clear) == answer
	return
}

```
#### 2.2.4 Generate Base64(image/audio) string
```go
func (c *Captcha) Generate() (id, b64s string, err error) {
	id = randomId()
	content, answer := c.Driver.GenerateQuestionAnswer()
	item, err := c.Driver.GenerateItem(content)
	if err != nil {
		return "", "", err
	}
	c.Store.Set(id, answer)
	b64s = item.EncodeB64string()
	return
}
```
#### 2.2.5 Verify Answer
```go
//if you has multiple captcha instances which shared a same store. You may want to use `store.Verify` method instead.
//Verify by given id key and remove the captcha value in store, return boolean value.
func (c *Captcha) Verify(id, answer string, clear bool) (match bool) {
	match = c.Store.Get(id, clear) == answer
	return
}
```

#### 2.2.6 Full Example

```go
func main(){
		dr := DefaultDriverAudio
    	st := DefaultMemStore
    	c := NewCaptcha(dr,st)
    	id,b64s,err := c.Generate()
    	if err != nil {
    		panic(err)
    	}
    	fmt.Println("ID ",id)
    	fmt.Println("Base64 Png or Wav ",b64s)
    	
    	//verify
    	//st.Verify(id,"xxxx",true)
    	//c.Verify(id,"xxxx",true)
}


```

#### 2.2.7 More Example


### func SetCustomStore

	func SetCustomStore(s Store)

SetCustomStore sets custom storage for captchas, replacing the default
memory store. This function must be called before generating any captchas.
##### [Customize Redis Store](_examples_redis/main.go)
### func NewMemoryStore

	func NewMemoryStore(collectNum int, expiration time.Duration) Store

NewMemoryStore returns a new standard memory store for captchas with the
given collection threshold and expiration time in seconds. The returned
store must be registered with SetCustomStore to replace the default one.
## use base64Captcha quick start a API server
```go

```
#### base64Captcha package function

#### Build and Run the Demo
    cd $GOPATH/src/github.com/mojocn/captcha/examples
    go run main.go

Congratulations! You've just built your first **base64Captcha-APIs** app.
Any question you can leave a message. If you like the package please star this repo
## Thanks

- [dchest/captha](https://github.com/dchest/captcha)
- [@slayercat][https://github.com/slayercat]
- [@amzyang][https://github.com/amzyang]
- [@Luckyboys](https://github.com/Luckyboys)

## License

base64Captcha source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).
