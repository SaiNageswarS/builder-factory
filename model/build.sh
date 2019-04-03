export GOPATH=/Users/saisatch/myWorkspace/builder-factory
export PATH=$GOPATH/bin:$PATH
rm -Rf services
rm -Rf clientjs
mkdir services clientjs
protoc --go_out=plugins=grpc:services *.proto
#js client
protoc -I=. *.proto \
  --js_out=import_style=commonjs:clientjs \
  --grpc-web_out=import_style=commonjs+dts,mode=grpcweb:clientjs