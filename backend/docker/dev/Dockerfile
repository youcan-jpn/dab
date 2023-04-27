FROM golang:1.20.1-bullseye

RUN apt-get update && \
    apt-get install -y --no-install-recommends\
        mariadb-client \
        uuid-runtime \
        curl && \
    apt-get -y clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /dab

# CMD ["go", "run", "cmd/dab/main.go"]
CMD ["/bin/bash"]
