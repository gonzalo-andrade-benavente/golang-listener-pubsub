package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"net/url"
	"os"
	"strings"

	"example.com/hexagonal-listener/src/domain"
)

type HttpClient struct{}

func GetToken() (string, error) {

	data := url.Values{}

	data.Set("grant_type", os.Getenv("GRANT_TYPE"))
	data.Set("client_id", os.Getenv("CLIENT_ID"))
	data.Set("client_secret", os.Getenv("CLIENT_SECRET"))

	resp, err := http.Post(os.Getenv("URL_AUTHORIZATION"), "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))

	if err != nil {
		return "Error making token request", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "Error reading response body:", err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("Token request failed with status %d: %s", resp.StatusCode, body), err
	}

	var tokenResponse map[string]interface{}

	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "Error parsing JSON response:", err
	}

	accessToken, ok := tokenResponse["access_token"].(string)

	if !ok {
		fmt.Println("Error extracting access token from response:", tokenResponse)
	}

	return accessToken, nil
}

func GetTransaction(token string, transaction domain.Transaction) (string, error) {

	url := fmt.Sprintf("%s/%s", os.Getenv("URL_PIX"), transaction.Data.TransactionNumber)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "Error creating request", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("x-environment", os.Getenv("ENV"))
	req.Header.Add("X-country", transaction.Country)
	req.Header.Add("X-commerce", transaction.Commerce)
	req.Header.Add("X-chRef", transaction.Channel)
	req.Header.Add("X-cmRef", transaction.Cmref)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return "Error making request:", err

	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "Error reading response body:", err
	}

	return string(body), nil

}
