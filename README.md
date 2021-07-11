# ddlcpd

## **Refer <https://github.com/MlgmXyysd/DDLC-Plus-Asset-Decrypter>**

## What is this

ddlcpad is short of *Doki Doki Literature Club Plus Asset Decrypter*

## How to install

You can download a binary file on release. Or you can build it yourself.

## How to use

Just open your terminal (or powershell on windows), run the binary file.

## How to build

Just install and config go on your computer, then run this command:

```bash
# On windows
go build -o ddlcpad.exe
```

If you wants to build linux binary on windows:

```bash
set GOARCH=amd64
set GOOS=linux
go build -o ddlcpad
```

## Why do this

There has a [project](https://github.com/MlgmXyysd/DDLC-Plus-Asset-Decrypter) do this. But I don't use php and don't want to install it on my computer. Actually, I don't know how to setup and use this php code.

What's more, php script for those who doesn't lean programming are very very hard to use. I wants to create a binary file which is easy to use. So that people can just download a binary file from release and use it without build or run code themselves.

## Why read file with `ioutil.ReadFile()`

The biggest assert file is `gallery_images.cy` and with `582 MB`, There is no way to use too much memory.

PR welcome. :)

## More Question

Post your question on issue please.