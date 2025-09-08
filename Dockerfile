FROM busybox:glibc
WORKDIR /
ADD ./rehabber-discord-integration /rehabber-discord-integration
ADD ./ui /ui

ENTRYPOINT ["/rehabber-discord-integration"]
