FROM golang:1.20.5 as builder

RUN apt-get update && apt-get install -y

WORKDIR /DharitriOne
COPY . .

WORKDIR /DharitriOne/cmd/elasticindexer

RUN go build -o elasticindexer

# ===== SECOND STAGE ======
FROM ubuntu:22.04
RUN apt-get update && apt-get install -y

RUN useradd -m -u 1000 appuser
USER appuser

COPY --from=builder /DharitriOne/cmd/elasticindexer /DharitriOne

EXPOSE 22111

WORKDIR /DharitriOne

ENTRYPOINT ["./elasticindexer"]
CMD ["--log-level", "*:DEBUG"]
