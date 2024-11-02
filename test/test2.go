package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	dir := "." // replace with the path of the directory you want to inspect

	// Open the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Iterate over the files
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}

		// File details
		fmt.Printf("Name: %s\n", info.Name())
		fmt.Printf("Size: %d bytes\n", info.Size())
		fmt.Printf("Permissions: %s\n", info.Mode())
		fmt.Printf("Last Modified: %s\n", info.ModTime().Format("Jan 2 15:04"))

		// Additional details using syscall
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			fmt.Printf("Inode: %d\n", stat.Ino)
			fmt.Printf("Number of Links: %d\n", stat.Nlink)
		}
		fmt.Println("------")
	}

	// now := time.Now()
	// fmt.Println(now.)
}
