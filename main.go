package main
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"

	"log"
	"os"

	"bufio"

)
//http://www.omdbapi.com/ could be used but no info about next episode..


func main() {
	readFile("seriekoll.txt")
	fmt.Println("--------------")
	fmt.Println("https://next-episode.net/\n")
	fmt.Println("Goto wanted series and paste in url\n")
	fmt.Println("--------------")
	url :=""
	fmt.Fscanf(os.Stdin,"%s",&url)
	writeFile("seriekoll.txt",url+"\n\r")
}

func Scrape(url string) {
	if(url != ""){
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
			fmt.Print("errrro")
		}


		pgtitle := doc.Find("div#next_episode").Contents().Text() //gets too much information right now

		fmt.Printf(pgtitle)
	}


}


func readFile(filename string) {

	file, err := os.Open(filename)
	if err == nil {

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() { // internally, it advances token based on sperator
			//fmt.Println(scanner.Text())  // token in unicode-char
			if (scanner.Text() != "") {
				fmt.Println("Info from: " + scanner.Text())
				Scrape(scanner.Text())
			}

		}
	}
	fmt.Println("")
}

func writeFile(filename string, msg string) {
	fo, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	w := bufio.NewWriter(fo)
	fmt.Fprint(w, msg)
	w.Flush()
}
