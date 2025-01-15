#/bin/bash

# Build and push the image to my docker repo
# You can set the docker login method as per your wish, I am going with ENV here.

sudo docker build -t nitinjilla/welcome:latest .
sudo docker login -u $DOCKER_USER -p $DOCKER_PASSWD docker.io
sudo docker tag nitinjilla/welcome welcome:latest
sudo docker push nitinjilla/welcome:latest
