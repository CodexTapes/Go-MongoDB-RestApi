package api

type Reservation struct {
	Id         bson.ObjectId `bson:"_id"`
	Day        string        `bson:"day" json:"day"`
	NumSeats   int           `bson:"numSeats" json:"numSeats"`
	TimeSlot   string        `bson:"timeSlot" json:"timeSlot"`
	Price      int           `bson:"price" json:"price"`
	Restaurant *Restaurant   `bson:"restaurant" json:"restaurant"`
	Menu       *Menu         `bson:"menu" json:"menu"`
}

type Restaurant struct {
	Name     string `bson:"name" json:"name"`
	Location string `bson:"Location" json:"Location"`
}

type Menu struct {
	Menu_Item string `bson:"menuItem" json:"menuItem"`
}

type Reservations []Reservation
