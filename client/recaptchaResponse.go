package client

type RecaptchaResponse struct {
	success      bool
	challenge_ts string
	hostname     string
}
