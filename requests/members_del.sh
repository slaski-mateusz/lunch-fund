#!/bin/bash
TEAM_NAME=$1
MEMBER_ID=$2

if [ -z "$TEAM_NAME" ]
then
	exit
fi

if [ -z "$MEMBER_ID" ]
then
	exit
fi

BODY_TEMPLATE=`cat members_del.json`

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{TEAM_NAME\}\}/"$TEAM_NAME"/g | sed s/\{\{MEMBER_ID\}\}/"$MEMBER_ID"/g`
echo $BODY

curl -v -X DELETE -d "$BODY" "127.0.0.1:8080/members/"
