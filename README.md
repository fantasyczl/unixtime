# A Tool for Unix timestamp


## Install

Install by go install:

```bash
go install github.com/fantasyczl/unixtime@latest
```

## Usage

```bash
$ unixtime -date="2023-10-01 12:00:00"

time:
        2023-10-01 12:00:00 +0800 CST
unix:
        1696132800
```

Show current time and unix timestamp:

```bash
$ unixtime -date=now

now:
2025-05-27 23:02:07.84816 +0800 CST m=+0.000192751
ts: 1748358127
```


