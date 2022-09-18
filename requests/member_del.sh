#!/bin/bash
ID=$1
if [ -z "$ID" ]
then
	exit
fi

BODY_TEMPLATE=`cat member_del.json`

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{ID\}\}/$ID/g `
echo $BODY

curl -v -X DELETE "$BODY" 127.0.0.1:8080/members/
