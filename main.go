package main

func main() {
	cfg := &config{commands: getCommands(),
		Previous: nil,
		page:     0,
	}
	startRepl(cfg)
}
