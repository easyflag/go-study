package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

var totalFile int

func listDir(path string, deepth uint8) error {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if deepth == 0 {
		fmt.Printf("%s\\%s\\\n", filepath.Dir(path), filepath.Base(path))
		fmt.Printf("@")
	} else {
		fmt.Printf("|")
	}
	for i := deepth; i > 0; i-- {
		fmt.Printf("	|")
	}
	fmt.Printf("-------%s\n", filepath.Base(path))
	deepth++

	separator := string(os.PathSeparator)
	for _, v := range dir {
		name := v.Name()
		if v.IsDir() {
			lowerPath := path + separator + name
			listDir(lowerPath, deepth)
		} else {
			for i := deepth; i > 0; i-- {
				fmt.Printf("|	")
			}
			fmt.Printf("|-------%s\n", name)

			totalFile++
		}
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "tree"
	app.Usage = "spread out all file under current directory"
	app.Version = "V1.0"

	app.Action = func(c *cli.Context) error {
		path := "."
		if c.NArg() > 0 {
			path = c.Args().First()
		}

		err := listDir(path, 0)
		fmt.Printf("there are %d file in %s", totalFile, path)
		return err
	}

	app.Run(os.Args)
}
