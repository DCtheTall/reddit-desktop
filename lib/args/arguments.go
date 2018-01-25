package args

import (
	"errors"
	"fmt"
)

const (
	// Cache image option
	Cache = "--save"
	// EmptyCache option
	EmptyCache = "--empty"
)

var acceptedOptions = []string{Cache, EmptyCache}

/*
Options provided maps for which option is provided
*/
type Options map[string]bool

func findOptions(searchVals []string, list []string) Options {
	opts := make(Options)
	for _, val := range list {
		for _, searchVal := range searchVals {
			if searchVal == val {
				opts[searchVal] = true
			}
		}
	}
	return opts
}

/*
ParseArgs parse the provided arguments
*/
func ParseArgs(args []string) (subreddits []string, opts Options, err error) {
	opts = findOptions(acceptedOptions, args)

	providedOptions := 0
	for _, optProvided := range opts {
		if optProvided {
			providedOptions++
		}
	}

	// validating
	if providedOptions > 1 {
		return nil, nil, fmt.Errorf("You cannot use the %s and %s options at the same time", Cache, EmptyCache)
	}

	// filtering subreddits from args
	for _, arg := range args {
		if !opts[arg] {
			subreddits = append(subreddits, arg)
		}
	}

	if len(subreddits) == 0 && !opts[EmptyCache] {
		return nil, nil, errors.New("You must provide at least one subreddit")
	} else if len(subreddits) > 0 && opts[EmptyCache] {
		return nil, nil, fmt.Errorf("You cannot provide any subreddits with the %s option", EmptyCache)
	}

	return subreddits, opts, nil
}
