package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type ContextData struct {
	Point     string `json:"point"`
	Sequence  string `json:"sequence"`
	State     string `json:"state"`
	CompTime  string `json:"comp_time"`
	Phase     string `json:"phase"`
	Step      string `json:"step"`
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

	sstat_path := main_path + strings.ToUpper(*net) + "/"

	dir, err := os.Open(sstat_path)
	if err != nil {
		log.Fatal(err)
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	out_slice := []ContextData{}
	out_json := []byte{}
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
					c.Sequence = strings.TrimSpace(fileScanner.Text()[39:49])
					c.State = strings.TrimSpace(fileScanner.Text()[49:55])
					c.CompTime = strings.TrimSpace(fileScanner.Text()[21:39])
					c.Phase = strings.TrimSpace(fileScanner.Text()[89:98])
					c.Step = strings.TrimSpace(fileScanner.Text()[99:109])
					c.Statement = strings.TrimSpace(fileScanner.Text()[109:])
					out_slice = append(out_slice, *c)
				}
			}
			readFile.Close()

			out_json, _ = json.MarshalIndent(out_slice,"", " ")


			fmt.Println(string(out_json))

			err = os.WriteFile(*plant+"_seqs.json", out_json, 0644)
			if err != nil {
				fmt.Println(err)
			}

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
