#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`

BODY=$( jq -n \
--arg teamname "$teamname" \
--arg membername "$membername" \
--arg email "$email" \
--arg phone "$phone" \
--argjson isactive "$isactive" \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X PUT -d "$BODY" 127.0.0.1:8080/members/
