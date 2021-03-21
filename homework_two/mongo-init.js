db.createCollection("plantType")
db.createCollection("plant")
db.createCollection("booking")

manualId_1 = ObjectId()
manualId_2 = ObjectId()
manualId_3 = ObjectId()
manualId_4 = ObjectId()

db.plantType.insertMany([
	{"_id": manualId_1, "plantTypeName": "mithril"},
	{"_id": manualId_2, "plantTypeName": "adamant"},
	{"_id": manualId_3, "plantTypeName": "rune"},
	{"_id": manualId_4, "plantTypeName": "dragon"}
	]
)

db.plant.insertMany([
	{"plantTypeId": manualId_3, "plantRentalDailyPrice": 1250, "plantName": "excavator"},
	{"plantTypeId": manualId_2, "plantRentalDailyPrice": 5000, "plantName": "bulldozer"},
	{"plantTypeId": manualId_1, "plantRentalDailyPrice": 62500, "plantName": "crane"},
	{"plantTypeId": manualId_3, "plantRentalDailyPrice": 2500, "plantName": "dumper"},
	{"plantTypeId": manualId_4, "plantRentalDailyPrice": 5000, "plantName": "forklift"},
	{"plantTypeId": manualId_2, "plantRentalDailyPrice": 1000, "plantName": "mewp"},
	{"plantTypeId": manualId_3, "plantRentalDailyPrice": 2500, "plantName": "sweeper"},
	{"plantTypeId": manualId_1, "plantRentalDailyPrice": 10000, "plantName": "road roller"}
	]
)
