# First stage: build the executable.
FROM golang:1.15.0-alpine3.12 AS builder

# git is required to fetch go dependencies
RUN apk add --no-cache ca-certificates git
RUN apk add --no-cache ca-certificates tree
RUN go env -w GOPRIVATE=github.com/AyokunlePaul/crud-pay-api

# Create the user and group files that will be used in the running
# container to run the process as an unprivileged user.
RUN mkdir /user && echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && echo 'nobody:x:65534:' > /user/group

# Copy the predefined netrc file into the location that git depends on
COPY ./.netrc /root/.netrc
RUN chmod 600 /root/.netrc

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Import the code from the context.
COPY src/ src/
RUN tree

# Build the executable to `/app`. Mark the build as statically linked.
RUN go build src/main.go
RUN ls

# Final stage: the running container.
FROM scratch AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the first stage.
COPY --from=builder /src/main /src

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Run the compiled binary.
ENTRYPOINT ["./main"]