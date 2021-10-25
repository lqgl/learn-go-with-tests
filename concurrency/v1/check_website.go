package concurrency

func CheckWebSite(url string) bool {
	return url == "https://www.cctv.com"
}

func CheckWebsites(urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = CheckWebSite(url)
	}
	return results
}
