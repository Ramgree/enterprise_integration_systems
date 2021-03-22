
# HTTP endpoints (for dear BuildIT):

## Get all plants (**HTTP GET**):

Lists all the plants.

Request format:
```
URL:8080/plants
```
Example:
```
http://localhost:8080/plants
```

Returns a list of "Plant" objects in JSON format.

Example: 

```
[
    {
        id: 6,
        plantType: "adamant",
        dailyPrice: 1000,
        plantName: "mewp"
    },
    {
        id: 2,
        plantType: "adamant",
        dailyPrice: 5000,
        plantName: "bulldozer"
    }
]
```

## Estimate price (**HTTP GET**):

Estimates the price for a plant for a specified time period (start and end dates must be specified).

Date should be in year-month-day format.


Request format:
```
URL:8080/estimate?name=<plant name>&from=<start date>&to=<end date>
```

Example:
```
http://localhost:8080/estimate?name=bulldozer&from=2020-01-01&to=2020-01-10
```

Returns a JSON object with "price" (float) field.

Example:

```
{
    price: 45000
}
```

## Check availability (**HTTP GET**):

Checks if the plant is available for a specified time period (start and end dates must be specified).

Date should be in year-month-day format.

Request format:
```
URL:8080/availability?name=<plant name>&from=<start date>&to=<end date>
```

Example:
```
http://localhost:8080/availability?name=road%20roller&from=2021-10-19&to=2021-10-21
```

Returns a JSON object with "isAvailable" (boolean) field.

Example:

```
{
    isAvailable: false
}
```
