package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	shop "github.com/zangar-tm/online_shop"
)

func main() {
	fmt.Println("asdf")
	srv := new(shop.Server)
	if err := srv.Run("8000"); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
