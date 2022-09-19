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

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{OLDNAME\}\}/$OLDNAME/g | sed s/\{\{NEWNAME\}\}/$NEWNAME/g `
echo $BODY

curl -v -X POST -d "$BODY" "127.0.0.1:8080/teams/"