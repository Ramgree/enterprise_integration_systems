## Homework 1

The whole back end is done, and the core functionality is 100% tested.

### Task 1 - Complete

For task 1 you can go to backend/src and then 

``` sh

go test -run=''

```

In order to test **core** functionality, CRUD.

Our todo list is slightly different than usual, in the sense that it's __meant__ to be a **TASK DAG**. That's our attempt at the "crazy/novel" feature.

It does **not** check the DAG's acyclicity, it's **on you** to not bug it out.

If a **TASK DAG** is hard to visualize for you, just look at this picture, you'll get it:

![DAG](https://i.stack.imgur.com/e0NQk.png)

### Task 2 - Complete

In order to use the CLI, cd to /frontend folder(make sure that todocli is executable) and just 

``` sh
todocli <command>
```

the commands are:

``` sh
1) todocli get-all
2) todocli todo <title> [...dependencies], e.g todocli todo "Insert your task here!" 1 2 3
3) todo remove <id> 
4) todocli check <id> <do|undo>
```

Running tests:

``` sh
cd frontend/test_non_docker
go test -run=''
```

Before testing, make sure that the back-end is up.

You can run the docker file as follows, after building it:

``` sh
docker run --network host hw1-frontend:1.0 <command>
```


### Task 3 - Complete

* All CRUD functional tests are done

* Dockerized the back end service

* Made the docker compose

Just write

``` sh
docker-compose -f docker-compose.yml up -d \
docker exec -it <id of the front-end container> todocli get-all
```

* All CRUD integration tests are done.

### Task 4 - Complete

* Docker registry set up

The docker registry was set up on a google cloud compute engine.

We did it exactly as on the docker website, not sure how to prove it :D ![here](https://docs.docker.com/registry/)

```sh

docker run -d -p 5000:5000 --name registry registry:2

```

Then we cloned our repo, built our images, and:

```sh

docker image tag hw1-backend:1.0 localhost:5000/hw1-backend
docker image tag hw1-frontend:1.0 localhost:5000/hw1-frontend

```

then tested by pulling it:

```sh

docker pull localhost:5000/hw1-backend
docker pull localhost:5000/hw1-frontend

```

And then ran everything using the `docker-compose-local-registry.yml`

* We just used a alpine linux image to reduce the size. The original debian one was around ~ 700 mb, the alpine linux hovers around 5 mb, hence we achieved **WHOPPING** 140x smaller image.

### Task 5 - Complete

Done, the compose file is `docker-compose-task-five.yml`

It will work as long as you have built our backend with the tag `hw1-backend:1.0`

We have put Grafana as the extra service under /mystery
