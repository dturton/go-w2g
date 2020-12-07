package main

const (
	merchantId = "M343"
	token      = "your-api-key-here"
	secret     = "your-secret-pass-here"
	apiVersion = "2"
)

var client = client.New(merchantId, token, secret, apiVersion)
