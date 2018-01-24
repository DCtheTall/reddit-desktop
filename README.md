# Reddit Desktop

A subreddit scraper which sets your desktop background
written in Go
---

This project uses [dep](https://github.com/golang/dep) to ensure the
integrity of its dependencies. To install dependencies run

`dep ensure`

in the root of the project directory to install dependencies.

To build the app, run

`sh build.sh`

It will compile into a directory in the root
of this project called `release`.

To run, change the working directory to the directory
where the app was built and run

`reddit-desktop <subreddit1> <subreddit2> ...`

### Supported Operating Systems
- OSX

### TODO
- Windows support
- Linux support
- Windows build script

Dependencies:
- golang.org/x/net/html
