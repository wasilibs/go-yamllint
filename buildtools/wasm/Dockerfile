FROM ghcr.io/astral-sh/uv:alpine

RUN apk add zip

COPY buildtools/wasm/pysite .
COPY buildtools/wasm/version.txt .
RUN uv sync
RUN uv add yamllint==$(cat version.txt | awk '{$1=$1};1')
RUN wget https://github.com/brettcannon/cpython-wasi-build/releases/download/v3.13.7/python-3.13.7-wasi_sdk-24.zip
RUN unzip python-3.13.7-wasi_sdk-24.zip
RUN rm -r .venv/lib64
RUN zip -r pysite.zip.wasm .venv lib

CMD ["cp", "-R", "python.wasm", "pysite.zip.wasm", "/out/"]
