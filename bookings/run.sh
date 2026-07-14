#!/bin/bash
go build -o bookings cmd/web/*.go && ./bookings -dbname=bookings -dbuser=bookings_ope -cache=false -production=false