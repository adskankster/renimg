# Exifsorter

Exifsorter is a command-line tool to merge and rename images from multiple
sources, so that you can view them in the order of their creation date. It reads
images (jpg) from a directory, sorts them ascending by their exif creation date
and renames them according to a specified prefix.

[![Build Status](https://drone.io/github.com/mrauh/exifsorter/status.png)](https://drone.io/github.com/mrauh/exifsorter/latest)

## Use case

Think of the following situation: there is a party and you and some of your
friends take pictures with your digicams and smartphones. After that, all those
pictures are shared among each other. Now you have - say 5 - directories on your
computer with images of yourself and from your friends and you want to watch
them in the order of their creation date (and not folder after folder).

So you copy all images in a single directory. But the "last modified date" is
the date where those pictures were copied to disk so there is little you can do
to accomplish the right order.

Of course there are "image viewers" that can sort images by their exif creation
date (e.g. digikam on linux), but there are a lot who can't. Or suppose you
want to copy the files on a sd card to view them on TV.

With exifsorter you can rename all images in a directory after they are sorted
by their creation date. The result is that the alphabetical order corresponds to
the exif creation date.

Exifsorter is tested on linux, but should also work on the other plattforms
that are supported by Go (e.g. OS X and Windows).

## Installation

With [Go](http://www.golang.org) installed on your machine:

	$ go get github.com/mrauh/exifsorter

## Downloads

### Linux
* [linux_386](http://www.linux-quiz.de/go/downloads/exifsorter/linux_386/exifsorter)
* [linux_amd64](http://www.linux-quiz.de/go/downloads/exifsorter/linux_amd64/exifsorter)

### OS X
* [darwin_386](http://www.linux-quiz.de/go/downloads/exifsorter/darwin_386/exifsorter)
* [darwin_amd64](http://www.linux-quiz.de/go/downloads/exifsorter/darwin_amd64/exifsorter)

### Windows
* [windows_386](http://www.linux-quiz.de/go/downloads/exifsorter/windows_386/exifsorter.exe)
* [windows_amd64](http://www.linux-quiz.de/go/downloads/exifsorter/windows_amd64/exifsorter.exe)

## Usage / Example

Assuming that the binary is in your 'PATH', just navigate to the directory where
your jpg files are and run 'exifsorter'. If you don't specify any flag, the
current directory is used and you are asked to provide a prefix for the
resulting filenames.

## Flags

You can also specify the following flags:

* -dir: Directory of the images to sort
* -p: Prefix of the renamed files

	$ exifsorter -dir="path/to/images" -p="imagesMerged"

## FAQ

**Why are the images still not in the right order?**

Maybe the date from the digicam / smartphone of one of the sources was not
set correctly at the time the pictures were made.

**I get an "EOF" error. What can I do?**

Unfortunately nothing. There seems to be a problem with (very) old images.
# renimg
