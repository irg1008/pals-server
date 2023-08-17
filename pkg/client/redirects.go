package client

import "fmt"

func (c *Client) ConfirmEmailURL(token string) string {
	return fmt.Sprintf("%s/confirm-email?id=%s", c.URL, token)
}

func (c *Client) ResetPasswordURL(token string) string {
	return fmt.Sprintf("%s/reset-password?id=%s", c.URL, token)
}
