package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebSiteChecker(url string) bool {
	return url == "https://www.cctv.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.baidu.com",
		"https://www.bilibili.com",
		"https://www.cctv.com",
	}

	actualResults := CheckWebsites(mockWebSiteChecker, websites)

	// want := len(websites)
	// got := len(actualResults)

	// if want != got {
	// 	t.Fatalf("Wanted %v, got %v", want, got)
	// }

	expectedResults := map[string]bool{
		"https://www.baidu.com":    false,
		"https://www.bilibili.com": false,
		"https://www.cctv.com":     true,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Errorf("Wanted %v, but got %v", expectedResults, actualResults)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebSites(b *testing.B) {
	websites := make([]string, 100)
	for i := 0; i < len(websites); i++ {
		websites[i] = "just a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, websites)
	}
}
