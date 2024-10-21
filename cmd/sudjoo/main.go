package main

import (
	"fmt"

	"github.com/vanvanni/sudjoo/internal/config"
	"github.com/vanvanni/sudjoo/internal/logger"
)

func main() {
	config := config.GetConfig()

	logger.Info("Hi!")
	fmt.Println(config)
}
