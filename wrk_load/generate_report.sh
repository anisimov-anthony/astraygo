#!/bin/bash

RESULTS_DIR="${1:-./wrk_results}"
REPORT_FILE="${RESULTS_DIR}/BENCHMARK_REPORT.md"

if [ ! -d "$RESULTS_DIR" ]; then
    echo "Error: Results directory not found: $RESULTS_DIR"
    exit 1
fi

echo "Generating benchmark report from: $RESULTS_DIR"


cat > "$REPORT_FILE" << EOF
# WRK Benchmark Report - AstrayGo

**Generated:** $(date '+%Y-%m-%d %H:%M:%S %Z')

## System Specifications

**CPU:**
$(lscpu | grep "Model name" || grep "model name" /proc/cpuinfo | head -1 | cut -d: -f2)
- Cores: $(nproc)
- Threads: $(lscpu | grep "^CPU(s):" | awk '{print $2}' || nproc)

**Memory:**
$(free -h | grep "Mem:" | awk '{print "Total: " $2 ", Available: " $7}')

**Operating System:**
$(cat /etc/os-release | grep PRETTY_NAME | cut -d= -f2 | tr -d '"')

**Kernel:**
$(uname -r)

**Architecture:**
$(uname -m)

---

EOF

for result_file in $(ls "$RESULTS_DIR"/*.txt 2>/dev/null | sort); do
    filename=$(basename "$result_file" .txt)

    if [[ $filename =~ warmup_t([0-9]+)_c([0-9]+) ]]; then
        test_name="Warmup - POST /objects"
    elif [[ $filename =~ read_by_id_t([0-9]+)_c([0-9]+) ]]; then
        test_name="GET /objects/:id"
    elif [[ $filename =~ read_by_status_t([0-9]+)_c([0-9]+) ]]; then
        test_name="GET /objects?status=true"
    elif [[ $filename =~ read_all_t([0-9]+)_c([0-9]+) ]]; then
        test_name="GET /objects"
    elif [[ $filename =~ healthz_t([0-9]+)_c([0-9]+) ]]; then
        test_name="GET /healthz"
    else
        test_name="$filename"
        threads="?"
        connections="?"
    fi

    cat >> "$REPORT_FILE" << EOF
## $test_name


\`\`\`
$(cat "$result_file")
\`\`\`

---

EOF

done
