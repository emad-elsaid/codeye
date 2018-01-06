all: assets.go

assets.go:
	go-bindata -pkg views -o views/assets.go assets/...
