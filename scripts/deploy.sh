#!/bin/bash
# Deploy BasketForm-AI to VM
# Run from repository root: ./scripts/deploy.sh
#
# Environment variables:
#   REMOTE_USER  — SSH user (default: root)
#   REMOTE_HOST  — target host (default: 80.74.30.14)
#   REMOTE_DIR   — remote app directory (default: /opt/basketform-ai)
#   SKIP_ML_DEPS — set to 1 to skip ML dependency installation

set -euo pipefail

REMOTE_USER="${REMOTE_USER:-root}"
REMOTE_HOST="${REMOTE_HOST:-80.74.30.14}"
REMOTE_DIR="${REMOTE_DIR:-/opt/basketform-ai}"
BINARY="bin/server"

echo "=== [1/6] Building for Linux ==="
GOOS=linux GOARCH=amd64 go build -race -o "${BINARY}" ./cmd/server/
echo "Binary built: $(ls -lh ${BINARY} | awk '{print $5}')"

echo "=== [2/6] Uploading to ${REMOTE_HOST} ==="
ssh "${REMOTE_USER}@${REMOTE_HOST}" "mkdir -p ${REMOTE_DIR}/{uploads,results,data,web,ML,scripts}"
scp "${BINARY}" "${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/"
scp -r web/templates "${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/web/"
scp -r web/static "${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/web/"
scp -r ML/*.py "${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/ML/"
scp requirements.txt "${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/"

echo "=== [3/6] Installing ML dependencies ==="
if [ "${SKIP_ML_DEPS:-0}" != "1" ]; then
    ssh "${REMOTE_USER}@${REMOTE_HOST}" \
        "cd ${REMOTE_DIR} && pip install -q -r requirements.txt" 2>&1 | tail -3
else
    echo "Skipped (--skip-ml-deps)"
fi

echo "=== [4/6] Restarting service ==="
ssh "${REMOTE_USER}@${REMOTE_HOST}" \
    "cd ${REMOTE_DIR} && pkill -f 'bin/server' || true; sleep 1; \
     PORT=80 nohup ./bin/server > /var/log/basketform-ai.log 2>&1 &"
sleep 2

echo "=== [5/6] Smoke test ==="
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" "http://${REMOTE_HOST}/login")
if [ "$HTTP_CODE" = "200" ] || [ "$HTTP_CODE" = "302" ]; then
    echo "OK: HTTP ${HTTP_CODE}"
else
    echo "FAIL: HTTP ${HTTP_CODE}"
    echo "Check logs: ssh ${REMOTE_USER}@${REMOTE_HOST} 'tail -20 /var/log/basketform-ai.log'"
    exit 1
fi

echo "=== [6/6] Upload status ==="
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" "http://${REMOTE_HOST}/api/videos")
echo "API /api/videos: HTTP ${HTTP_CODE}"

echo ""
echo "=== Deploy complete: http://${REMOTE_HOST}/ ==="
