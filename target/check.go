package target

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (t *Target) Check() (bool, error) {

	c := t.Client
	if c == nil {
		c = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	resp, err := c.Head(t.Url)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("%w: status - %d", ErrUnexpectedStatusCode, resp.StatusCode)
	}

	r := resp.Header.Get(AcceptRangesHeader)
	if r == "bytes" {
		t.AcceptRanges = true
	}

	l := resp.Header.Get(ContentLengthHeader)
	if l != "" {
		t.ContentLength, err = strconv.ParseInt(l, 10, 64)
		if err != nil {
			return false, err
		}
	}

	if t.ContentLength < 1 {
		return false, ErrIncorrectFile
	}

	return t.AcceptRanges, nil
}
