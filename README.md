# hostat

[![Go](https://github.com/SSARCandy/hostat/workflows/Go/badge.svg?branch=master)](https://github.com/SSARCandy/hostat/actions)

> **Host** + **Status** = **hostat**

A simple CLI tool to print out host status in one line. Support [slurm](https://slurm.schedmd.com/documentation.html) status also.

```sh
$ hostat
hostname   | CPUs |     1m |     5m |    15m | memory % | disk % | UpTime | Avg Mhz |
cluster01  |    8 |    0.9 |    1.1 |    1.4 |     40 % |   19 % |    4 d | 2283.42 |
```

## Install

Fetch the latest release for your platform:

```sh
# Linux
sudo wget https://github.com/SSARCandy/hostat/releases/download/v1.0.0/hostat-linux -O /usr/local/bin/hostat
sudo chmod +x /usr/local/bin/hostat

# Windows
wget https://github.com/SSARCandy/hostat/releases/download/v1.0.0/hostat-win10.exe -O hostat.exe
.\hostat.exe
```

## Options

```sh
$ hostat --help
Usage of hostat:
  -header
        Print Header or not (default true)
  -thres_disk int
        Threshold for Disk. Render red color if >= thres (default 80)
  -thres_load int
        Threshold for Load. Render red color if >= thres (default 8)
  -thres_mem int
        Threshold for Memory. Render red color if >= thres (default 80)
```

## Fetch multiple nodes status

You can use [pdsh](https://linux.die.net/man/1/pdsh) to fetch multiple nodes status:

```sh
$ pdsh -w 'cluster[01-05]' -N -R ssh '/usr/local/bin/hostat --header=false' | sort 
cluster01  |    8 |    1.3 |    1.2 |    1.4 |     40 % |   19 % |    4 d | Avg Mhz | drain |
cluster02  |    8 |    8.0 |    8.0 |    8.0 |      8 % |   83 % |   77 d | 3900.00 |  idle | 
cluster03  |    8 |    8.0 |    8.1 |    8.0 |      7 % |   84 % |   77 d | 3900.00 | alloc | ssarcandy(8)
cluster04  |    8 |    8.1 |    8.0 |    8.0 |      7 % |   82 % |   77 d | 3900.00 | alloc | ssarcandy(8)
cluster05  |    8 |    8.2 |    8.1 |    8.1 |      7 % |   81 % |   77 d | 3900.00 | alloc | ssarcandy(8)
```
