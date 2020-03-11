###########################################################################
# Stage 1 Start
###########################################################################
FROM golang AS build-golang

# Maintainer
LABEL MAINTAINER = "Rian Eka Cahya <rian.eka.cahya@gmail.com>"

################################
# Build IMS Service:
################################
WORKDIR /usr/share/kumparan/news

COPY  . .

RUN make deploy

###########################################################################
# Stage 2 Start
###########################################################################
FROM ubuntu:18.04

# Copy Binary
COPY --from=build-golang /usr/share/kumparan/news/bin /usr/share/kumparan/news/bin/

WORKDIR /usr/share/kumparan/news

# Create group and user to the group
RUN groupadd -r kumparan && useradd -r -s /bin/false -g kumparan kumparan

# Set ownership golang directory
RUN chown -R kumparan:kumparan /usr/share/kumparan/news

# Make docker container rootless
USER kumparan
