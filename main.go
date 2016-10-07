/**
* Author: CZ cz.devnet@gmail.com
 */

package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n \t%s dest \n", os.Args[0], os.Args[0])
	flag.PrintDefaults()
	os.Exit(-1)
}

func initFlag() {
	flag.Usage = usage
}

func init() {
	initFlag()
}

func main() {
	fmt.Println("igen- Impress.js project generator")
	flag.Parse()

	if len(flag.Args()) != 1 {
		usage()
	}
	dirName := flag.Arg(0)
	err := mkdir(flag.Arg(0))
	if err != nil {
		println("err is ", err.Error())
		usage()
	}

	err = createIdxHTML(dirName)
	if err != nil {
		println("err is ", err.Error())
		usage()
	}

	err = createIdxCSS(dirName)
	if err != nil {
		println("err is ", err.Error())
		usage()
	}

	err = createImpressJS(dirName)
	if err != nil {
		println("err is ", err.Error())
		usage()
	}

}
