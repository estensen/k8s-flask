FROM python:3.7.4-slim as base
ENV PIP_NO_CACHE_DIR=off \
    PIP_DISABLE_PIP_VERSION_CHECK=on \
    PIP_DEFAULT_TIMEOUT=100 \
    POETRY_PATH=/opt/poetry \
    VENV_PATH=/opt/venv
ENV PATH="$POETRY_PATH/bin:$VENV_PATH/bin:$PATH"


FROM base as poetry
RUN apt-get update \
    && apt-get install --no-install-recommends -y \
        curl \
        build-essential \
    && curl -sSL https://raw.githubusercontent.com/sdispater/poetry/master/get-poetry.py | python \
    && mv /root/.poetry $POETRY_PATH \
    && poetry --version \
    \
    # configure poetry & make a virtualenv ahead of time since we only need one
    && python -m venv $VENV_PATH \
    && poetry config settings.virtualenvs.create false \
    && rm -rf /var/lib/apt/lists/*

COPY poetry.lock pyproject.toml ./
RUN poetry install --no-interaction --no-ansi -vvv


FROM base as runtime
WORKDIR /app

COPY --from=poetry $VENV_PATH $VENV_PATH
COPY . ./

CMD ["python", "server.py"]
