package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"sort"
	"syscall"
)

func ReadAll(path string) ([]fs.FileInfo, error) {
	var Entries []fs.FileInfo
	entrs, err := os.ReadDir(path)
	if err != nil {
		return Entries, errors.New("error reading directory: " + err.Error())
	}

	dot, err := os.Stat(".")
	if err != nil {
		return Entries, err
	}
	ddot, err := os.Stat("..")
	if err != nil {
		return Entries, err
	}
	Entries = append(Entries, dot, ddot)

	for _, entry := range entrs {
		eInfo, err := entry.Info()
		if err != nil {
			return Entries, err
		}
		Entries = append(Entries, eInfo)

	}
	return Entries, nil
}

func listDirectoryRecursively(path string, recursive bool) error {
	// var elements []string

	entries, err := ReadAll(path)
	if err != nil {
		return err
	}

	// PrintResult(path, elements)
	for _, entry := range entries[2:] {
		// fmt.Println(entry.Name())
		if entry.IsDir() && recursive {
			// fmt.Println(i)
			subdir := path + "/" + entry.Name()
			if err := listDirectoryRecursively(subdir, recursive); err != nil {
				return errors.New("Error reading subdirectory:" + err.Error())
			}
		}
	}
	// var stat syscall.Statfs_t
    // syscall.Statfs(path, &stat)
	// println("syscall.Statfs_t Type: ",  stat.Type)
    // println("syscall.Statfs_t Bsize: ",  stat.Bsize)
    // println("syscall.Statfs_t Blocks: ",  stat.Blocks)
	// SortEntriesByModTime(entries)
	// entries = reverseFileInfo(entries)

	printLongFormat(entries)
	// fmt.Println(printLongFormat(entries))
	return nil
}

// printLongFormat prints the file info in a format similar to `ls -l`
func printLongFormat(entries []os.FileInfo) []string {
	var totalBlocks int64
	var longFormatList []string
	for _, entry := range entries {
		// File permissions
		permissions := entry.Mode().String()

		// Number of links
		var numLinks uint64 = 1
		var uid, gid string
		if user, err := user.LookupId(uid); err == nil {
			uid = user.Username
		}
		if group, err := user.LookupGroupId(gid); err == nil {
			gid = group.Name
		}
		if stat, ok := entry.Sys().(*syscall.Stat_t); ok {
			numLinks = uint64(stat.Nlink)
			totalBlocks += stat.Blocks
		}

		// File size
		size := entry.Size()

		// Modification time
		modTime := entry.ModTime().Format("Jan 2 15:04")

		// Print in "ls -l" format
		// fmt.Printf("[%s] [%2d] [%s] [%s] [%d] [%s] [%s]\n", permissions, numLinks, uid, gid, size, modTime, entry.Name())
		longFormatList = append(longFormatList, fmt.Sprintf("%s %2d %s %s %6d %s %s\n", permissions, numLinks, uid, gid, size, modTime, entry.Name()))
	}
	
	totalBlocks /=2
	fmt.Println(totalBlocks)
	return longFormatList
}


func reverseFileInfo(entries []os.FileInfo) []os.FileInfo {
	size := len(entries) - 1
	for i := 0; i < size/2; i++ {
		entries[i], entries[size-i] = entries[size-i], entries[i]
	}
	return entries
}

func SortEntriesByModTime(entries []fs.FileInfo) {
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].ModTime().After(entries[j].ModTime())
	})
}

func PrintResult(path string, elements []string) {
	// if path != "." {
	fmt.Printf("Path= %s:\n", path)

	// }
	for i := 0; i < len(elements); i++ {
		fmt.Printf("%s", elements[i])
		if i < len(elements)-1 {
			fmt.Print(string(13) + string(10))
		}
	}
	fmt.Printf("\n\n")
}

func main() {
	if err := listDirectoryRecursively(".", false); err != nil {
		fmt.Println("Error:", err)
	}
}
