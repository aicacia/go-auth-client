# Aicacia Auth Go Client

```go
authClient := openapi.NewAPIClient(openapi.NewConfiguration())
SetAuthorization(authClient, "authorization-token")
currentUser, _httpResponse, _err := authClient.CurrentUserAPI.GetCurrentUser(context.Background()).Execute()
currentUserJson, _err := json.Marshal(currentUser)
log.Printf("currentUser: %s", currentUserJson)
```
