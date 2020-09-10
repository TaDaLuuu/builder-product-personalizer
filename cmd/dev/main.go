package main

import "builder/internal/config"

func main() {
	config.Replace()
	//config.StartVue()
	config.BuildVue()
}
