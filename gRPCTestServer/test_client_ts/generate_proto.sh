PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"

# Directory to write generated code to (.js and .d.ts files)
OUT_DIR="."

PROTO_FILE="service.proto"

# protoc \
#     --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
#     --js_out="import_style=commonjs,binary:${OUT_DIR}" \
#     --grpc-web_out="import_style=commonjs,mode=grpcweb,out=echo_grpc_pb.js:generated" \
#     --ts_out="service=false:${OUT_DIR}" \
#     ${PROTO_FILE}

protoc ${PROTO_FILE} \
    --js_out="import_style=commonjs:generated" \
    --grpc-web_out="import_style=commonjs,mode=grpcweb,out=echo_grpc_pb.js:generated" \
