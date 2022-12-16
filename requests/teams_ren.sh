#!/bin/bash
OLDNAME=$1
NEWNAME=$2
if [ -z "$OLDNAME" ]
then
	exit
fi
if [ -z "$NEWNAME" ]
then
	exit
fi


BODY_TEMPLATE=`cat teams_ren.json`
BACKEND_URL=`cat backend.url`

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{OLDNAME\}\}/$OLDNAME/g | sed s/\{\{NEWNAME\}\}/$NEWNAME/g `
echo $BODY

curl -v -X POST -d "$BODY" "$BACKEND_URL/teams/"