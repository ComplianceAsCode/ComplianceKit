package clifactory

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
)

type ReaderConfig struct {
	Regex     string
	Multiline bool
	Prompt    bool
}

// CliReader handles getting input from command line interactively
func CliReader(msg string, val string, options ...ReaderConfig) string {
	opts := ReaderConfig{
		Regex:     "",
		Multiline: false,
		Prompt:    false,
	}

	for _, opt := range options {
		opts = ReaderConfig(opt)
	}

	if val == "" || opts.Prompt {
		var scan_text string
		fmt.Printf(" %s: ", msg)
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			if !opts.Multiline {
				return scanner.Text()
			} else if scanner.Text() == "done" {
				return scan_text
			} else {
				scan_text += scanner.Text() + "\n"
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	if len(opts.Regex) > 0 {
		re := regexp.MustCompile(opts.Regex)
		test := re.MatchString(val)

		if !test {
			fmt.Println("ERROR: ", test)
			log.Fatal("Something failed but I'm not quitting.")

		}
	}
	return val
}
