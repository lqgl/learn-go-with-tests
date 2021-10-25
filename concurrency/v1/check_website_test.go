package concurrency

import (
	"reflect"
	"testing"
)

func TestCheckWebsites(t *testing.T) {

	want := map[string]bool{
		"https://www.baidu.com":    false,
		"https://www.bilibili.com": false,
		"https://www.cctv.com":     true,
	}
	urls := []string{
		"https://www.baidu.com",
		"https://www.bilibili.com",
		"https://www.cctv.com",
	}
	got := CheckWebsites(urls)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, but want %v", got, want)
	}
}
