package ip

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPAddressData struct {
	IPv4         string `json:"ip"`
	Hostname     string `json:"hostname"`
	City         string `json:"city"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	Location     string `json:"loc"`
	Organization string `json:"org"`
	Postal       string `json:"postal"`
	Timezone     string `json:"timezone"`
}

func GetIPInfo() (IPAddressData, error) {
	var ipData IPAddressData

	resp, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		return ipData, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ipData, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&ipData)
	if err != nil {
		return ipData, fmt.Errorf("failed to decode JSON: %v", err)
	}

	return ipData, nil
}

func GetIPv4() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(ip), nil
}
