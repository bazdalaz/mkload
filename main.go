package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type ContextData struct {
	Point string `json:"point"`
	Sequence string `json:"sequence"`
	State string `json:"state"`
	CompTime string `json:"comp_time"`
	Phase string `json:"phase"`
	Step string `json:"step"`
	Statement string `json:"statement"`

}

const main_path = "/mnt/data/ASPEN_WORKS/LCNDB/MAIN/"

func main() {
	plant := flag.String("p", "", "the plant")

	net := flag.String("n", "A", "LCN [a/b]")
	flag.Parse()

	prefix, err := make_prefix(*plant)
	if err != nil {
		panic(err)

	}

	sstat_path := main_path + *net + "/"

	dir, err := os.Open(sstat_path)
	if err != nil {
		log.Fatal(err)
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), "SSTAT") {
			readFile, err := os.Open(sstat_path + f.Name())

			if err != nil {
				fmt.Println(err)
			}
			fileScanner := bufio.NewScanner(readFile)

			fileScanner.Split(bufio.ScanLines)

			for fileScanner.Scan() {
				if strings.HasPrefix(fileScanner.Text(), prefix) {
					var c = new(ContextData)
					c.Point = strings.TrimSpace(fileScanner.Text()[1:9])
					fmt.Println(c.Point)
				}
			}

			readFile.Close()

		}
	}
}

func make_prefix(plant string) (string, error) {
	if plant == "" {
		return "", errors.New("missing plant")
	}

	if plant == "651" {
		return " 51", nil
	}

	if plant == "634" {
		return " 41", nil
	}

	return string(' ') + plant[:2], nil
}
