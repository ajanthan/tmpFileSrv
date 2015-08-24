#!/bin/bash
if [ -d public ];
then
    echo "Deleting public folder"
    rm -rf public
fi
go build 
cd ui
grunt build:dist
cd ..
mkdir public
cp -r ui/dist/* public

