FROM scratch
COPY hivelime /usr/local/bin/hivelime
ENTRYPOINT [ "/usr/local/bin/hivelime" ]