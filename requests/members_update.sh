#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`
BACKEND_URL=`cat backend.url`

BODY=$( jq -n \
--arg teamname "$teamname" \
--argjson id "$id" \
--arg membername "$membername" \
--arg email "$email" \
--arg phone "$phone" \
--argjson isactive "$isactive" \
--argjson isadmin "$isadmin" \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X PUT -d "$BODY" "$BACKEND_URL/api/members/"
