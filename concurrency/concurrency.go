package concurrency

type UrlChecker func(string) bool
type result struct {
	string
	bool
}

func CheckUrl(uc UrlChecker, urls []string) (results map[string]bool) {
	results = make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		// go func(url string) { results[url] = uc(url) }(url)
		go func(url string) { resultChannel <- result{url, uc(url)} }(url)

	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	
	return
}
