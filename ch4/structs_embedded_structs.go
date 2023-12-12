package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("============================================")
}

func main() {
	test(sender{
		rateLimit: 100000,
		user: user{
			name:   "Deborah",
			number: 981273981,
		},
	})
	test(sender{
		rateLimit: 41231,
		user: user{
			name:   "Sarah",
			number: 981272233981,
		},
	})
	test(sender{
		rateLimit: 2312,
		user: user{
			name:   "Sally",
			number: 981123273981,
		},
	})
}
