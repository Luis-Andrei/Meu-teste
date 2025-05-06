# Build stage
FROM python:3.8-slim as builder

WORKDIR /app

COPY requirements.txt .
RUN pip install --user -r requirements.txt

# Production stage
FROM python:3.8-slim

WORKDIR /app

COPY --from=builder /root/.local /root/.local
COPY . .

ENV PATH=/root/.local/bin:$PATH
ENV FLASK_APP=app.py
ENV FLASK_ENV=production

EXPOSE 5000

CMD ["flask", "run", "--host=0.0.0.0"] 