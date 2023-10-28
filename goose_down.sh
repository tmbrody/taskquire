#!/bin/bash

source .env

dbUser=$(echo $DATABASE_URL | cut -d':' -f1)
dbPass=$(echo $DATABASE_URL | cut -d':' -f2 | cut -d'@' -f1)
dbHost=$(echo $DATABASE_URL | cut -d'@' -f2 | cut -d':' -f1)
dbPort=$(echo $DATABASE_URL | cut -d':' -f3 | cut -d'/' -f1)
dbName=$(echo $DATABASE_URL | cut -d'/' -f2)

goose -dir sql/schema mysql "$dbUser:$dbPass@tcp($dbHost:$dbPort)/$dbName" down
goose -dir sql/schema mysql "$dbUser:$dbPass@tcp($dbHost:$dbPort)/$dbName" down