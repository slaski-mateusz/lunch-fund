#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`
BACKEND_URL=`cat backend.url`

BODY=$( jq -n \
--arg teamname "$teamname" \
--argjson id $id \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X DELETE -d "$BODY" "$BACKEND_URL/api/members/"
