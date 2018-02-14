// UNIX哲学に出来る限り従う方向で！
// isearとホームページを作りたい
// lazyの構成を参考に
package main

import (
	"flag"
	"net/http"

	_ "github.com/intelfike/web/contents"
	"github.com/intelfike/web/policy"
)

func init() {
	flag.Parse()

	http.HandleFunc("/", policy.Listener)
}
