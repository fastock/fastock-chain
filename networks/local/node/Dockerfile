FROM golang:stretch as build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y $PACKAGES

WORKDIR /fastock-chain
# Add source files
COPY . .

# build faStock-Chain
RUN make build-linux

# Final image
FROM golang:1.14 as final

WORKDIR /fastock-chain-daemon
# Copy over binaries from the build-env
COPY --from=build-env /fastock-chain/build/fastock-chain-daemon /usr/bin/fastock-chain-daemon
COPY --from=build-env /fastock-chain/build/fastock-chain-cli /usr/bin/fastock-chain-cli
COPY --from=build-env /fastock-chain/networks/local/node/wrapper.sh /usr/bin/wrapper.sh

EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start"]
STOPSIGNAL SIGTERM

