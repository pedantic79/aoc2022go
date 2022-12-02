#!/bin/bash

YEAR=2022
VANITY="github.com/pedantic79/aoc${YEAR}go"
NUM=$(printf "%02d" "$1")
CREDS="$HOME/Library/Preferences/com.github.gobanos.cargo-aoc/credentials.toml"

if [ -e "$CREDS" ]; then
    COOKIE=$(yj -t < "$CREDS" | jq '.session' -r)
else
    COOKIE=""
fi

if [ -d "day$NUM" ]; then
    echo "day$NUM already exists"
    exit 1
fi

cp -r template "day$NUM"
mv "day$NUM/dayNN.go" "day$NUM/day$NUM.go"
mv "day$NUM/dayNN_test.go" "day$NUM/day${NUM}_test.go"
mv "day$NUM/dayNN.txt" "day$NUM/day$1.txt"

gsed -i "s/package dayNN/package day$NUM/" "day$NUM/day$NUM.go"
gsed -i "s/const day uint = 0/const day uint = $1/" "day$NUM/day$NUM.go"

gsed -i "s/package dayNN/package day$NUM/" "day$NUM/day${NUM}_test.go"

if [ -n "$COOKIE" ]; then
    curl -s --cookie "session=$COOKIE" "https://adventofcode.com/$YEAR/day/$1/input" > "day$NUM/day$1.txt"
fi

entry=$(printf "_ \"%s/day%s\"" "$VANITY" "$NUM")
gsed -i "/^)$/i $entry" main.go
gofmt -w main.go
