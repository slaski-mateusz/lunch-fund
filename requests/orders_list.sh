#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`

BODY=$( jq -n \
--arg teamname "$teamname" \
"$BODY_TEMPLATE"\
)

# BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{TEAM_NAME\}\}/"$TEAM_NAME"/g `
echo $BODY

# curl -v -X GET -d "$BODY" "127.0.0.1:8080/api/orders/"
curl -v -X GET -d "$BODY" "127.0.0.1:8080/api/orders/" | jq .