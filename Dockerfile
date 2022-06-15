FROM alpine:3.15
COPY trusted-ports /trusted-ports

RUN apk add --no-cache tini
# Tini is now available at /sbin/tini
ENTRYPOINT ["/sbin/tini", "--"]

# Run your program under Tini
CMD ["/trusted-ports"]
# ENTRYPOINT ["/httpserver"]