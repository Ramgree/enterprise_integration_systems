
# HTTP endpoints (for dear BuildIT):

## Get all plants (**HTTP GET**):

Lists all the plants.

Request format:
```
URL/plants
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
        Plant_id: 5,
        Plant_type_name: "dragon",
        Plant_daily_rental_price: 5000,
        Plant_name: "forklift"
    },
    {
        Plant_id: 1,
        Plant_type_name: "rune",
        Plant_daily_rental_price: 1250,
        Plant_name: "excavator"
    }
]
```

## Estimate price (**HTTP GET**):

Estimates the price for a plant for a specified time period (start and end dates must be specified).

Date should be in year-month-day format.


Request format:
```
URL/estimate?name=<plant name>&from=<start date>&to=<end date>
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
URL/availability?name=<plant name>&from=<start date>&to=<end date>
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
