#!/bin/sh

while getopts "h:u:" opt; do
    case $opt in
    h) SERVER=$OPTARG ;;
    u) USER=$OPTARG ;;
    esac
done

docker buildx build --platform linux/arm64 -t line-bot:pi --load .
docker save line-bot:pi -o linebot_arm64.tar
scp linebot_arm64.tar ${USER}@${SERVER}:~/image
rm linebot_arm64.tar
