#!/bin/bash
# Deploy BasketForm-AI to VM
# Run this from the repository root after pushing to main

set -e

REMOTE_USER="${REMOTE_USER:-root}"
REMOTE_HOST="${REMOTE_HOST:-80.74.30.14}"
REMOTE_DIR="/opt/basketform-ai"
BINARY="bin/server"

echo "=== Building for Linux ==="
GOOS=linux GOARCH=amd64 go build -o ${BINARY} ./cmd/server/

echo "=== Uploading to ${REMOTE_HOST} ==="
scp ${BINARY} ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/

echo "=== Uploading static files and templates ==="
scp -r web/templates ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/web/
scp -r web/static ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/web/
scp -r scripts ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/
scp requirements.txt ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/

echo "=== Restarting service ==="
ssh ${REMOTE_USER}@${REMOTE_HOST} "cd ${REMOTE_DIR} && pkill -f basketform-ai || true; PORT=80 nohup ./${BINARY} > /var/log/basketform-ai.log 2>&1 &"

echo "=== Verifying ==="
sleep 2
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" http://${REMOTE_HOST}/login)
if [ "$HTTP_CODE" = "200" ] || [ "$HTTP_CODE" = "302" ]; then
    echo "=== Deploy successful! http://${REMOTE_HOST}/ ==="
else
    echo "=== WARNING: Got HTTP ${HTTP_CODE}, check logs ==="
fi
