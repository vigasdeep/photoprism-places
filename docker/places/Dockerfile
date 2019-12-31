FROM photoprism/places-dev:20191230 as build

# Set up project directory
WORKDIR "/go/src/github.com/photoprism/photoprism-places"
COPY . .

# Build PhotoPrism Places
RUN make dep install

# Same base image as photoprism/places-dev
FROM ubuntu:18.04

# Set environment variables
ENV DEBIAN_FRONTEND noninteractive

# Install additional distribution packages
RUN apt-get update && apt-get install -y --no-install-recommends \
        curl \
        unzip \
        nano \
        wget \
        ca-certificates \
        tzdata && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Create and configure places user
RUN useradd -ms /bin/bash places
USER places

WORKDIR /home/places
ENV PATH /home/places/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ENV PLACES_CONFIG_FILE /home/places/.config/places/places.yml

# Copy assets to /home/places
COPY --chown=places --from=build /root/.local/bin/places /home/places/.local/bin/places
COPY --chown=places --from=build /root/.local/share/places /home/places/.local/share/places
COPY --chown=places --from=build /root/.config/places /home/places/.config/places

# Create directories
RUN mkdir -p /home/places/.cache/places

# Fix permissions
USER root
RUN chown -R places:places /home/places
USER places

# Show places version
RUN places -v

# Expose http and database ports
EXPOSE 8080

# Run server
CMD places start
