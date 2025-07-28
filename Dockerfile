ARG PHP_VERSION

FROM ghcr.io/shopware/heycart-cli-base:${PHP_VERSION}

COPY heycart-cli /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/heycart-cli"]
CMD ["--help"]
