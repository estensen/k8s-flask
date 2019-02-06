FROM python:3.7.2-alpine

WORKDIR /app

COPY poetry.lock pyproject.toml ./
RUN pip install --no-cache-dir --upgrade pip \
 && pip install --no-cache-dir poetry \
 && poetry config settings.virtualenvs.create false \
 && poetry install --no-dev \
 && pip uninstall --yes poetry

COPY . /app

EXPOSE 5000
CMD ["python", "server.py"]

