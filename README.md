# ddlcpad

## What is this

ddlcpad is short of *Doki Doki Literature Club Plus Asset Decrypter*

You can decrypt the `*.cy` file from *Doki Doki Literature Club Plus* with this cli binary.

After decrypt, you will get a `*.out` file. Put it into [Perfare/AssetStudio](https://github.com/Perfare/AssetStudio/) to get the assert file.

## How to install

You can download a binary file on [release](https://github.com/Perfare/AssetStudio/releases). Or you can [build](https://github.com/aimerneige/DDLC-Plus-Asset-Decrypter#how-to-build) it yourself.

## How to use

Just open your terminal (or powershell on windows), run the binary file.

Basic usage:

```bash
# On Linux
./ddlcpad <file_name>
```

```powershell
# On windows
.\ddlcpad.exe <file_name>
```

Use `--help` or `-h` to get more help.

## How to build

Just install and config [golang](https://golang.org/doc/install) on your computer, then run this command:

```bash
# On windows
go build -o ddlcpad.exe
```

```bash
# On Linux
go build -o ddlcpad
```

If you wants to build exe file on Linux:

```bash
export GOARCH=amd64
export GOOS=windows
go build -o ddlcpad.exe
```

## Why do this

There has a [project](https://github.com/MlgmXyysd/DDLC-Plus-Asset-Decrypter) do this. But I don't use php and don't want to install it on my computer. <sub>~~Actually, I don't know how to setup and use the php code. And also lazy to learn it.~~</sub>

What's more, php script for those who doesn't lean programming are very very hard to use. I wants to create a binary file which is easy to use. So that people can just download a binary file from release and use it without build or run code themselves.

## Why read file with `ioutil.ReadFile()`

The biggest assert file is `gallery_images.cy` and with `582 MB`, There is no way to use too much memory.

PR welcome. :)

## Open Source

**Refer <https://github.com/MlgmXyysd/DDLC-Plus-Asset-Decrypter>**

**Cli <https://github.com/urfave/cli>**

## More Question

Post your question on issue please.
