#!/usr/bin/env bash

set -eux

# KEYCLOAK_URL=http://localhost:8081

# wait 60 seconds for keycloak to start
timeout 60 bash -c 'while [[ "$(curl -s -o /dev/null -w ''%{http_code}'' http://localhost:${KC_HTTP_PORT}/admin/master/console/)" != "200" ]]; do sleep 1; done'

# Get Keycloak Access Token
TOKEN=$(curl --url http://localhost:${KC_HTTP_PORT}/realms/master/protocol/openid-connect/token \
             --data "grant_type=password&client_id=admin-cli&username=${KC_BOOTSTRAP_ADMIN_USERNAME}&password=${KC_BOOTSTRAP_ADMIN_PASSWORD}" \
             --silent \
          | jq -r '.access_token')

# delete the keycloak realm if it exists

curl --url http://localhost:${KC_HTTP_PORT}/admin/realms/myrealm \
     --header "Authorization: Bearer $TOKEN" \
     --request DELETE

# Create a keycloak realm
curl --url http://localhost:${KC_HTTP_PORT}/admin/realms \
     --header "Authorization: Bearer $TOKEN" \
     --json '{ "realm": "myrealm", "enabled": true }'

# Create a keycloadk openid client
curl --url http://localhost:${KC_HTTP_PORT}/admin/realms/myrealm/clients \
     --header "Authorization: Bearer $TOKEN" \
     --json '{ 
            "clientId": "myclient", 
            "protocol": "openid-connect",
            "redirectUris": [
               "https://www.keycloak.org/app/*",
               "http://localhost:8080/swagger/oauth2-redirect.html",
               "http://127.0.0.1:8082/oauth/callback"
           ],
            "webOrigins": [ "https://www.keycloak.org", "http://localhost:8080" ]
        }'

# Create a keycloak user
curl --url http://localhost:${KC_HTTP_PORT}/admin/realms/myrealm/users \
     --header "Authorization: Bearer $TOKEN" \
     --json '{
            "username": "myuser",
            "firstName": "My",
            "lastName": "User",
            "email": "myuser@example.com",
            "enabled": true,
            "credentials": [{ "type": "password", "value": "mypassword" }]
        }'
