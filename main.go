package main

import (
	"blog_Agregator/internal/config"
	"fmt"
)

func main() {
	cfg := config.Read()
	cfg.SetUser()
	fmt.Printf("%v\n", cfg)

}
