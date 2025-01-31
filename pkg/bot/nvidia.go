package bot

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/http2"
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
	req.Header.Set("Referer", "https://www.nvidia.com/de-de/geforce/graphics-cards/50-series/rtx-5090/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	jar, err := cookiejar.New(nil)
	if err != nil {
		return response, fmt.Errorf("error creating cookie jar", err)
	}

	// create http2 client
	client := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			TLSClientConfig: &tls.Config{
				// Enable TLS renegotiation
				Renegotiation: tls.RenegotiateOnceAsClient,
				// You might need these depending on the server
				InsecureSkipVerify: true, // Remove in production!
				MaxVersion:         tls.VersionTLS13,
				MinVersion:         tls.VersionTLS12,
			},
		},
		Jar: jar,
	}

	resp, err := client.Do(req)
	if err != nil {
		return response, fmt.Errorf("error fetching API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return response, fmt.Errorf("error fetching API: %d %s", resp.StatusCode, resp.Status)
	}

	var responseBuffer []byte
	teeReader := io.TeeReader(resp.Body, bytes.NewBuffer(responseBuffer))

	body, err := io.ReadAll(teeReader)
	if err == nil {
		fmt.Println("Response Body:", string(body))
	}

	err = json.NewDecoder(bytes.NewReader(responseBuffer)).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("error parsing API: %w", err)
	}

	// Ensure list is empty
	if response.ListMap == nil {
		response.ListMap = make([]any, 0)
	}

	return response, nil
}
