#!/bin/sh

echo "---------------------------------"
echo "---------------------------------"
echo ${EVANS_HOST}
echo ${EVANS_PORT}
echo ${EVANS_PROTO_FILE}
echo "---------------------------------"
echo "---------------------------------"
pwd
echo "---------------------------------"
echo "---------------------------------"
# Start evans
evans --host ${EVANS_HOST} ${EVANS_PROTO_FILE}

