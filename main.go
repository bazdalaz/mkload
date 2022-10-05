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

	"github.com/joho/godotenv"
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


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	main_path := os.Getenv("BASE_PATH")


	plant := flag.String("p", "", "the plant")

	net := flag.String("n", "A", "LCN [a/b]")
	flag.Parse()



	sstat_path := main_path + strings.ToUpper(*net) + "/"



	out_json, _ := create_json(plant, sstat_path)
	fmt.Printf("%s\n", out_json)

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


func create_json(plant *string, sstat_path string) ([]byte, error) {
	dir, err := os.Open(sstat_path)
	if err != nil {
		log.Fatal(err)
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	var out_json = []byte{}


	for _, f := range files {
		if strings.Contains(f.Name(), "SSTAT") {
			readFile, err := os.Open(sstat_path + f.Name())

			if err != nil {
				fmt.Println(err)
			}
			fileScanner := bufio.NewScanner(readFile)

			fileScanner.Split(bufio.ScanLines)

			prefix, err := make_prefix(*plant)
			if err != nil {		panic(err)

			}
			out_slice := []ContextData{}

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

			out_json, err = json.MarshalIndent(out_slice,"", " ")
			if err!= nil {
                fmt.Println(err)
				return nil, err
			}

			if string(out_json) == "[]" {
				log.Printf("no data available for plant " + *plant)
				return nil, err
			}

			err = os.WriteFile(*plant+"_seqs.json", out_json, 0644)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

		}
	}
	return out_json, nil
}









