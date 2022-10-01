#!/bin/bash
TEAM_NAME=$1
if [ -z "$TEAM_NAME" ]
then
	exit
fi
echo "Listing team $TEAM_NAME members"
BODY_TEMPLATE=`cat teams_new.json`

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{TEAM_NAME\}\}/"$TEAM_NAME"/g `
echo $BODY

curl -v -X GET -d "$BODY" "127.0.0.1:8080/members/"