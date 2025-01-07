package parser

import (
	"fmt"
	"gowikicli/internal/utils"
	"strings"

	"github.com/BenSnedeker/GoWiki/gowiki"
)

// Parse command and arguments then handle command
func HandleInput(args []string) error {
	// Checks if any arguments were entered
	if len(args) <= 1 {
		return fmt.Errorf("no command detected, try `wiki help`")
	}

	// Grabs the first argument and checks what command it is
	command := strings.ToLower(args[1])
	switch command {

	case "help":
		fmt.Printf("Help!")

	case "search":
		// If there aren't any arguments
		if len(args[2:]) == 0 {
			return fmt.Errorf("search command needs at least 1 argument")
		}

		// Parse the search arguments
		query := parseSearch(args[2:])

		// Call API using query
		gowiki.HandleSearch(query)

	case "read":
		// If there aren't any arguments
		if len(args[2:]) == 0 {
			return fmt.Errorf("read command needs at least 1 argument")
		}

		// Parse the read arguments
		flags, query, err := parseRead(args[2:]) /// SOLVE ALTER parseRead
		if err != nil {
			return fmt.Errorf("failed to parse read arguments\n     : %s", err)
		}

		// Call API using flags and query
		gowiki.HandleRead(flags, query)

	default:
		return fmt.Errorf("command not recognized -> \"%s\"", command)
	}

	return nil
}

// -------------
//    FLAGS
// -------------

func parseFlags(flag_args []string) (gowiki.Flags, error) {
	flags := gowiki.NewFlags()

	// For each flag arg, alter flags or throw error
	for _, arg := range flag_args {
		fmt.Printf("Checking arg: \"%s\"\n", arg)
		switch arg {
		// Text flags
		case "--html", "-h", "--text", "-t":
			flags = utils.UpdateTextFlags(arg, flags)
		// Section flags
		case "--all", "-a", "--limited`", "-l", "--summary", "-s", "--references", "-r", "--wikilinks", "-w":
			flags = utils.UpdateSectionFlags(arg, flags)
		// Visual flags
		case "--clean", "-c", "--fancy", "-f":
			flags = utils.UpdateVisualFlags(arg, flags)
		// No recognized flag present
		default:
			return flags, fmt.Errorf("unrecognized flag \"%s\"", arg)
		}
	}

	return flags, nil
}

// -------------
//    SEARCH
// -------------

func parseSearch(args []string) string {
	return strings.Join(args, " ")
}

// -------------
//     READ
// -------------

func parseRead(args []string) (gowiki.Flags, string, error) {
	var flag_args []string
	query := ""

	// For each argument, find all flags and find query term
	for i, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flag_args = append(flag_args, arg)
		} else {
			query = strings.Join(args[i:], " ")
			break
		}
	}

	// Parse flag args
	flags, err := parseFlags(flag_args)
	if err != nil {
		return flags, "", fmt.Errorf("error parsing flags\n     : %s", err)
	}

	// Check if query is present
	if query == "" {
		return flags, "", fmt.Errorf("no article name found")
	}

	return flags, query, nil
}
