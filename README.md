# Data Sync

This is a solution for data synchronization problem between two storage systems

## Motivation

I created this project with the purpose to study some technologies that I want to learn.

## Supported Storages
- Cassandra
- ElasticSearch


## How to use

Inside datasync/config, add a json config file for your project.

In config folder already have a example.json, that used to connect to example demo

### Config: example.json
```json
{
    "App": "Example",
    "Timer": 30,
    "Storages": [
        {
            "Driver": "cassandra",
            "Host": "192.168.59.103",
            "Port": "9042",
            "Keyspace": "example",
            "Table": "post",
            "ID": "id",
            "Created": "created_datetime",
            "Updated": "updated_datetime"
        },
        {
            "Driver": "elasticsearch",
            "Host": "192.168.59.103",
            "Port": "9200",
            "Index": "example",
            "Table": "tweet",
            "ID": "mid",
            "Created": "created",
            "Updated": "updated"
        }
    ]
}
```

## Pre-requisites
  - [Fabric](http://www.fabfile.org/installing.html)
  - [Docker](https://docs.docker.com/installation/#installation)

## Getting Started

Clone this repository:

```bash
git clone git@github.com:guilhermebr/datasync.go
```
Run supported storages:

```bash
cd datasync
fab run_cassandra
fab run_elastic
```

### 1 - Using Docker for Datasync

Run datasync container:

```bash
fab run_datasync
```

Run demo:

```
fab run_example
```

Access example page in your browser: http://127.0.0.1:3000


### 2 - Run local

* Make sure you have Go environment configured.

Install godep:

```bash
go get github.com/tools/godep
```

Run datasync:

```bash
cd datasync
godep get
go run datasync.go
```

Run demo (in other shell):

```bash
cd example
go run example.go
```

## Running tests

```bash
cd datasync
go test
cd storages
go test
```

If you want to change storage host:

    -cashost = cassandra host
    -casport = cassandra port
    -caskeyspace = cassandra keyspace

    -eshost = elasticsearch host
    -esport = elasticsearch port

```bash
go test -cashost=192.168.59.103 -eshost=192.168.59.103
```
