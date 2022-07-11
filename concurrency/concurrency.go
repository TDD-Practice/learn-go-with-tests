package concurrency

type UrlChecker func(string) bool

func CheckUrl(uc UrlChecker, urls []string) (results map[string]bool) {
	results = make(map[string]bool)
	for _, url := range urls {
		results[url] = uc(url)
	}
	return
}
