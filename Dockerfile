FROM photoprism/places-dev:20191230

# Set up project directory
WORKDIR "/go/src/github.com/photoprism/photoprism-places"
COPY . .
