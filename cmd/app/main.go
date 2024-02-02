package main

import (
	"flag"
	"wblzero/internal/app"
)

func main() {
	cfgPath := flag.String("cfg", "./.env", "USAGE -cfg='path_to_config_file'")
	flag.Parse()
	app.RunServer(*cfgPath)
}
