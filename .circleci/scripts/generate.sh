#@IgnoreInspection BashAddShebang
SERVICE="charon"
PROTO_INCLUDE="-I=./tmp/pb -I=./vendor/github.com/piotrkowalczuk -I=${GOPATH}/src"

: ${PROTOC:="${PWD}/tmp/bin/protoc"}

protobufs=(
    "rpc/${SERVICE}d/v1"
)

for protobuf in "${protobufs[@]}"
do
    case $1 in
        lint)
            ${PROTOC} ${PROTO_INCLUDE} --lint_out=. ${PWD}/pb/${protobuf}/*.proto
            ;;
        python)
            python -m grpc_tools.protoc ${PROTO_INCLUDE} --python_out=publish/python --grpc_python_out=publish/python ${PWD}/pb/${protobuf}/*.proto
            cp publish/python/github.com/piotrkowalczuk/charon/pb/${protobuf}/* publish/python/github/com/piotrkowalczuk/charon/pb/${protobuf}/
            rm -rf publish/python/github.com
            ;;
        java)
            rm -rf ./publish/java
            mkdir -p ./publish/java
            ${PROTOC} ${PROTO_INCLUDE} --java_out=./publish/java ${PWD}/pb/${protobuf}/*.proto
            ;;
        golang | go)
            ${PROTOC} ${PROTO_INCLUDE} --go-grpc_out=${GOPATH}/src --go_out=${GOPATH}/src ${PWD}/pb/${protobuf}/*.proto
            mockery -case=underscore -dir=./pb/${protobuf} -all -outpkg=$(basename $(dirname "./pb/${protobuf}mock"))mock -output=./pb/${protobuf}mock
            goimports -w ${PWD}/pb
            ;;
        swift)
            ${PROTOC} ${PROTO_INCLUDE} --swift_opt=Visibility=Public --swift_out=publish/swift --grpc-swift_out=publish/swift ${PWD}/pb/${protobuf}/*.proto
            mkdir -p publish/swift/Sources/Generated/charon/${protobuf}
            cp publish/swift/github.com/piotrkowalczuk/charon/pb/${protobuf}/* publish/swift/Sources/Generated/charon/${protobuf}/
            rm -rf publish/swift/github.com
            ;;
        *)
            echo "code generation failure due to unknown language: ${1}"
            exit 1
            ;;
    esac
done
