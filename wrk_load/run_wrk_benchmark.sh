#!/bin/bash

BASE_URL="${BASE_URL:-http://localhost:8080}"
RESULTS_DIR="./wrk_results"
SCENARIOS_DIR="./wrk_scenarios"

export WRK_THREADS=24
export WRK_CONNECTIONS=10000
export WRK_DURATION="30s"

if [ ! -d "$SCENARIOS_DIR" ]; then
	echo "Scenarios directory not found: $SCENARIOS_DIR"
	exit 1
fi

mkdir -p "$RESULTS_DIR"
echo "Results directory: $RESULTS_DIR"

echo "Checking server availability at $BASE_URL"
if ! curl -s -f "$BASE_URL/healthz" > /dev/null 2>&1; then
	echo "Server not reachable at $BASE_URL/healthz"
	exit 1
fi
echo "Server is up"


echo "Warmup - POST /objects"
output_file="${RESULTS_DIR}/01_warmup_post.txt"

wrk -t "$WRK_THREADS" -c "$WRK_CONNECTIONS" -d "$WRK_DURATION" -s "$SCENARIOS_DIR/warmup_post.lua" "$BASE_URL" > "$output_file" 2>&1

echo "Results saved to: $output_file"
sleep 10


echo  "GET /objects/:id"
output_file="${RESULTS_DIR}/02_read_by_id.txt"

wrk -t "$WRK_THREADS" -c "$WRK_CONNECTIONS" -d "$WRK_DURATION" -s "$SCENARIOS_DIR/read_by_id.lua" "$BASE_URL" > "$output_file" 2>&1

echo "Results saved to: $output_file"
sleep 10


echo  "GET /objects?status=true"
output_file="${RESULTS_DIR}/03_read_by_status.txt"

wrk -t "$WRK_THREADS" -c "$WRK_CONNECTIONS" -d "$WRK_DURATION" -s "$SCENARIOS_DIR/read_by_status.lua" "$BASE_URL" > "$output_file" 2>&1

echo "Results saved to: $output_file"
sleep 10


echo "GET /objects"
output_file="${RESULTS_DIR}/04_read_all.txt"

wrk -t "$WRK_THREADS" -c "$WRK_CONNECTIONS" -d "$WRK_DURATION" -s "$SCENARIOS_DIR/read_all.lua" "$BASE_URL" > "$output_file" 2>&1

echo "Results saved to: $output_file"
sleep 10


echo "GET /healthz"
output_file="${RESULTS_DIR}/05_healthz.txt"

wrk -t "$WRK_THREADS" -c "$WRK_CONNECTIONS" -d "$WRK_DURATION" -s "$SCENARIOS_DIR/healthz.lua" "$BASE_URL" > "$output_file" 2>&1

echo "Results saved to: $output_file"
