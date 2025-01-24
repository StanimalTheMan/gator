package main

import (
	"fmt"

	"github.com/StanimalTheMan/gator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("stanimal")
	fmt.Println(config.Read())
}
