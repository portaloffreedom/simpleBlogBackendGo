# This file creates a container that runs X11 and SSH services
# The ssh is used to forward X11 and provide you encrypted data
# communication between the docker container and your local 
# machine.
#
# Xpra allows to display the programs running inside of the
# container such as Firefox, LibreOffice, xterm, etc. 
# with disconnection and reconnection capabilities
#
# Xephyr allows to display the programs running inside of the
# container such as Firefox, LibreOffice, xterm, etc. 
#
# Fluxbox and ROX-Filer creates a very minimalist way to 
# manages the windows and files.
#
# Author: Roberto Gandolfo Hashioka
# Date: 07/28/2013


FROM golang:onbuild
MAINTAINER Matteo De Carlo "matteo.dek@gmail.com"

#RUN go get github.com/gorilla/mux
#RUN go get gopkg.in/mgo.v2

# Set the env variable DEBIAN_FRONTEND to noninteractive
#ENV HOST 
ENV PORT 80

EXPOSE 80
# Start xdm and ssh services.
CMD ["simpleBlogBackendGo"]

