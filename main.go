package main

import (
	"fmt"

	"github.com/kyoruni/go-internal-test/hoge"
	"github.com/kyoruni/go-internal-test/hoge/hogefuga"
	"github.com/kyoruni/go-internal-test/hoge/hogefuga/hogefugapiyo"
	"github.com/kyoruni/go-internal-test/hoge/hogefuga/hogefugapiyo/hogefugapiyopika"
	"github.com/kyoruni/go-internal-test/hoge/internal/hello"
)

func main() {
	// internalの1つ上の階層と、その配下からは呼び出すことができる
	fmt.Println(hoge.SayHello())             // hello
	fmt.Println(hogefuga.SayHello())         // hello
	fmt.Println(hogefugapiyo.SayHello())     // hello
	fmt.Println(hogefugapiyopika.SayHello()) // hello
	fmt.Println(hello.Say())                 // use of internal package github.com/kyoruni/go-internal-test/hoge/internal/hello not allowed
}
