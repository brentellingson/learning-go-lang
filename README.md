# Learning Go Lang

Walkthrough of https://learninggolang.com

See:
* https://github.com/renato0307/learning-go-api
* https://github.com/renato0307/learning-go-lib
* https://github.com/renato0307/learning-go-cli
* https://github.com/renato0307/learning-go-api-iac

## Uses RFC9068 JSON Web Token (JWT) Profile for OAuth 2.0 Access Tokens

* see: https://datatracker.ietf.org/doc/html/rfc9068
* inspiration: https://github.com/coreos/go-oidc
* inspiration: https://github.com/auth0/go-jwt-middleware

* use: https://github.com/golang-jwt/jwt
* use: 

## Keycloak URLs

Sign into Account Console:

* url: http://localhost:8081/realms/myrealm/account
* username: myuser
* password: mypassword

Sign into Admin Console:

* url: http://localhost:8081/admin
* username: admin
* password: admin

Sign in with OAuth:

* url:  https://www.keycloak.org/app/#url=http://localhost:8081&realm=myrealm&client=myclient
* Keycloak URL:  http://localhost:8081
* Keycloak Realm: myrealm
* Keycloak Client: myclient
* username: myuser
* password: mypassword

