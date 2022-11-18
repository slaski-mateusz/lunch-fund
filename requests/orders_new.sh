#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`

BODY=$( jq -n \
--arg teamname "$teamname" \
--arg ordername "$ordername" \
--argjson timestamp "$timestamp" \
--argjson deliverycost "$deliverycost" \
--argjson tipcost "$tipcost" \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X PUT -d "$BODY" 127.0.0.1:8080/api/orders/
