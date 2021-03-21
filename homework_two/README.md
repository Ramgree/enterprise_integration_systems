<<<<<<< HEAD
### Issues from last homework to **not** do again

So, whenever you kids write tests, make sure that **it works on docker-compose**

How to do that?

Just write

```shell
docker-compose -f docker-compose.yml up
```

You should see, at the end, "homework_two_grpc-test_1 exited with code 0" and "homework_two_repository-service-test_1 exited with code0"

If the exit codes are ANYTHING other than 0 then it ain't right, and thus, just like our dear Alex says, **transitively** we won't get points.

Ok?

⣠⣶⡾⠏⠉⠙⠳⢦⡀⠀⠀⠀⢠⠞⠉⠙⠲⡀⠀
⠀⠀⠀⣴⠿⠏⠀⠀⠀⠀⠀⠀⢳⡀⠀⡏⠀⠀⠀⠀⠀⢷
⠀⠀⢠⣟⣋⡀⢀⣀⣀⡀⠀⣀⡀⣧⠀⢸⠀⠀⠀⠀⠀ ⡇
⠀⠀⢸⣯⡭⠁⠸⣛⣟⠆⡴⣻⡲⣿⠀⣸⠀⠀OK⠀ ⡇
⠀⠀⣟⣿⡭⠀⠀⠀⠀⠀⢱⠀⠀⣿⠀⢹⠀⠀⠀⠀⠀ ⡇
⠀⠀⠙⢿⣯⠄⠀⠀⠀⢀⡀⠀⠀⡿⠀⠀⡇⠀⠀⠀⠀⡼
⠀⠀⠀⠀⠹⣶⠆⠀⠀⠀⠀⠀⡴⠃⠀⠀⠘⠤⣄⣠⠞⠀
⠀⠀⠀⠀⠀⢸⣷⡦⢤⡤⢤⣞⣁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢀⣤⣴⣿⣏⠁⠀⠀⠸⣏⢯⣷⣖⣦⡀⠀⠀⠀⠀⠀⠀
⢀⣾⣽⣿⣿⣿⣿⠛⢲⣶⣾⢉⡷⣿⣿⠵⣿⠀⠀⠀⠀⠀⠀
⣼⣿⠍⠉⣿⡭⠉⠙⢺⣇⣼⡏⠀⠀⠀⣄⢸⠀⠀⠀⠀⠀⠀
⣿⣿⣧⣀⣿………⣀⣰⣏⣘⣆⣀⠀⠀

peace

Setup 

```
go get -d -v ./...  to remove annoying VSC errors
```

Running

```
docker build -t rentit:1.0 .
docker-compose up
```

OR (if you feel really cool today)

```
python up.py
```

------------------------

Also, plants are cached during testing, maybe do smth about it or idk?

=======
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
python up.py
```

## Checking if we are successful individuals

Mocks set of commands that will run in Github actions

```
python autograding.py
```

## Stuff to think about
------------------------

* Plants are cached during testing, maybe do smth about it or idk?
* Availability returns true when a plant doesn't exist (based on SQL query), not good

>>>>>>> 31430d26951db2ea2737c8b5117e4f1fe740b022
