#!/usr/bin/env bash

eval "$(ssh-agent -s)"
ssh-add "$HOME/.ssh/pkey"

TEMP="/tmp"
REPO="git@github.com:Berveglieri/challenge.git"
STRUCT="challenge/hive/backups"

cd $TEMP
git clone $REPO
cd $STRUCT

sudo pgtool -operation backup -h app-db.c1vfijgzuyxj.eu-central-1.rds.amazonaws.com -p 5432 -u master -P '' -d hive

git add . && git commit -m "new backup" && git push

cd $TEMP && rm -rf challenge