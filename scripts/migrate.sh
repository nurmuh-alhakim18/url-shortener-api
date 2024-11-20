#!/bin/bash

if [ -f .env ]; then
  source .env
fi

goose -dir sql/schemas turso $DB_URL up