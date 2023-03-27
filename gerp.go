package main

// NOTE: important: this is a joke command of course ^^

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// Liste de blagues
	jokes := []string{
		"Not the valid grep command!",
		"(╯°□°）╯︵ ┻━┻",
		"(ﾉಥ益ಥ）ﾉ ┻━┻",
		"(／‵Д′)／~ ┻━┻",
		"┻━┻︵ \\(°□°)/ ︵ ┻━┻",
	}

	// Init random
	rand.Seed(time.Now().UnixNano())

	// Choose joke
	joke := jokes[rand.Intn(len(jokes))]

	// Print joke on stderr
	fmt.Fprintf(os.Stderr, "\033[31m%s\033[0m\n", joke)

	// get args and run them (in a shell as the gerp will be in a shell too)
	args := os.Args[1:]
	cmd := exec.Command("/bin/sh", "-c", "grep "+strings.Join(args, " "))

	// Mapping de l'entrée standard vers le stdin de la commande "grep"
	cmd.Stdin = os.Stdin

	// Map stdin and err for grep to be ok
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Go!
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erreur in the grep command :", err)
		os.Exit(1)
	}
}
