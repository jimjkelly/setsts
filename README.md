# setsts
A simple tool for setting AWS STS tokens

## What does this do?

Currently not much - in truth this project is more for me to mess with Go a bit.
Basically this reformats output of STS token generation so it can be copied
directly into your `~/.aws/credentials` file.

## How to install

Man who knows.  I guess clone the repo and then `go get ./...` to install
dependencies and then `go build -o setsts main.go`.  Still not up to speed
on how this Go stuff works.

## Usage

Usage flags can be seen by running `setsts -h` but here's the gist:

```
$ setsts --serial-number <YOUR MFA SERIAL> 123456
```

where `123456` is your MFA token.