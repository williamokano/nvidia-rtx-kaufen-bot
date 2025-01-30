package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func CheckAPI() (APIResponse, error) {
	var response APIResponse

	baseURL, err := url.Parse(endpointURL)
	if err != nil {
		return response, fmt.Errorf("error parsing URL: %w", err)
	}

	// Define query parameters
	params := url.Values{}
	params.Add("status", "1")
	params.Add("skus", "NVGFT590")
	params.Add("locale", "de-de")

	// Set the RawQuery
	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return response, fmt.Errorf("error creating request: %w", err)
	}

	// Add some random header to appear more generic, coming from browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Referer", "https://www.google.com")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("error fetching API: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("error parsing API: %w", err)
	}

	// Ensure list is empty
	if response.ListMap == nil {
		response.ListMap = make([]any, 0)
	}

	return response, nil
}
