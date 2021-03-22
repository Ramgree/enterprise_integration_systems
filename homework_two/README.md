## Setup 

```
go get -d -v ./...  to remove annoying VSC errors
```

## Running

```
docker build -t rentit:1.0 .
docker-compose up
```

OR better

```
python scripts/up.py
```

## Checking if we are successful individuals

Mocks set of commands that will run in Github actions

```
python scripts/autograding.py
```

## Stuff to think about
------------------------

* Plants are cached during testing, maybe do smth about it or no?
