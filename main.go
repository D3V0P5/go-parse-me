package main

//for i in {1..3} ; do mkdir $i ; for z in {1..3} ; do cp ex $i/$z.log ; done ; done

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ua-parser/uap-go/uaparser"
)

var path = "log"
var parser, errr = uaparser.New("./regexes.yaml")

var fo, err = os.Create("output.csv")

func main() {

	//err :=
	filepath.Walk(path, walkDir)
	defer fo.Close()
}

func walkDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("error 「%v」 at a path 「%q」\n", err, path)
		return err
	}
	if info.IsDir() {
		return nil
	}

	realpath := filepath.Dir(path) + "/" + info.Name()
	fmt.Println(realpath)
	readFile(realpath)
	return nil
}

func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !strings.HasPrefix(scanner.Text(), "2") {
			continue
		}
		uaarr := strings.Split(scanner.Text(), " ")
		ua := strings.Replace(uaarr[9], "+", " ", -1)
		fmt.Println(ua)
		//bla := "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us; Silk/1.1.0-80) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16 Silk-Accelerated=true"

		client := parser.Parse(ua)
		row := uaarr[8] + "," + client.UserAgent.Family + "," + client.UserAgent.Major +
			"," + client.UserAgent.Minor + "\r\n"
		fo.WriteString(row)
		row = uaarr[8] + "," + client.Os.Family + "," + client.Os.Major + "," +
			client.Device.Family + "\r\n"
		fo.WriteString(row)
		/*
			fmt.Println(client.UserAgent.Family) // "Amazon Silk"
			fmt.Println(client.UserAgent.Major)  // "1"
			fmt.Println(client.UserAgent.Minor)  // "1"
			fmt.Println(client.UserAgent.Patch)  // "0-80"
			fmt.Println(client.Os.Family)        // "Android"
			fmt.Println(client.Os.Major)         // ""
			fmt.Println(client.Os.Minor)         // ""
			fmt.Println(client.Os.Patch)         // ""
			fmt.Println(client.Os.PatchMinor)    // ""
			fmt.Println(client.Device.Family)    // "Kindle Fire"
		*/

	}

}
