#!/bin/bash

if [ "${SONAR_SCANNER_HOME}" != "" ]; then
    COMMAND="sonar-scanner -Dsonar.host.url=$SONAR_HOST_URL -Dsonar.login=$SONAR_TOKEN"
else
    COMMAND=""
fi

echo ${COMMAND}
${COMMAND}
