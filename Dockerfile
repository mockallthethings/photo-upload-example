FROM scratch

COPY photo-upload-example /photo-upload-example
COPY public /public

ENTRYPOINT ["/photo-upload-example"]
