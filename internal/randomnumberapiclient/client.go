package randomnumberapiclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const apiURL = "https://www.randomnumberapi.com/api/v1.0/"

func GetRandomNumber() (int, error) {
	url := apiURL + "random"
	var client http.Client
	client.Get(url)
	resp, err := client.Get(url)
	if err != nil {

		return 0, errors.New("error connecting to url to retrieve random number")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, errors.New("error reading body to retrieve random number")
		}
		var random []int
		err = json.Unmarshal(bodyBytes, &random)
		if err != nil || len(random) < 1 {
			return 0, errors.New("error linearizing json")
		}
		return random[0], nil
	}
	return 0, errors.New("error getting random number, received bad status")
}
