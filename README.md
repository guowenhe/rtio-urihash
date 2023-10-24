# urihash - urihash tool for RTIO

## 概述

`urihash`是一个用于计算RTIO URI哈希摘要的工具。

## 安装

```sh
go get github.com/guowenhe/rtio-urihash
go install github.com/guowenhe/rtio-urihash
$ rtio-urihash -h
Usage of rtio-urihash:
  -u string
        uri string, example: /uri/example
  -x    display digest with hex
```

## 示例

```sh
$ rtio-urihash -u /uri/example -x
URI: /uri/example, CRC: 0x4faf2a16
```
