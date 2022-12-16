#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`
BACKEND_URL=`cat backend.url`

BODY=$( jq -n \
--arg teamname "$teamname" \
--argjson orderid "$orderid" \
--argjson memberid "$memberid" \
--argjson founder "$isfounder" \
--argjson amount "$amount" \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X PUT -d "$BODY" $BACKEND_URL/api/orders_details/
