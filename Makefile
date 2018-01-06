all: assets.go

assets.go:
	go-bindata -pkg templates -o templates/assets.go assets/...
