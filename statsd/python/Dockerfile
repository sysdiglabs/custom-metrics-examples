FROM python:alpine
WORKDIR /app
COPY main.py /app/
COPY requirements.txt /app/
RUN cd /app && pip install -r requirements.txt
CMD ["python", "/app/main.py"]