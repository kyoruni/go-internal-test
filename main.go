package main

import (
	"fmt"

	"github.com/kyoruni/go-internal-test/hoge"
	"github.com/kyoruni/go-internal-test/hoge/hogefuga"
	"github.com/kyoruni/go-internal-test/hoge/internal/hello"
)

func main() {
	fmt.Println(hoge.SayHello())     // hello ... internalの直接の親からは呼び出せる
	fmt.Println(hogefuga.SayHello()) // hello ... internakの兄弟からは呼び出せる
	fmt.Println(hello.Say())         // use of internal package github.com/kyoruni/go-internal-test/hoge/internal/hello not allowed ... mainはinternalの直接の親でも兄弟でもないので、呼び出せない
}
