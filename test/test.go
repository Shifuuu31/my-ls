package main

import (
    "errors"
    "fmt"
    "os"
    "time"
)

// listDirectoryRecursively recursively lists files and directories.
func listDirectoryRecursively(path string, showAll, longFormat bool) error {
    fmt.Println(path + ":")

    entries, err := os.ReadDir(path)
    if err != nil {
        return errors.New("error reading directory: " + err.Error())
    }

    for _, entry := range entries {
        // Skip hidden files unless `showAll` (-a) is true
        if !showAll && entry.Name()[0] == '.' {
            continue
        }

        // Handle `-l` flag (long format)
        if longFormat {
            info, err := entry.Info()
            if err != nil {
                continue // Skip files with errors
            }

            modTime := info.ModTime().Format(time.RFC822)
            fmt.Printf("%-10s %10d %s %s\n", info.Mode(), info.Size(), modTime, entry.Name())
        } else {
            fmt.Println(entry.Name())
        }
    }

    // Recursively visit directories if `-R` is enabled
    for _, entry := range entries {
        if entry.IsDir() {
            subdir := path + "/" + entry.Name()
            err := listDirectoryRecursively(subdir, showAll, longFormat)
            if err != nil {
                fmt.Println("Error reading subdirectory:", err)
            }
        }
    }
    return nil
}

func main() {
    // Simulate flag values for demonstration
    showAll := true
    longFormat := true

    if err := listDirectoryRecursively(".", showAll, longFormat); err != nil {
        fmt.Println("Error:", err)
    }
}
