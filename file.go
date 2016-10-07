/**
* Author: CZ cz.devnet@gmail.com
 */

package main

import (
	"os"
	"path"
)

func createIdxHTML(dir string) (err error) {
	fp := path.Join(dir, "index.html")
	fd, err := os.Create(fp)
	if err != nil {
		return
	}
	defer fd.Close()
	_, err = fd.Write([]byte(idxHTML))
	return
}

func createIdxCSS(dir string) (err error) {
	fp := path.Join(dir, "css", "index.css")
	fd, err := os.Create(fp)
	if err != nil {
		return
	}
	defer fd.Close()
	_, err = fd.Write([]byte(idxCSS))
	return
}

func createImpressJS(dir string) (err error) {
	fp := path.Join(dir, "js", "impress.js")
	fd, err := os.Create(fp)
	if err != nil {
		return
	}
	defer fd.Close()
	_, err = fd.Write([]byte(impressJS))
	return
}
