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
