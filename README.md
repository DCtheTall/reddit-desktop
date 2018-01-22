# Reddit Desktop

A subreddit scraper which sets your desktop background
written in Go
---

This project uses [dep](https://github.com/golang/dep) to ensure the
integrity of its dependencies. To install dependencies run

`dep ensure`

in the root of the project directory to install dependencies.

To build the app, run

`go install reddit-desktop`

To run, change the working directory to the directory
where the app was built and run

`reddit-desktop <subreddit1> <subreddit2> ...`

*Important note:*

This app will write (and promptly delete) an image file to
the directory it runs in. It will delete any image it successfully writes even if it encouter an error.

### Supported Operating Systems
- OSX

### TODO
- Windows support
- Linux support

Dependencies:
- golang.org/x/net/html
