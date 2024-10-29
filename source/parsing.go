package source

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ParseArgs(ls *MyLsFlag) error {
	argCount := len(os.Args[1:])
	if argCount == 0 {
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
				if err := checkPath(arg); err != nil {
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

func getOption(ls *MyLsFlag, argument string) error {
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

func checkPath(path string) error {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println(path)
		return errors.New("my-ls: cannot access '" + path + "': No such file or directory")
	}
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		Dir = append(Dir, path)
	} else {
		File = append(File, path)
	}
	return nil
}
