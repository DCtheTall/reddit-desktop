BUILD_GOOS=(darwin)
BUILD_GOARCHS=(amd64 386)

[ -e release/ ] && rm -r release/;
mkdir -p release/;

[ -e release/data/ ] && rm -r release/data/;
mkdir -p release/data/;

for GOOS in ${BUILD_GOOS[@]}; do
  for GOARCH in ${BUILD_GOARCHS[@]}; do
    NAME="reddit-desktop-$GOOS-$GOARCH"
    if [ "$OS" == "windows" ]; then
      NAME="$NAME.exe"
    fi
    go build -o "release/$NAME" -a ./cmd/main.go
  done
done
