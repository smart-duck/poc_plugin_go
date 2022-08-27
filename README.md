# Build plugin

go build -buildmode=plugin -o plug1/plug.so plug1/plug.go

# Execute dummy pgm

go run *.go