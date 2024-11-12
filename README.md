# Learning Go Lang

Walkthrough of https://learninggolang.com

See:
* https://github.com/renato0307/learning-go-api
* https://github.com/renato0307/learning-go-lib
* https://github.com/renato0307/learning-go-cli
* https://github.com/renato0307/learning-go-api-iac

## Uses RFC9068 JSON Web Token (JWT) Profile for OAuth 2.0 Access Tokens

* see: https://datatracker.ietf.org/doc/html/rfc9068
* see: https://datatracker.ietf.org/doc/html/rfc8414
* inspiration: https://github.com/coreos/go-oidc
* inspiration: https://github.com/auth0/go-jwt-middleware
* use: https://github.com/golang-jwt/jwt
* use: https://github.com/MicahParks/keyfunc/v3


## Keycloak URLs

Sign into Realm with KeyCloack UI:

* url: http://localhost:8081/realms/myrealm/account
* username: myuser
* password: mypassword

Sign into Admin Console:

* url: http://localhost:8081/admin
* username: admin
* password: admin

Get OAuth Client with Keycloak Client:

* url:  https://www.keycloak.org/app/#url=http://localhost:8081&realm=myrealm&client=myclient
* Keycloak URL:  http://localhost:8081
* Keycloak Realm: myrealm
* Keycloak Client: myclient
* username: myuser
* password: mypassword

OAuth Parameters:

* issuer: http://localhost:8081/realms/myrealm
* well known endpoint: http://localhost:8081/realms/myrealm/.well-known/oauth-authorization-server
* JWKS URl: http://localhost:8081/realms/myrealm/protocol/openid-connect/certs
* Auth Endpoint: http://localhost:8081/realms/myrealm/protocol/openid-connect/auth
* Token Endpoint: http://localhost:8081/realms/myrealm/protocol/openid-connect/token

