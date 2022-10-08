#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`

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

curl -v -X POST -d "$BODY" 127.0.0.1:8080/members/
