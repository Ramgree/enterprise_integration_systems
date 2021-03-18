
Just do dis:

Get all:
```
GET URL/plants

http://localhost:8080/plants
```

Estimate price:
```
GET URL/estimate?name=<plant name>&from=<start date>&to=<end date>

http://localhost:8080/estimate?name=bulldozer&from=2020-01-01&to=2020-01-10
```

Check availability:

```
GET URL/availability?name=<plant name>&from=<start date>&to=<end date>

http://localhost:8080/availability?name=road%20roller&from=2021-10-19&to=2021-10-21
```
