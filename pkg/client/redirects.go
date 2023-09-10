package client

import "fmt"

func (c *Client) ConfirmEmailURL(token string) string {
	return fmt.Sprintf("%s/check-email/confirm/%s", c.URL, token)
}

func (c *Client) ResetPasswordURL(token string) string {
	return fmt.Sprintf("%s/reset-password/%s", c.URL, token)
}
