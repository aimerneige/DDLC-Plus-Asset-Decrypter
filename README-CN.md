# ddlcpad

## 这是什么 🤔

ddlcpad 是 *Doki Doki Literature Club Plus Asset Decrypter* 的简写。

你可以使用这个命令行工具解密来自[《心跳文学部 Plus》](https://ddlc.plus/)的 `*.cy` 资源文件。

解密成功后，你会得到一个 `*.out` 文件。之后使用 [Perfare/AssetStudio](https://github.com/Perfare/AssetStudio/) 即可得到所有的素材资源。

## 如何安装 🤗

你可以在 [release](https://github.com/aimerneige/DDLC-Plus-Asset-Decrypter/releases) 下载一个编译好的二进制文件。或者你可以 [自己编译](https://github.com/aimerneige/DDLC-Plus-Asset-Decrypter/blob/master/README-CN.md#%E5%A6%82%E4%BD%95%E7%BC%96%E8%AF%91)。

## 如何使用 😍

打开你的终端（在 Windows 上打开 powershell），执行二进制文件。

基本使用:

```bash
# On Linux
./ddlcpad <file_name>
```

```powershell
# On windows
.\ddlcpad.exe <file_name>
```

使用 `--help` 或 `-h` 查看更多的帮助信息。

## 如何编译 🤓

在你的电脑上安装和配置 [golang](https://golang.org/doc/install)，然后执行下面的指令：

```bash
# On windows
go build -o ddlcpad.exe
```

```bash
# On Linux
go build -o ddlcpad
```

如果你想在 Linux 下编译 exe 文件：

```bash
export GOARCH=amd64
export GOOS=windows
go build -o ddlcpad.exe
```

## 为什么做这个项目 🧠

实际上已经有一个[项目](https://github.com/MlgmXyysd/DDLC-Plus-Asset-Decrypter)在做这个了。但是我不用 php，同时也不想在我的电脑上装一个 php 的环境。<sub>~~实际上，我不知道如何安装和执行 php 脚本。并且也懒得学。~~</sub>

此外，php 脚本对于那些没有学过编程的人来说用起来比较困难。我想要的是一个易于使用的二进制文件。这样就可以直接在 release 下载一个二进制文件而不需要自己来编译或是运行它。

## 为什么使用 `ioutil.ReadFile()` 来读取文件 ❓

最大的资源文件是 `gallery_images.cy` ，它的大小是 `582 MB`，即使全部读取到内存里也不会占用太大的内存。

欢迎 PR。:)

## 开源相关 📖

**算法参考 <https://github.com/MlgmXyysd/DDLC-Plus-Asset-Decrypter>**

**命令行库 <https://github.com/urfave/cli>**

## 其他问题 ⁉️

欢迎提 [issue](https://github.com/aimerneige/DDLC-Plus-Asset-Decrypter/issues)。
