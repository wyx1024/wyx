package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/text/transform"

	"golang.org/x/net/html/charset"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

var ratertime = time.Tick(time.Microsecond * 1500)

func Fetch(url string) ([]byte, error) {
	<-ratertime
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("http status %s", resp.StatusCode)
		return nil, err
	}
	body := bufio.NewReader(resp.Body)
	e := DetermineEncoding(body)
	bodyIO := transform.NewReader(body, e.NewDecoder())
	ioResult, err := ioutil.ReadAll(bodyIO)
	return ioResult, err
}
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
