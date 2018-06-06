# Data Accelerator

[![Build Status](https://travis-ci.org/JohnGarbutt/data-acc.svg?branch=master)](https://travis-ci.org/JohnGarbutt/data-acc)
[![Go Report Card](https://goreportcard.com/badge/github.com/johngarbutt/data-acc)](https://goreportcard.com/report/github.com/johngarbutt/data-acc)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/JohnGarbutt/data-acc/internal/pkg/registry)
[![Releases](https://img.shields.io/github/release/JohnGarbutt/pfsaccel/all.svg?style=flat-square)](https://github.com/JohnGarbutt/data-acc/releases)
[![LICENSE](https://img.shields.io/github/license/JohnGarbutt/pfsaccel.svg?style=flat-square)](https://github.com/JohnGarbutt/data-acc/blob/master/LICENSE)

Data Accelerator uses commodity storage to accelerate HPC jobs.
Currently targeting initial integration with the Slurm Burst Buffer plugin,
with Lustre over Intel P4600 attached to Dell R730 with 2x100Gb/s OPA.

To see end to end demo with Slurm (not currently working):
```
cd docker-slurm
./update_burstbuffer.sh
```

To clean up after the demo:
```
docker-compose down --vol
```

To try this out, run etcd then run:
```
go get https://github.com/RSE-Cambridge/cmd/fakewarp
fakewarp --help
```

To build it locally and run tests:
```
make
make test
```