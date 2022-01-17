package main

import "github.com/itsbocchi/postScraper/scraper"

func main() {
	//scraper.FetchURL("https://pbs.twimg.com/media/FJNKWr3aUAQTT37?format=jpg&name=4096x4096")
	url := "https://pbs.twimg.com/media/FJNKWr3aUAQTT37?format=jpg&name=4096x4096"
	fileExt := scraper.GetFileExtension(url)

	scraper.FetchAndSave(url, fileExt)
}
