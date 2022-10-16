#!/bin/bash
TEAM_NAME=$1
if [ -z "$TEAM_NAME" ]
then
	exit
fi

BODY_TEMPLATE=`cat teams_del.json`

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{TEAM_NAME\}\}/"$TEAM_NAME"/g `
echo $BODY

curl -v -X PUT -d "$BODY" "127.0.0.1:8080/api/teams/"
curl -v -X DELETE -d "$BODY" "127.0.0.1:8080/teams/"
