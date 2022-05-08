# Web Boiler plate

This is a Golang web boiler plate which is using web framework [Fiber](https://github.com/gofiber/fiber) and fantastic ORM [Gorm](https://gorm.io/docs/)

## Docker
First build docker image with `docker-compose`
```
docker-compose build
```
 
 then fire docker containers using following command
```
docker-compose up -d
```
 
 
## Local Setup

### Configuration file

Create a rest_task.ini configuration file and place it wherever you want. 
Configuration should look like this
```
[database]
host = localhost
name = postgres
password =postgres
```

### Export config path

put this line into ~/.bashrc file

```
export SETTINGS=path/to/config/rest_task.ini 
```

**change** ``path/to/config/`` with your config path.

### Build 
Run make command it will download all libraries and builds binary files
```
make
```

### Help

```
./bin/app --help
```

### Serve api
`serve_api` command starts server on `3000` port.

```
./bin/app serve_api
```

