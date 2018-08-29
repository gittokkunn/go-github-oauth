# Githubのoauth認証

## About
### RedirectAuthrizeClient(c *gin.Context, clientID string)
### GetAccessTokenClient(c *gin.Context, clientID string, clientSecret string) (*CredentialInfo)
### CredentialInfo
##### `.AccessToken string`
- 認証に必要なアクセストークン
##### `.Scope string`
- APIアクセスのスコープ
##### `.TokenType`
- トークンの種類