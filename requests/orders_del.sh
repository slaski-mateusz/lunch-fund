#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`

BODY=$( jq -n \
--arg teamname "$teamname" \
--argjson id $id \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X DELETE -d "$BODY" "127.0.0.1:8080/api/orders/"
