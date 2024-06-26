package requester

import (
	"crypto/tls"
	"fmt"
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
