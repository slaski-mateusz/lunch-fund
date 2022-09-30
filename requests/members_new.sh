#!/bin/bash
TEAM_NAME=$1
MEMBER_NAME=$2
EMAIL=$3
PHONE=$4

if [ -z "$TEAM_NAME" ]
then
	exit
fi

if [ -z "$MEMBER_NAME" ]
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

BODY_TEMPLATE=`cat members_new.json`
echo $MEMBER_NAME
BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{TEAM_NAME\}\}/$TEAM_NAME/g | sed s/\{\{MEMBER_NAME\}\}/"$MEMBER_NAME"/g | sed s/\{\{EMAIL\}\}/$EMAIL/g | sed s/\{\{PHONE\}\}/$PHONE/g`
# BODY=`echo $BODY_TEMPLATE | tr '\n' ' ' | sed s/\{\{TEAM_NAME\}\}/$TEAM_NAME/g`


echo $BODY

curl -v -X PUT -d "$BODY" 127.0.0.1:8080/members/
