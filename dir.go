/**
* Author: CZ cz.devnet@gmail.com
 */

package main

import (
	"os"
	"path"
	"syscall"
)

func mkdir(dir string) (err error) {
	if len(dir) < 1 {
		return ErrDirName
	}
	oldMask := syscall.Umask(0)
	defer syscall.Umask(oldMask)
	cssDir := path.Join(dir, "css")
	jsDir := path.Join(dir, "js")
	imgDir := path.Join(dir, "img")
	const dirMode = 0755
	err = os.MkdirAll(dir, dirMode)
	if err != nil {
		err = ErrMkdir
		return
	}

	err = os.MkdirAll(cssDir, dirMode)
	if err != nil {
		err = ErrMkdir
		return
	}

	err = os.MkdirAll(jsDir, dirMode)
	if err != nil {
		err = ErrMkdir
		return
	}

	err = os.MkdirAll(imgDir, dirMode)
	if err != nil {
		err = ErrMkdir
		return
	}

	return
}
