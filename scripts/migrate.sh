#!/bin/bash

if [ -f .env ]; then
  source .env
fi

goose -dir sql/schemas turso -allow-missing $DB_URL up