package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
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

	// Set headers to avoid nvidia block... cuz... cuz... couscous
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en,pt-BR;q=0.9,pt;q=0.8,en-US;q=0.7,ja;q=0.6")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("DNT", "1")
	req.Header.Set("Origin", "https://marketplace.nvidia.com")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Referer", "https://marketplace.nvidia.com/")
	req.Header.Set("Sec-CH-UA", `"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"`)
	req.Header.Set("Sec-CH-UA-Mobile", "?0")
	req.Header.Set("Sec-CH-UA-Platform", `"Linux"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36")

	jar, err := cookiejar.New(nil)
	if err != nil {
		return response, fmt.Errorf("error creating cookie jar: %w", err)
	}

	// create http2 client
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout
		Transport: &http.Transport{
			ForceAttemptHTTP2: true, // Ensure HTTP/2 is used
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

	//var responseBuffer []byte
	//teeReader := io.TeeReader(resp.Body, bytes.NewBuffer(responseBuffer))
	//
	//body, err := io.ReadAll(teeReader)
	//if err == nil {
	//	fmt.Println("Response Body:", string(body))
	//}
	//
	//err = json.NewDecoder(bytes.NewReader(responseBuffer)).Decode(&response)
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("error parsing API: %w", err)
	}

	// Ensure list is empty
	if response.ListMap == nil {
		response.ListMap = make([]Product, 0)
	}

	return response, nil
}
