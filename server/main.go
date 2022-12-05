package main

import (
	"github.com/tuki9ko/cola_aketa/route"
)

func main() {
	r := route.GetRouter()
	r.Run(":8080")
}
