package main
import ("strings"
	"bufio"
	"fmt"
	"os")

func startRepl(){
	scanner := bufio.NewScanner (os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if len(scanner.Text()) == 0 {
			continue
		}
	
		cleanedInput := cleanInput(scanner.Text())
	fmt.Printf("Your command was: %s\n", cleanedInput[0])
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	lowered := strings.ToLower(text)
	stringList := strings.Fields(lowered)	
	return stringList
	}
