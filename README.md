# Build plugin

go build -buildmode=plugin -o plug1/plug.so plug1/plug.go

# Execute dummy pgm

go run *.go

# To read

Top:
https://github.com/douglasmakey/admissioncontroller

file2configmap
https://github.com/go-guoyk/file2configmap/blob/master/main.go
# Package a plugin as a tgz/base64

[Not retained] tar cvzf plug1/plug.tgz -C plug1/ plug.so

tar cvzf plug1/plug.tgz plug1/plug.so

base64 plug1/plug.tgz > plug1/plug1.tgz.base64

[Not retained] Plugin library .so must be at root path.

[Not retained] cp -f plug1/plug.yaml.tmpl plug1/plug.yaml;base64 plug1/plug.tgz >> plug1/plug.yaml

cp -f plug1/plug.yaml.tmpl plug1/plug.yaml;base64 plug1/plug.tgz | sed -E 's/(^.+$)/  \1/g' >> plug1/plug.yaml

# Docker

sudo docker build -t veradco/dummy:0.1 .
sudo docker run --rm veradco/dummy:0.1

# Bulk

kubectl -n ns-test create configmap plug1 --from-file=plug1/plug.so