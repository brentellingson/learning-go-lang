#!/usr/bin/env bash

set -eux

KEYCLOAK_URL=http://localhost:8081

# Get Keycloak Access Token
KEYCLOAK_TOKEN=$(curl --url $KEYCLOAK_URL/realms/master/protocol/openid-connect/token \
                    --data "grant_type=password&client_id=admin-cli&username=admin&password=admin" \
                    --silent \
                | jq -r '.access_token')

# Create a keycloak realm
curl --url $KEYCLOAK_URL/admin/realms \
     --header "Authorization: Bearer $KEYCLOAK_TOKEN" \
     --json '{ "realm": "myrealm", "enabled": true }'

# Create a keycloadk openid client
curl --url $KEYCLOAK_URL/admin/realms/myrealm/clients \
     --header "Authorization: Bearer $KEYCLOAK_TOKEN" \
     --json '{ 
            "clientId": "myclient", 
            "protocol": "openid-connect",
            "redirectUris": [ "https://www.keycloak.org/app/*" ],
            "webOrigins": [ "https://www.keycloak.org" ]
        }'

# Create a keycloak user
curl --url $KEYCLOAK_URL/admin/realms/myrealm/users \
     --header "Authorization: Bearer $KEYCLOAK_TOKEN" \
     --json '{
            "username": "myuser",
            "firstName": "My",
            "lastName": "User",
            "email": "myuser@example.com",
            "enabled": true,
            "credentials": [{ "type": "password", "value": "mypassword" }]
        }'
