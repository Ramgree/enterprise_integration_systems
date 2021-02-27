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
