#!/bin/bash

echo "-----------------Script-01----------------- [localstack]"

echo "########### Creating profile ###########"

aws configure set aws_access_key_id default_access_key
aws configure set aws_secret_access_key default_secret_key
aws configure set region us-east-2

echo "########### Listing profile ###########"
aws configure list
