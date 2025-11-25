CONTAINER_DIR=/home/naokotani/.config/containers/coredns/
SRC_DIR=/home/naokotani/src/coredns

cd $SRC_DIR

git pull

cd $CONTAINER_DIR

podman run --rm -i -t \
  -v $SRC_DIR:/go/src/github.com/coredns/coredns:Z -w /go/src/github.com/coredns/coredns \
  golang:1.24 sh -c 'GOFLAGS="-buildvcs=false" make gen && GOFLAGS="-buildvcs=false" make'

cp $SRC_DIR/coredns $CONTAINER_DIR
