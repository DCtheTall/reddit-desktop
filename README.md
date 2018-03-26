# Reddit Desktop

A subreddit scraper which sets your desktop background
written in Go
---

This project uses [dep](https://github.com/golang/dep) to ensure the
integrity of its dependencies. To install dependencies run

`dep ensure`

in the root of the project.

To build the app, run

`sh build.sh`

It will compile into a directory in the root
of this project called `release`. It is recommended
that you run the project from this directory so
that it can write to the `release/data/` subdirectory
it creates during the build.

To run (from project root), run the command

`reddit-desktop <subreddit1> <subreddit2> ...`

The program will choose one of the provided
subreddits at random and scrape an image from
its front page.

### Options

Running the executable with the optional parameter `--save`,

ex: `reddit-desktop --save <subreddit1> <subreddit2> ...`

will cache the images in the `data/` directory
in the releases folder in case you want to
use some of the images you like later.

Running the executable with the parameter `--empty`,

ex: `reddit-desktop --empty`

will empty the images out of the `data/` directory.
This is a hard delete, so use at your own peril.

### Supported Operating Systems
- OSX

### TODO
- Windows support
- Linux support
- Windows build script

Dependencies:
- golang.org/x/net/html
