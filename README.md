envcon
======

A quick and dirty way to grab environment variables and place them in a config.


### Usage

```go
config = envcon.Setup("MYAPPNAME")
```

This will be looking for env variables like:

```
MYAPPNAME_DATABASE_USER
```


The variables its looking for are:

```go
type Config struct {
	Environment      Environment
	DatabasePort     string // PREFIX_DATABASE_PORT
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseHost     string
	Port             string
	Key              string
	AdminEmail       string
	AppName          string
	AppURL           string
}
```
