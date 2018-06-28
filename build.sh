#! /bin/bash

# Build web UI
cd ./web
pwd
go install
cp ../../../bin/web.exe ../../../bin/video_server_web_ui/web
pwd
cp -R /f/GoProjects/src/video_server/template /f/GoProjects/bin/video_server_web_ui/