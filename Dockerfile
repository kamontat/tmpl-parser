FROM scratch

WORKDIR /app

# copy compiled script
COPY tmpl-parser /app

ENTRYPOINT [ "/app/tmpl-parser"]
