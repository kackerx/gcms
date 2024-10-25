package main

import "gcms/internal/conf"

func main() {
	cfg, err := conf.New()
	if err != nil {
		panic(err)
	}

	app, f, err := wireApp(cfg.Data)
	defer f()
	if err != nil {
		panic(err)
	}

	app.Run()
}
