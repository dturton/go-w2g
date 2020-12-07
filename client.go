package gow2g

import "fmt"

// Client makes all the API calls to w2g
type Client struct {
	merchantId string
	token      string
	secret     string
	apiVersion int
}

const (
	baseURL = "https://openapi.ware2go.io/v1/merchants/"
)

func New(merchantId, token, secret string, apiVersion int) Client {

	client := Client{merchantId: merchantId, token: token, secret: secret, apiVersion: apiVersion}
	fmt.Println("[New] Creating W2G client with: ", merchantId, apiVersion)

	return client
}
