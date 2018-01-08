all: compile_assets

compile_assets:
	go-bindata -pkg codeye -o assets.go assets/...
