# Author: Nitin Jilla

FROM golang:1.22-alpine AS build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o welcome

FROM scratch 
WORKDIR /
COPY --from=build /build/welcome /
EXPOSE 8000
ENTRYPOINT ["/welcome"]
