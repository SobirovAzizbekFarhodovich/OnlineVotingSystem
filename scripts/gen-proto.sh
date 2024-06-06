#!/bin/bash
CURRENT_DIR=$1
rm -rf ${CURRENT_DIR}/genproto
mkdir -p ${CURRENT_DIR}/genproto

for x in $(find ${CURRENT_DIR}/protos -type f -name "*.proto"); do
  protoc -I=${CURRENT_DIR}/protos --go_out=${CURRENT_DIR}/genproto \
    --go-grpc_out=${CURRENT_DIR}/genproto ${x}
done
