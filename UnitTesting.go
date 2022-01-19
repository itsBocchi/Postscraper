package main

import "github.com/itsbocchi/postScraper/handlers"

func main() {
	//scraper.FetchURL("https://pbs.twimg.com/media/FJNKWr3aUAQTT37?format=jpg&name=4096x4096")
	urls := []string{
		"https://pbs.twimg.com/media/FJNKWr3aUAQTT37?format=jpg&name=4096x4096",
		"https://pbs.twimg.com/media/FJTLxsvaMAA0V1x?format=jpg&name=medium",
		"https://pbs.twimg.com/media/FJUSdBfVkAAbgDL?format=jpg&name=large",
		"https://pbs.twimg.com/media/FJTsLPmaAAEdBSW?format=jpg&name=small",
		"https://pbs.twimg.com/media/FJRwAJ5aUAMmAe-?format=jpg&name=4096x4096",
		"https://pbs.twimg.com/media/FJTGBvKaUAE3T5g?format=jpg&name=large",
		"https://pbs.twimg.com/media/FJSEG14VgAQoNwE?format=jpg&name=medium",
	}

	fileExt := handlers.GetFileExtension(urls[2])

	handlers.FetchAndSave(urls[2], fileExt)
}
