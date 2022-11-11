#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`

BODY=$( jq -n \
--arg teamname "$teamname" \
--argjson orderid "$orderid" \
--argjson memberid "$memberid" \
--argjson founder "$isfounder" \
--argjson amount "$amount" \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X PUT -d "$BODY" 127.0.0.1:8080/api/orders_details/
