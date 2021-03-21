db.createCollection("plant")

db.plant.insertMany([
	{"_id": 5, "plantTypeName": "dragon", "plantRentalDailyPrice": 5000, "plantName": "forklift"},
	{"_id": 6, "plantTypeName": "adamant", "plantRentalDailyPrice": 1000, "plantName": "mewp"},
	{"_id": 7, "plantTypeName": "rune", "plantRentalDailyPrice": 2500, "plantName": "sweeper"},
	{"_id": 8, "plantTypeName": "mithril", "plantRentalDailyPrice": 10000, "plantName": "road roller"}
	]
)
