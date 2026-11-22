package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const cloudflareUrl = "https://www.cloudflare.com/cdn-cgi/trace"

func Country() (string, error) {
	client := &http.Client{Timeout: 2 * time.Second}

	resp, err := client.Get(cloudflareUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("geo lookup failed: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(body), "\n") {
		if code, ok := strings.CutPrefix(line, "loc="); ok {
			return code, nil
		}
	}

	return "", fmt.Errorf("country not found in geo response")
}
