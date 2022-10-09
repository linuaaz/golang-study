package ch2

import (
	"net/http"
	"testing"
)

func TestParallel(t *testing.T) {
	var urls = map[string]string{
		"baidu": "https://www.baidu.com",
		"bing":  "https://cn.bing.com",
		//"google": "https://www.google.com", will time out
	}

	for k, v := range urls {
		v := v
		ok := t.Run(k, func(t *testing.T) {
			t.Parallel()
			resp, err := http.Get(v)
			if err != nil {
				t.Fatalf("failed to get %s: %v", v, err)
			}

			resp.Body.Close()
		})

		t.Logf("run: %t", ok)
	}
}
