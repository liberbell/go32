#!/bin/bash
go build -o bookings cmd/web/*.go && ./bookings -dbuser=bookings_ope -cache=false -production=false