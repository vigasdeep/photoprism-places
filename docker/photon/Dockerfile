FROM openjdk:11

LABEL maintainer="Michael Mayer <michael@liquidbytes.net>"

# Set environment variables
ENV DEBIAN_FRONTEND noninteractive

# Install additional distribution packages
RUN apt-get update && apt-get install -y --no-install-recommends \
        wget \
        ca-certificates \
        tzdata && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Create and configure user
RUN useradd -ms /bin/bash photon
USER photon

WORKDIR /home/photon

# Install Photon
ENV PHOTON_VERSION 0.3.2
RUN set -eux; \
	\
	url="https://github.com/komoot/photon/releases/download/${PHOTON_VERSION}/photon-${PHOTON_VERSION}.jar"; \
	wget -O photon.jar "$url"

# Expose HTTP port 2322
EXPOSE 2322

# Start server
CMD java -jar photon.jar -listen-port 2322
