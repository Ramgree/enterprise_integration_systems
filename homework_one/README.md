## Homework 1

The whole back end is done, and the core functionality is 100% tested.

### Task 1

For task 1 you can go to backend/src and then 

``` sh

go test -run=''

```

In order to test **core** functionality, CRUD.

Our todo list is slightly different than usual, in the sense that it's __meant__ to be a **TASK DAG**. That's our attempt at the "crazy/novel" feature.

It does **not** check the DAG's acyclicity, it's **on you** to not bug it out.

If a **TASK DAG** is hard to visualize for you, just look at this picture, you'll get it:

![DAG](https://i.stack.imgur.com/e0NQk.png)

### Task 2

No status

### Task 3

* I guess we've done all functional tests?

* Dockerized the back end service

* Made the docker compose, but no tests yet

* No integration tests done

### Task 4

* Docker registry set up

The docker registry was set up on a google cloud compute engine.

We did it exactly as on the docker website, not sure how to prove it :D ![here](https://docs.docker.com/registry/)

```sh

docker run -d -p 5000:5000 --name registry registry:2

```

Then we cloned our repo, and:

```sh

docker image tag hw1-backend:1.0 localhost:5000/hw1-backend

```

then tested by pulling it:

```sh

docker pull localhost:5000/hw1-backend

```

* Backend image pushed

* We just used a alpine linux image to reduce the size. The original debian one was around ~ 700 mb, the alpine linux hovers around 5 mb, hence we achieved **WHOPPING** 140x smaller image.

### Task 5


