package source

import "io/fs"

type (
	MyLsFlag struct {
		Flag_a bool
		Flag_R bool
		Flag_t bool
		Flag_r bool
		Flag_l bool
		Help   bool
	}

	Inputs struct {
		Dir  []string
		File []string
	}

	// Entry struct {
	// 	Path         string
	// 	Files_Hidden []fs.FileInfo
	// 	Dirs_Hidden  []fs.FileInfo
	// 	Files        []fs.FileInfo
	// 	Dirs         []fs.FileInfo
	// }
)

var (
	Flags *MyLsFlag
	In    *Inputs
	Entries []fs.FileInfo
)
