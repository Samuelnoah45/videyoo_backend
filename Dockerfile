
# Final Stage
FROM debian:stable-slim

RUN apt update && apt -y upgrade
RUN apt install -y wget
RUN apt -y install libjpeg62-turbo-dev
RUN apt-get update \
    && apt-get install -y \
        curl \
        libxrender1 \
        libjpeg62-turbo \
        fontconfig \
        libxtst6 \
        xfonts-75dpi \
        xfonts-base \
        xz-utils

# Copy built artifacts from the build stage
# COPY --from=build /server/server /server/server
# COPY --from=build /server/publicKey.pem /server/publicKey.pem
# COPY --from=build /server/privateKey.pem /server/privateKey.pem

RUN wget http://archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.0g-2ubuntu4_amd64.deb
RUN dpkg -i libssl1.1_1.1.0g-2ubuntu4_amd64.deb

RUN curl "https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.buster_amd64.deb" -L -o "wkhtmltopdf.deb"
RUN dpkg -i wkhtmltopdf.deb