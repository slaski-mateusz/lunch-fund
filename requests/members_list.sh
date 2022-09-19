#!/bin/bash
NAME=$1
if [ -z "$NAME" ]
then
	exit
fi
echo "Listing team $NAME members"
BODY_TEMPLATE=`cat teams_new.json`

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{NAME\}\}/$NAME/g `
echo $BODY

curl -v -X GET -d "$BODY" "127.0.0.1:8080/members/"