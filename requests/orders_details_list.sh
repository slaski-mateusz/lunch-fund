#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`
BACKEND_URL=`cat backend.url`

BODY=$( jq -n \
--arg teamname "$teamname" \
--argjson id "$id" \
"$BODY_TEMPLATE"\
)

# BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{TEAM_NAME\}\}/"$TEAM_NAME"/g `
echo "Request body:"
echo $BODY

# curl -v -X GET -d "$BODY" "127.0.0.1:8080/api/orders_details/"
curl -v -X GET -d "$BODY" "$BACKEND_URL/api/orders_details/" | jq .