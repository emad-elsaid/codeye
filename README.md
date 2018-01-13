# Codeye (Code Eye)

<p align="center">
<img src="http://www.emadelsaid.com/images/codeye.png" width="150" align="center">
</p>

When I join a new project I always find it hard to understand the current code
structure, its history like when was it started, who worked on it so far, and
other helpful tips that gives me context about code before diving in the code
itself, so I decided to build that tool, to be a single binary web server, you
can run it in any project, and explore that project code from a web interface.

The technique itself helped me discover anomalies in code, especially what is so
called legacy code, get an idea about the messy places and have a general plan
to straighten the project before even diving in the refactoring phase.

In the beginning I didn't know whether to choose ruby or Go for the
implementation, ruby will make it faster to implement it with nice help from
rails and other gems, and go will give me speed and one binary distribution, so
I decided to go with Go for these reasons, also I wanted to push my skills
towards Go and explore more of this language.

## Installation

Download the binary for your system from here:

| Windows                                                                                                                                                   | Linux                                                                                                                                               | FreeBSD                                                                                                                                                   | MacOS                                                                                                                                               |
|-----------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| [![windows](http://www.emadelsaid.com/images/windows.png)](https://github.com/emad-elsaid/codeye/releases/download/v0.2/codeye_0.2_windows_64-bit.tar.gz) | [![linux](http://www.emadelsaid.com/images/linux.png)](https://github.com/emad-elsaid/codeye/releases/download/v0.2/codeye_0.2_linux_64-bit.tar.gz) | [![freebsd](http://www.emadelsaid.com/images/freebsd.png)](https://github.com/emad-elsaid/codeye/releases/download/v0.2/codeye_0.2_freebsd_64-bit.tar.gz) | [![macos](http://www.emadelsaid.com/images/macos.png)](https://github.com/emad-elsaid/codeye/releases/download/v0.2/codeye_0.2_macOS_64-bit.tar.gz) |

Or latest version from [here](https://github.com/emad-elsaid/codeye/releases/latest)

## Usage

Execute `codeye` in any git project you have, it will open a port service a web insterface for you to navigate

## Development

### Prerequisite

* You need of course Go installed on your system with `$GOPATH` defined in your shell
* `go-bindata` installed with `go get -u github.com/jteeuwen/go-bindata/...`
* [goreleaser](https://goreleaser.com/) for releasing, if you don't want to release then ignore it,

### Having your own fork

* Fork the repository to your account
* Get the repository `go get github.com/emad-elsaid/codeye`
* Go there `cd $GOPATH/src/github.com/emad-elsaid/codeye`
* Rename remote `origin` to `upstream` `git remote rename origin upstream`
* Add your repository as origin `git remote add origin <your repository remote url here>`
* Happy hacking

### Assets

Codeye uses `go-bindata` to generate assets, so if you changed any of the assets you'll need to regenerate them in code.

```bash
go-bindata -pkg views -o templates/assets.go assets/...
```


### Running codeye in another project

I usually execute that command while developing, to run codeye in another project called `news`

```shell
cd ~/code/news; go run ~/go/src/github.com/emad-elsaid/codeye/cmd/codeye/codeye.go
```

### Releasing a new version

1. add a tag to the branch and push it
2. execute `goreleaser`
