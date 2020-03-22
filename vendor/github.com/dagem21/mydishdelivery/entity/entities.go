package entity

type Customer struct {
	Id int
	Name string
	Phone string
	Location string
	Password string
	Account int
	Imagepath string
}

type Company struct {
	Id int
	Name string
	Phone string
	Location string
	Email string
	Password string
	Account int
	Activated bool
	Rating int
	Delrange int
	Description string
	Imagepath string
}

type Employee struct {
	Id int
	Name string
	Fee int
	Phone string
	Email string
	Ability string
	Available bool
	Employedby int
	Password string
	Account int
	Rating int
	Imagepath string
	Activated bool
}

type Food struct {
	Id int
	Name string
	Company int
	Ingridient string
	Price int
	Rating int
	Imagepath string
	Discount int
	Available bool
}

type Favorite struct {
	Id int
	Foodid int
	Custid int
}

type Order struct {
	Id int
	Foodid int
	Amount int
	Location string
	Custphone string
	Custname string
	Ordertime string
	Delivered bool
}

type Employed struct {
	Id int
	Employeeid int
	Cusphone string
	Date string
	Accepted bool
}

type Manager struct {
	Id int
	Name string
	Phone string
	Password string
	Imagepath string
	Position string
}

type Comment struct {
	Id int 					`json:"id"`
	Forusertype string 		`json:"forusertype"`
	Foruserid int 			`json:"foruserid"`
	Byid int 				`json:"byid"`
	Message string 			`json:"message"`
	Date string 			`json:"placedat"`
}
