#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`

BODY=$( jq -n \
--arg TEAM_NAME "$TEAM_NAME" \
--argjson ID $ID \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X DELETE -d "$BODY" "127.0.0.1:8080/members/"
