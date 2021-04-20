package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	listName:="saulgoodman"
	uid:=110000283
	a:=base64.StdEncoding.EncodeToString([]byte(listName))

	listName = base64.StdEncoding.EncodeToString([]byte(listName))
	k:=fmt.Sprintf("%s_%d", a, uid)
	fmt.Println(k)
}
