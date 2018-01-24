package args

import (
	"errors"
	"strconv"
	"strings"
)

var acceptedOptions = []string{"--cache", "--limit"}

type option struct {
	provided bool
	index    int
}

func findOptions(searchVals []string, list []string) (result map[string]option) {
	for i, val := range list {
		for _, searchVal := range searchVals {
			if searchVal == val {
				result[searchVal] = option{provided: true, index: i}
			}
		}
	}
	return result
}

func getCacheLimitWithUnits(limit string) (int, string, error) {
	kbIndex := strings.Index("kb", limit)
	mbIndex := strings.Index("mb", limit)
	gbIndex := strings.Index("gb", limit)
	unit := ""
	strippedSizeStr := ""
	if kbIndex != -1 {
		unit = "kb"
		strippedSizeStr = string([]byte(limit)[:kbIndex])
	} else if mbIndex != -1 {
		unit = "mb"
		strippedSizeStr = string([]byte(limit)[:mbIndex])
	} else if gbIndex != -1 {
		unit = "gb"
		strippedSizeStr = string([]byte(limit)[:gbIndex])
	} else {
		unit = "files"
		strippedSizeStr = limit
	}
	size, err := strconv.Atoi(strippedSizeStr)
	if err != nil {
		return 0, "", errors.New("Invalid value after --limit option")
	}
	return size, unit, nil
}

/*
ParseArgs parse the provided arguments
*/
func ParseArgs(args []string) (subreddits []string, err error) {
	options := findOptions(acceptedOptions, args)

	// validating
	switch true {
	case options["--limit"].provided && !options["--cache"].provided:
		return nil, errors.New("You must provide --cache option in order to specify a limit")
	case options["--limit"].index == len(args)-1:
		return nil, errors.New("You must supply an argument after --limit")
	}

	if options["--limit"].provided {
		indexAfterLimit := options["limit"].index + 1
		sizeAsStr := strings.ToLower(args[indexAfterLimit])
		_, _, err := getCacheLimitWithUnits(sizeAsStr)
		if err != nil {
			return nil, err
		}
	}

	return args, nil
}
