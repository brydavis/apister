package main

import "fmt"

func main() {

	{ // GOV JOBS API
		req := Request{}

		req["query"] = "jobs in king wa"
		req["size"] = 100
		req["tags"] = []string{"city", "county", "state"}

		fmt.Println(PrettyJSON(req.Get("https://api.usa.gov/jobs/search.json?")))
	}

	{ // NHL NEWS XML
		req := Request{}

		fmt.Println(req.Get("http://www.nhl.com/rss/news.xml"))
	}
}
