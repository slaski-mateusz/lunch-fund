#!/bin/bash
NAME=$1
EMAIL=$2
PHONE=$3
if [ -z "$NAME" ]
then
	exit
fi
if [ -z "$EMAIL" ]
then
	exit
fi
if [ -z "$PHONE" ]
then
	exit
fi

BODY_TEMPLATE=`cat member_new.json`

BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{NAME\}\}/$NAME/g | sed s/\{\{EMAIL\}\}/$EMAIL/g | sed s/\{\{PHONE\}\}/$PHONE/g`
echo $BODY

curl -v -X PUT "$BODY" 127.0.0.1:8080/members/
