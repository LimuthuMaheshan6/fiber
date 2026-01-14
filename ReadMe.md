package main

import "log"

type con struct {
	
}
func (c *con) hello() {

}

func main() {
	log.Println("Air Works... 31")


	var c con

	c.hello()
	

}

go run github.com/99designs/gqlgen generate







