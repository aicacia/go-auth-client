package auth_client

import (
	openapi "github.com/aicacia/go-auth-client/openapi"
)

func SetAuthorization(client *openapi.APIClient, token string) *openapi.APIClient {
	configuration := client.GetConfig()
	configuration.AddDefaultHeader("Authorization", "Bearer "+token)
	return client
}

func SetTenentID(client *openapi.APIClient, tenentId string) *openapi.APIClient {
	configuration := client.GetConfig()
	configuration.AddDefaultHeader("Tenent-ID", tenentId)
	return client
}
