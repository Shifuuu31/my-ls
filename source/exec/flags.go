package exec

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
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


	entries, err := ReadAll(path)
	if err != nil {
		return err
	}

	for _, entry := range entries[2:] {

		if entry.IsDir() && recursive {

			subdir := path + "/" + entry.Name()
			if err := listDirectoryRecursively(subdir, recursive); err != nil {
				return errors.New("Error reading subdirectory:" + err.Error())
			}
		}
	}
	// SortEntriesByModTime(entries)
	// entries = reverseFileInfo(entries)

	printLongFormat(entries)
	// fmt.Println(printLongFormat(entries))
	return nil
}

func SortEntriesByModTime(entries []fs.FileInfo) {
	n := len(entries)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if entries[j].ModTime().Before(entries[j+1].ModTime()) {
				entries[j], entries[j+1] = entries[j+1], entries[j]
			}
		}
	}
}

func reverseFileInfo(entries []os.FileInfo) []os.FileInfo {
	size := len(entries) - 1
	for i := 0; i < size/2; i++ {
		entries[i], entries[size-i] = entries[size-i], entries[i]
	}
	return entries
}

func printLongFormat(entries []os.FileInfo) []string {
	var totalBlocks int64
	var longFormatList []string
	for _, entry := range entries {
		permissions := entry.Mode().String()

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

		size := entry.Size()

		modTime := entry.ModTime().Format("Jan 2 15:04")

		longFormatList = append(longFormatList, fmt.Sprintf("%s %2d %s %s %6d %s %s\n", permissions, numLinks, uid, gid, size, modTime, entry.Name()))
	}

	totalBlocks /= 2
	fmt.Println(totalBlocks)
	return longFormatList
}
