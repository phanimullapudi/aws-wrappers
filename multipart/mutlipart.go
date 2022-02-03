package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func checkfiles(foldername string) {

	files, err := ioutil.ReadDir("/Users/phanimullapudi/Documents/test/" + foldername)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() {
			fmt.Println("Need code for recurssive")
		} else {
			content, err := ioutil.ReadFile("/Users/phanimullapudi/Documents/test/" + foldername + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Sub -- File contents: %s", content)
		}
	}

}

func main() {

	files, err := ioutil.ReadDir("/Users/phanimullapudi/Documents/test")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			go checkfiles(file.Name())
		} else {
			content, err := ioutil.ReadFile("/Users/phanimullapudi/Documents/test/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("File contents: %s", content)

		}

	}

	var input string
	fmt.Scanln(&input)
}
