/**
* Author: CZ cz.devnet@gmail.com
 */

package main

import (
	"errors"
)

var (
	//ErrDirName is a dir error
	ErrDirName = errors.New("A unavaliable directory name!")

	//ErrMkdir is a make error
	ErrMkdir = errors.New("Make Directory error!")
)
