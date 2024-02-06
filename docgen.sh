env USER='$USER' go run handler/storage/gen/main.go
rm -rf docs/en/cli-reference
env USER='$USER' go run docs/gen/clireference/main.go

