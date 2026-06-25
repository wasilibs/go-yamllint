FROM ghcr.io/astral-sh/uv:debian-slim

RUN apt-get update && apt-get install -y wget zip

WORKDIR /app

COPY buildtools/wasm/pysite .
COPY buildtools/wasm/version.txt .
RUN uv sync
RUN uv add yamllint==$(cat version.txt | awk '{$1=$1};1')
RUN wget https://github.com/brettcannon/cpython-wasi-build/releases/download/v3.14.6/python-3.14.6-wasi_sdk-24.zip
RUN unzip python-3.14.6-wasi_sdk-24.zip
RUN rm -r .venv/bin/python*
RUN rm -r .venv/lib64
RUN zip -r pysite.zip.wasm .venv lib

CMD ["cp", "-R", "python.wasm", "pysite.zip.wasm", "/out/"]
