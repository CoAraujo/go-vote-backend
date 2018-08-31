package client

import (
	"fmt"

	client "github.com/coaraujo/go-vote-backend/client"
	domain "github.com/coaraujo/go-vote-backend/domain"
)

const (
	secretKey = "6LcFjWQUAAAAALVs39sxT2Jo-1yyLldfF8EeBfaa"
	off       = "off"
	baseURL   = "https://www.google.com/recaptcha/api/siteverify?secret="
)

type RecaptchaClient struct {
	Vote     *domain.Vote
	response *RecaptchaResponse
}

func Verify(vote *domain.Vote, auth string) bool {
	fmt.Println("[RECAPTCHACLIENT] Verify method invoked with auth = ", auth)

	if auth == "off" {
		fmt.Println("[RECAPTCHACLIENT] Auth is off.")
		return true
	}

	c := client.NewClient(nil)
	c.BaseURL = baseURL
	path := secretKey + "&response=" + vote.RecaptchaToken

	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		fmt.Println("[RECAPTCHACLIENT] Error while calling POST.")
		return false
	}

	var response RecaptchaResponse
	_, err = c.Do(req, &response)

	fmt.Println("[RECAPTCHACLIENT] Response.success = ", response.success)
	return response.success
}
