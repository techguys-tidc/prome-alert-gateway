package requester

import (
	"alertmanager-gateway/pkg/models"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var client = &http.Client{}

func init() {
	client = &http.Client{}
	transport := &http.Transport{}

	proxy := os.Getenv("PROXY")
	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			fmt.Println("Error parsing proxy URL:", err)
			return
		}

		transport.Proxy = http.ProxyURL(proxyURL)
	}

	if skipTLS, _ := strconv.ParseBool(os.Getenv("SKIP_TLS_VERIFY")); skipTLS {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	client.Transport = transport
}

func PostForm(url, payload, token string) {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	q := req.URL.Query()
	q.Add("message", payload)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Line response: %s\n", resp.Status)
}

func PostMsTeams(url, payload string) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("MsTeams response: %s\n", resp.Status)
}

func PostUVdeskAccesstoken(url string, payload string) string {
	// Create a new HTTP POST request
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	// Convert response to string
	responseString := string(body)
	// fmt.Println("Response from Server:", responseString)

	// Return the response
	return responseString
}

func PostUVdeskCreateTicket(url, payload, token string) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// q := req.URL.Query()
	// q.Add("message", payload)
	// req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response Body:", string(body))
	respBody := &models.UVdeskRespBody{}
	if err := json.Unmarshal(body, respBody); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
		// return 0, fmt.Errorf("error unmarshalling response: %v", err)
	}

	fmt.Printf("Ticket ID: %d\n", respBody.TicketId)
	fmt.Printf("Incident Number: %d\n", respBody.Values.IncidentNumber)
	fmt.Printf("UVdesk response: %s\n", resp.Status)

}
