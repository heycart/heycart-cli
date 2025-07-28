ARG PHP_VERSION

FROM ghcr.io/heycart/heycart-cli-base:${PHP_VERSION}

COPY heycart-cli /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/heycart-cli"]
CMD ["--help"]
