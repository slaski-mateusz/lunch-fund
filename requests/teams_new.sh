#!/bin/bash
TEAM_NAME=$1
if [ -z "$TEAM_NAME" ]
then
	exit
fi
echo "Adding team: $TEAM_NAME"
BODY_TEMPLATE=`cat teams_new.json`
BACKEND_URL=`cat backend.url`

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{TEAM_NAME\}\}/$TEAM_NAME/g `
echo $BODY

curl -v -X PUT -d "$BODY" "$BACKEND_URL/api/teams/"