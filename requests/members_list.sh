#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`
BACKEND_URL=`cat backend.url`

BODY=$( jq -n \
--arg teamname "$teamname" \
"$BODY_TEMPLATE"\
)

# BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{TEAM_NAME\}\}/"$TEAM_NAME"/g `
echo $BODY

curl -v -X GET -d "$BODY" "$BACKEND_URL/api/members/"
# curl -v -X GET -d "$BODY" "BACKEND_URL/api/members/" | jq .