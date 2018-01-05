# Codeye (Code Eye)

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

## Development

### Prerequisite

* You need of course Go installed on your system with `$GOPATH` defined in your shell

### Having your own fork

* Fork the repository to your account
* Get the repository `go get github.com/emad-elsaid/codeye`
* Go there `cd $GOPATH/src/github.com/emad-elsaid/codeye`
* Rename remote `origin` to `upstream` `git remote rename origin upstream`
* Add your repository as origin `git remote add origin <your repository remote url here>`
* Happy hacking

### Running codeye in another project

I usually execute that command while developing, to run codeye in another project called `news`

```shell
cd ~/code/news; go run ~/go/src/github.com/emad-elsaid/codeye/cmd/codeye/codeye.go
```
