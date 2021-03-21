db.createCollection("plantType")
db.createCollection("plant")
db.createCollection("booking")

manualId_1 = 1
manualId_2 = 2
manualId_3 = 3
manualId_4 = 4

db.plantType.insertMany([
	{"_id": manualId_1, "plantTypeName": "mithril"},
	{"_id": manualId_2, "plantTypeName": "adamant"},
	{"_id": manualId_3, "plantTypeName": "rune"},
	{"_id": manualId_4, "plantTypeName": "dragon"}
	]
)

db.plant.insertMany([
	{"_id": 10, "plantTypeId": manualId_3, "plantRentalDailyPrice": 1250, "plantName": "excavator"},
	{"_id": 11, "plantTypeId": manualId_2, "plantRentalDailyPrice": 5000, "plantName": "bulldozer"},
	{"_id": 12, "plantTypeId": manualId_1, "plantRentalDailyPrice": 62500, "plantName": "crane"},
	{"_id": 13, "plantTypeId": manualId_3, "plantRentalDailyPrice": 2500, "plantName": "dumper"},
	{"_id": 14, "plantTypeId": manualId_4, "plantRentalDailyPrice": 5000, "plantName": "forklift"},
	{"_id": 15, "plantTypeId": manualId_2, "plantRentalDailyPrice": 1000, "plantName": "mewp"},
	{"_id": 16, "plantTypeId": manualId_3, "plantRentalDailyPrice": 2500, "plantName": "sweeper"},
	{"_id": 17, "plantTypeId": manualId_1, "plantRentalDailyPrice": 10000, "plantName": "road roller"}
	]
)
