#!/bin/bash
NAME=$1
if [ -z "$NAME" ]
then
	exit
fi
echo "Adding team: $NAME"
BODY_TEMPLATE=`cat teams_new.json`

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{NAME\}\}/$NAME/g `
echo $BODY

curl -v -X PUT -d "$BODY" "127.0.0.1:8080/teams/"