#!/bin/bash
# API
cd golang
go run ./cmd &

# Servidor Webb
cd ../react/app-web
npm run dev
