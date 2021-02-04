FROM golang as builder
ENV CGO_ENABLED=0
WORKDIR /treeserver
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \ 
    make
RUN mkdir /new_tmp

FROM scratch
COPY --from=builder /treeserver/treeserver /treeserver
COPY --from=builder /new_tmp /tmp
ENTRYPOINT [ "/treeserver" ]
