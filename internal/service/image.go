package service

import (
	"io/ioutil"
	"net/http"
)

func Download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return ioutil.ReadAll(resp.Body)
}

// PyModelResult 打开图片
func PyModelResult(data []byte) {
	return
}
