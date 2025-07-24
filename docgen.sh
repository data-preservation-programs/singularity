
env USER='$USER' go run singularity.go
rm -rf docs/en/cli-reference
env USER='$USER' go run singularity.go

