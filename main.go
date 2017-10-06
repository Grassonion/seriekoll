package main
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"

	"log"
)
/*
https://jonathanmh.com/web-scraping-golang-goquery/




 */
func main() {
	fmt.Print("Yeahllo\n")

	Scrape()
}

func Scrape() {
	doc, err := goquery.NewDocument("http://www.imdb.com/title/tt2467372/episodes?ref_=tt_ov_epl")
	if err != nil {
		log.Fatal(err)
		fmt.Print("errrro")
	}


	pgtitle := doc.Find("h3").Contents().Text()
	fmt.Printf(pgtitle)


/*
	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
*/
}

