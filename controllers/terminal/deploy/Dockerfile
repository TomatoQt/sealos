FROM scratch

USER 65532:65532

COPY manifests ./terminal-controller
CMD ["kubectl apply -f terminal-controller/deploy.yaml"]
