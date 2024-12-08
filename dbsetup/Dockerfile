FROM mongo:8.0.3-noble

WORKDIR /data
COPY states-data.json .

CMD ["mongoimport", "--uri=mongodb://mongodb:27017", "--db", "us_states", "--collection", "states", "--file", "states-data.json", "--jsonArray"]