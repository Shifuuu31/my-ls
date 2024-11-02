package parser

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"my-ls/source"
)

var ArgCount int

func ParseArgs(ls *source.MyLsFlag, In *source.Inputs) error {
	if ArgCount == 0 {
		return nil
	} else {
		for _, arg := range os.Args[1:] {
			if arg == "--help" {
				ls.Help = true
				return nil
			} else if isValidFlag(arg) {
				err := getOption(ls, arg)
				if err != nil {
					return err
				}
			} else {
				if err := checkPath(In, arg); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func isValidFlag(flag string) bool {
	if strings.Count(flag, "-") == 1 && flag[0] == '-' {
		return true
	}
	return false
}

func getOption(ls *source.MyLsFlag, argument string) error {
	for _, c := range argument[1:] {
		// fmt.Println(string(c))
		switch c {
		case 'a':
			ls.Flag_a = true
		case 'R':
			ls.Flag_R = true
		case 't':
			ls.Flag_t = true
		case 'r':
			ls.Flag_r = true
		case 'l':
			ls.Flag_l = true
		default:
			return errors.New("invalid flag: " + argument)

		}
	}
	return nil
}

func checkPath(In *source.Inputs, path string) error {
    if In == nil {
        return errors.New("Inputs struct is not initialized")
    }

    fileInfo, err := os.Stat(path)
    if os.IsNotExist(err) {
        fmt.Println(path)
        return errors.New("my-ls: cannot access '" + path + "': No such file or directory")
    }
    if err != nil {
        return err
    }

    if fileInfo.IsDir() {
        In.Dir = append(In.Dir, path)
    } else {
        In.File = append(In.File, path)
    }
    return nil
}
