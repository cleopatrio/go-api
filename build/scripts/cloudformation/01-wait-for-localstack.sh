#!/bin/bash

echo "-----------------Script-01-----------------"

echo "########### Check if localstack is up ###########"
until curl http://localstack:4566/health --silent; do
  echo "Localstack not up yet"
  sleep 1
done

