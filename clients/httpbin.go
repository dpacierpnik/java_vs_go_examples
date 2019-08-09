package clients

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Httpbin struct {
	Url string
}

func (c *Httpbin) Json() (string, error) {

	resp, getErr := http.Get(c.Url)
	if getErr != nil {
		return "", errors.New("Error getting response from Httpbin service. Root cause: " + getErr.Error())
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("ERROR: Closing response body failed. Root cause: " + err.Error())
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Error calling Httpbin service. Response status code is: " + strconv.Itoa(resp.StatusCode))
	}

	respBytes, readRespErr := ioutil.ReadAll(resp.Body)
	if readRespErr != nil {
		return "", errors.New("Error reading response from Httpbin service. Root cause: " + readRespErr.Error())
	}

	return string(respBytes), nil
}
