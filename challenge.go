package main

import (
	"encoding/base64"
	"fmt"
)

func reverse(str string) (res string) {
	for _ , i := range str {
		res = string(i) + res
	}
	return
}

func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	
	sd, _ := base64.StdEncoding.DecodeString(secret)

  	fmt.Println("Non-reverse: ", string(sd))
	// Non-reverse:  iangnoW:NAM:ENIL:ta:su:nioJ

	fmt.Println("Final result: ", reverse(string(sd)))
	// Final result:  Join:us:at:LINE:MAN:Wongnai
}