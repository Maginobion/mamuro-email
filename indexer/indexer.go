package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var JSONname = "data.json"

func getJSON(path string) []byte {

	//Extract data from absolute path

	dataArr := strings.Split(path, "/")

	//Read the file content

	content, err := os.ReadFile(path)

	check(err)

	//Find first three email matches

	re := regexp.MustCompile("[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*")
	match := re.FindAllString(string(content), 3)

	//Find subject

	re = regexp.MustCompile("Subject: ([A-Za-z0-9 -_:#.]*)")
	subMatch := re.FindString(string(content))

	//If length less than 3, set sender and receiver as Unknown, subject starts as Unknown too

	sender := "Unknown"
	receiver := "Unknown"
	subject := "Unknown"

	if len(match) > 2 {
		sender = match[1]
		receiver = match[2]
	}

	//If subject is found, extract "Subject: " part and set to variable

	if subMatch != "Subject: " {
		subject = subMatch[9:len(subMatch)]
	}

	//Build struct for parsing

	type JSONvalues struct {
		User     string
		Sender   string
		Receiver string
		Subject  string
		Category string
		Content  string
	}

	group := JSONvalues{
		User:     dataArr[2],
		Sender:   sender,
		Receiver: receiver,
		Subject:  subject,
		Category: dataArr[3],
		Content:  string(content),
	}

	//Parse the struct as a JSON

	bytes, _ := json.Marshal(group)

	return bytes
}

func searchInside(path string) {

	//Tracking path on console

	fmt.Println(path)

	//Read all items inside a directory and check if they are directories or files

	items, _ := os.ReadDir(path)

	for _, item := range items {
		if item.IsDir() {
			//If they are directories, keep searching inside

			searchInside(path + "/" + item.Name())
		} else if filepath.Ext(item.Name()) != ".txt" {

			//If file, open JSON data file for writing

			f, err := os.OpenFile(JSONname, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

			check(err)

			//Write the JSON-parsed content

			_, err = f.Write(append(getJSON(path+"/"+item.Name()), []byte(",")...))

			check(err)

			//Write a comma at the end and close the file

			f.Close()
		}
	}
}

func main() {

	//Check that an argument is being passed

	if len(os.Args) < 2 {
		fmt.Print("Ingrese el directorio a indexar." + "\n")
		return
	}

	//Get argument

	fileArg := os.Args[1]

	//Remove previous JSON files if existing

	os.Remove(JSONname)

	//Create JSON file

	f, err := os.Create(JSONname)

	check(err)

	//Push initial JSON structure

	startingjson := "{\"index\":\"mails\",\"records\":["

	f.WriteString(startingjson)

	f.Close()

	//Search inside the initial directory

	searchInside(fileArg)

	//After all files pushed into the JSON, close the JSON structure

	f, err = os.OpenFile(JSONname, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	check(err)

	//Write a comma at the end and close the file

	_, err = f.WriteString("{}]}")

	f.Close()

	//Post the JSON into the localhost

	localUrl := "http://localhost:4080/api/_bulkv2"
	includeFlag := "-i"
	authFlag := "-u"
	credentials := "admin:Complexpass#123"
	mode := "--data-binary"
	file := "@" + JSONname

	fmt.Println(file)

	cmd := exec.Command("curl", localUrl, includeFlag, authFlag, credentials, mode, file)

	fmt.Println("Posting, please wait...")

	stdout, err := cmd.Output()

	check(err)

	//Print the standard output
	fmt.Println(string(stdout))
}
