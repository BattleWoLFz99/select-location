#!/bin/bash

up() {
    docker compose up -d --build
}


down() {
   docker compose down
}

if [ $# -eq 0 ]; then
    echo "No arguments provided. Please use 'start' or 'down'."
    exit 1
fi

case $1 in
    up)
        up
        ;;
    down)
        down
        ;;
    *)
        echo "Invalid argument: $1"
        echo "Usage: $0 {start|down}"
        exit 1
        ;;
esac