package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/coaraujo/go-mongo-rabbitmq/domain"
)

const (
	secretKey = "6LcFjWQUAAAAALVs39sxT2Jo-1yyLldfF8EeBfaa"
	off       = "off"
	baseURL   = "https://www.google.com/recaptcha/api/siteverify?secret="
)

type RecaptchaClient struct {
	Vote       *domain.Vote
	httpClient *http.Client
	response   *RecaptchaResponse
}

func Verify(vote *domain.Vote, auth string) bool {
	fmt.Println("[RECAPTCHACLIENT] Verify method invoked with auth = ", auth)

	if auth == "off" {
		fmt.Println("[RECAPTCHACLIENT] Auth is off.")
		return true
	}

	client := NewRecatchaClient(nil)
	path := secretKey + "&response=" + vote.RecaptchaToken

	req, err := client.newRequest("POST", path, nil)
	if err != nil {
		fmt.Println("[RECAPTCHACLIENT] Error while calling POST.")
		return false
	}

	var response RecaptchaResponse
	_, err = client.do(req, &response)

	fmt.Println("[RECAPTCHACLIENT] Response.success = ", response.success)
	return response.success

}

func NewRecatchaClient(httpClient *http.Client) *RecaptchaClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &RecaptchaClient{httpClient: httpClient}
	return c
}

func (c *RecaptchaClient) newRequest(method, path string, body interface{}) (*http.Request, error) {
	u := baseURL + path

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (r *RecaptchaClient) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
