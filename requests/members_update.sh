#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`

BODY=$( jq -n \
--arg TEAM_NAME "$TEAM_NAME" \
--argjson ID "$ID" \
--arg MEMBER_NAME "$MEMBER_NAME" \
--arg EMAIL "$EMAIL" \
--arg PHONE "$PHONE" \
--argjson IS_ACTIVE "$IS_ACTIVE" \
--argjson IS_ADMIN "$IS_ADMIN" \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X POST -d "$BODY" 127.0.0.1:8080/members/
