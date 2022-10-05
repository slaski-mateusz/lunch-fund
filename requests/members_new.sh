#!/bin/bash

source ./functions.sh

parseOptions "$@"

BODY_TEMPLATE=`loadBodyTemplate`

BODY=$( jq -n \
--arg TEAM_NAME "$TEAM_NAME" \
--arg MEMBER_NAME "$MEMBER_NAME" \
--arg EMAIL "$EMAIL" \
--arg PHONE "$PHONE" \
--argjson IS_ACTIVE "$IS_ACTIVE" \
"$BODY_TEMPLATE"\
)

echo $BODY

curl -v -X PUT -d "$BODY" 127.0.0.1:8080/members/
