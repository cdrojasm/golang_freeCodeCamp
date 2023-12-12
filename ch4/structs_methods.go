package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	authString := fmt.Sprintf("Authorization: Basic %s:%s", authInfo.username, authInfo.password)
	return authString
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("===============================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
