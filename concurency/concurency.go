package concurency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resutlChannel := make(chan result)
	for _, url := range urls {
		go func() {
			resutlChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		result := <-resutlChannel
		results[result.string] = result.bool
	}

	return results
}
