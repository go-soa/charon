echo "protoc installation in ${PWD}/tmp/bin"

curl -L https://github.com/google/protobuf/releases/download/v28.2/protoc-28.2-$1-x86_64.zip > protoc.zip

rm -rf ./tmp/protoc ./tmp/pb/google
mkdir -p ./tmp/protoc ./tmp/bin ./tmp/pb/google

unzip protoc.zip -d ./tmp/protoc

mv -f ./tmp/protoc/bin/protoc ./tmp/bin/protoc
mv -f ./tmp/protoc/include/google ./tmp/pb

rm -rf ./tmp/protoc

./tmp/bin/protoc --version
