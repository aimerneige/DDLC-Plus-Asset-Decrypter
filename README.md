# ddlcpad
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[įŽäŊä¸­æ](./README-CN.md)

## What is this đ¤

ddlcpad is short of *Doki Doki Literature Club Plus Asset Decrypter*

You can decrypt the `*.cy` file from [*Doki Doki Literature Club Plus*](https://ddlc.plus/) with this cli binary.

After decrypt, you will get a `*.out` file. Put it into [Perfare/AssetStudio](https://github.com/Perfare/AssetStudio/) to get the assert file.

## How to install đ¤

You can download a binary file on [release](https://github.com/aimerneige/DDLC-Plus-Asset-Decrypter/releases). Or you can [build](https://github.com/aimerneige/DDLC-Plus-Asset-Decrypter#how-to-build-) it yourself.

## How to use đ

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

## How to build đ¤

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

## Why do this đ§ 

There has a [project](https://github.com/MlgmXyysd/DDLC-Plus-Asset-Decrypter) do this. But I don't use php and don't want to install it on my computer. <sub>~~Actually, I don't know how to setup and use the php code. And also lazy to learn it.~~</sub>

What's more, php script for those who doesn't lean programming are very very hard to use. I wants to create a binary file which is easy to use. So that people can just download a binary file from release and use it without build or run code themselves.

## Why read file with `ioutil.ReadFile()` â

The biggest assert file is `gallery_images.cy` and with `582 MB`, There is no way to use too much memory.

PR welcome. :)

## Open Source đ

**Refer <https://github.com/MlgmXyysd/DDLC-Plus-Asset-Decrypter>**

**Cli <https://github.com/urfave/cli>**

## Support This Project đ

If you love this project, you can simply click a star or contributes codes through an issue, or create a pull request.

## More Question âī¸

Post your question on [issue](https://github.com/aimerneige/DDLC-Plus-Asset-Decrypter/issues) please.

## Contributors â¨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://aimerneige.com"><img src="https://avatars.githubusercontent.com/u/51701792?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Aimer Neige</b></sub></a><br /><a href="https://github.com/aimerneige/DDLC-Plus-Asset-Decrypter/commits?author=aimerneige" title="Code">đģ</a> <a href="https://github.com/aimerneige/DDLC-Plus-Asset-Decrypter/commits?author=aimerneige" title="Documentation">đ</a></td>
    <td align="center"><a href="https://allcontributors.org"><img src="https://avatars.githubusercontent.com/u/46410174?v=4?s=100" width="100px;" alt=""/><br /><sub><b>All Contributors</b></sub></a><br /><a href="https://github.com/aimerneige/DDLC-Plus-Asset-Decrypter/commits?author=all-contributors" title="Documentation">đ</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
