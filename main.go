package main

func main() {
	cfg := &config{commands: getCommands(),
		page: 0,
	}
	startRepl(cfg)
}
