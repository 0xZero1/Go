package shodan

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type APIInfo struct {
	QueryCredits 	int 	`json:"query_credits"`
	ScanCredits 	int 	`json:"scan_credits"`
	Telnet			int		`json:"telnet"`
	Plan			string	`json:"plan"`
	https			bool	`json:"https"`
	Unlocked		bool	`json:"unlocked"`
}

func (s *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Client) HostSearch(q string) (*HostSearch, error) {
	res, err := http.Get(
		fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s",
		BaseURL, s.apiKey, q),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret HostSearch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}