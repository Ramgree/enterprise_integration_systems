
# WebSocket docs (for dear ReBuildIT):

## Deployed at

```
ws://135.181.157.155:8081
```


## Setup: 

Connect to:

```
URL/ws 
```

## Get all plants:

Lists all the plants.

Send the following JSON message:
```
{
    "command": "all"
}
```

Service sends back a message with "Plant" objects in JSON format or "Error field".

Example: 

```
[
    {
        "id":8,
        "plantType":"mithril",
        "dailyPrice":10000,
        "plantName":"road roller"
    },
    {
        "id":5,
        "plantType":"dragon",
        "dailyPrice":5000,
        "plantName":"forklift"
    }
]
```

Example of response if error happened:

```
{
    "error": "Invalid query type"
}
```

## Estimate price:

Estimates the price for a plant for a specified time period (start and end dates must be specified).

Date should be in year-month-day format.


Send the following JSON message:

```
{
    "command": "estimate",
    "name": <plant name>
    "startDate": <date1> (e.g 2020-10-11)
    "endDate": <date2> (e.g 2020-10-12)
}
```

Service sends back a message with JSON object with "price" (float) field or "error" field.

Example:

```
{
    "price": 45000
}
```

Example of response if error happened:

```
{
    "error": "Failed to get estimate,  Plant not found"
}
```

## Check availability:

Checks if the plant is available for a specified time period (start and end dates must be specified).

Date should be in year-month-day format.

Send the following JSON message:

```
{
    "command": "availability",
    "name": <plant name>
    "startDate": <date1> (e.g 2020-10-11)
    "endDate": <date2> (e.g 2020-10-12)
}
```

Service sends back a message with JSON object with "isAvailable" (bool) field or "error" field.

Example:

```
{
    "isAvailable": false
}
```

Example of response if error happened:

```
{
    "error": "Query has missing fields"
}
```
