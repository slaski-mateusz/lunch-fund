#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`
BACKEND_URL=`cat backend.url`

BODY=$( jq -n \
--arg teamname "$teamname" \
--arg ordername "$ordername" \
--argjson timestamp "$timestamp" \
--argjson deliverycost "$deliverycost" \
--argjson tipcost "$tipcost" \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X PUT -d "$BODY" $BACKEND_URL/api/orders/
