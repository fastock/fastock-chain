# Simple usage with a mounted data directory:
# > docker build -t fastock-chain .
# > docker run -it -p 36657:36657 -p 36656:36656 -v ~/.fastock-chain-daemon:/root/.fastock-chain-daemon -v ~/.fastock-chain-cli:/root/.fastock-chain-cli fastock-chain fastock-chain-daemon init mynode
# > docker run -it -p 36657:36657 -p 36656:36656 -v ~/.fastock-chain-daemon:/root/.fastock-chain-daemon -v ~/.fastock-chain-cli:/root/.fastock-chain-cli fastock-chain fastock-chain-daemon start
FROM golang:alpine AS build-env

# Install minimum necessary dependencies, remove packages
RUN apk add --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev

# Set working directory for the build
WORKDIR /go/src/github.com/fastock/fastock-chain

# Add source files
COPY . .

# Build OKExChain
RUN GOPROXY=http://goproxy.cn make install

# Final image
FROM alpine:edge

WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/fastock-chain-daemon /usr/bin/fastock-chain-daemon
COPY --from=build-env /go/bin/fastock-chain-cli /usr/bin/fastock-chain-cli

# Run fastock-chain-daemon by default, omit entrypoint to ease using container with fastock-chain-cli
CMD ["fastock-chain-daemon"]
