#!/bin/bash
NAME=$1
if [ -z "$NAME" ]
then
	exit
fi

BODY_TEMPLATE=`cat teams_del.json`

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{NAME\}\}/$NAME/g `
echo $BODY

curl -v -X DELETE -d "$BODY" "127.0.0.1:8080/teams/"
