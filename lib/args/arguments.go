package args

import (
	"fmt"
)

const (
	// Cache image option - saves the scraped image
	Cache = "--save"
	// EmptyCache option - deletes saved images
	EmptyCache = "--empty"
	// Undo option - sets desktop to most previous image (only if caching)
	Undo = "--undo"
)

var acceptedOptions = []string{Cache, EmptyCache, Undo}

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

	if len(subreddits) > 0 && opts[EmptyCache] {
		return nil, nil, fmt.Errorf("You cannot provide any subreddits with the %s option", EmptyCache)
	} else if len(subreddits) > 0 && opts[Undo] {
		return nil, nil, fmt.Errorf("You cannot provide any subreddits with the %s option", Undo)
	} else if len(subreddits) == 0 && opts[Cache] {
		return nil, nil, fmt.Errorf("You must provide at least one subreddit with the %s option", Cache)
	} else if len(subreddits) == 0 && providedOptions == 0 {
		return nil, nil, fmt.Errorf("You must provide at least one subreddit")
	}
	return subreddits, opts, nil
}
