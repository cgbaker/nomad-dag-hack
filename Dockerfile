FROM scratch
COPY nomad-dag-hack /
ENTRYPOINT ["/nomad-dag-hack"]
