#!/bin/bash

echo "-----------------Script-02-----------------"

echo "########### Cloudformation Start ###########"
aws cloudformation deploy \
 --stack-name notes-api-infra \
 --template-file "/cloudformation/cloudformation-test.yml" \
 --endpoint-url http://localstack:4566

