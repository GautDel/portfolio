#!/bin/bash

# Run all commands in parallel
npx @tailwindcss/cli -i ./static/css/main.css -o ./static/css/output.css --watch --verbose
