package main

import (
	"net/http"
	"os"

	commentrepo "github.com/dagem21/mydishdelivery/comments/repository"
	commentserv "github.com/dagem21/mydishdelivery/comments/service"

	companyrepo "github.com/dagem21/mydishdelivery/companies/repository"
	companyserv"github.com/dagem21/mydishdelivery/companies/service"

	customerrepo "github.com/dagem21/mydishdelivery/customers/repository"
	customerserv "github.com/dagem21/mydishdelivery/customers/service"

	employedrepo "github.com/dagem21/mydishdelivery/employed/repository"
	employedserv "github.com/dagem21/mydishdelivery/employed/service"

	employeerepo "github.com/dagem21/mydishdelivery/employees/repository"
	employeeserv "github.com/dagem21/mydishdelivery/employees/service"

	favoriterepo "github.com/dagem21/mydishdelivery/favorites/repository"
	favoriteserv "github.com/dagem21/mydishdelivery/favorites/service"

	foodrepo "github.com/dagem21/mydishdelivery/foods/repository"
	foodserv "github.com/dagem21/mydishdelivery/foods/service"

	managerrepo "github.com/dagem21/mydishdelivery/managers/repository"
	managerserv "github.com/dagem21/mydishdelivery/managers/service"

	orderrepo "github.com/dagem21/mydishdelivery/orders/repository"
	orderserv "github.com/dagem21/mydishdelivery/orders/service"

	"github.com/dagem21/mydishdelivery/delivery/http/handler"
	"github.com/julienschmidt/httprouter"


	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" 

)

func main()  {
	dbconn, err := gorm.Open("mysql","x8CvpP7EjM:U908Nz50fv@tcp(remotemysql.com:3306)/x8CvpP7EjM")

	if err!=nil {
		panic(err.Error())
	}

	defer dbconn.Close()

	commentRepo := commentrepo.NewCommentGormRepo(dbconn)
	commentSrv := commentserv.NewCommentService(commentRepo)
	CommentHandler := handler.NewCommentHandler(commentSrv)

	companyRepo := companyrepo.NewCompanyGormRepo(dbconn)
	companySrv := companyserv.NewCompanyService(companyRepo)
	CompanyHandler := handler.NewCompanyHandler(companySrv)

	customerRepo := customerrepo.NewCustomerGormRepo(dbconn)
	customerSrv := customerserv.NewCustomerService(customerRepo)
	CustomerHandler := handler.NewCustomerHandler(customerSrv)

	employedRepo := employedrepo.NewEmployedGormRepo(dbconn)
	employedSrv := employedserv.NewEmployedService(employedRepo)
	EmployedHandler := handler.NewEmployedHandler(employedSrv)

	employeeRepo := employeerepo.NewEmployeeGormRepo(dbconn)
	employeeSrv := employeeserv.NewEmployeeService(employeeRepo)
	EmployeeHandler := handler.NewEmployeeHandler(employeeSrv)

	favoriteRepo := favoriterepo.NewFavoriteGormRepo(dbconn)
	favoriteSrv := favoriteserv.NewFavoriteService(favoriteRepo)
	FavoriteHandler := handler.NewFavoriteHandler(favoriteSrv)

	foodRepo := foodrepo.NewFoodGormRepo(dbconn)
	foodSrv := foodserv.NewFoodService(foodRepo)
	FoodHandler := handler.NewFoodHandler(foodSrv)

	managerRepo := managerrepo.NewManagerGormRepo(dbconn)
	managerSrv := managerserv.NewManagerService(managerRepo)
	ManagerHandler := handler.NewManagerHandler(managerSrv)

	orderRepo := orderrepo.NewOrderGormRepo(dbconn)
	orderSrv := orderserv.NewOrderService(orderRepo)
	OrderHandler := handler.NewOrderHandler(orderSrv)

	router := httprouter.New()

	router.GET("/", CommentHandler.GetComments)
	router.GET("/comments", CommentHandler.GetComments)
	router.GET("/comments/:id", CommentHandler.GetSingleComment)
	router.PUT("/comments/:id", CommentHandler.PutComment)
	router.POST("/comments", CommentHandler.PostComment)
	router.DELETE("/comments/:id", CommentHandler.DeleteComment)

	router.GET("/companies", CompanyHandler.GetCompanies)
	router.GET("/companies/:id", CompanyHandler.GetSingleCompany)
	router.PUT("/companies/:id", CompanyHandler.PutCompany)
	router.POST("/companies", CompanyHandler.PostCompany)
	router.DELETE("/companies/:id", CompanyHandler.DeleteCompany)

	router.GET("/customers", CustomerHandler.GetCustomers)
	router.GET("/customers/:id", CustomerHandler.GetSingleCustomer)
	router.PUT("/customers/:id", CustomerHandler.PutCustomer)
	router.POST("/customers", CustomerHandler.PostCustomer)
	router.DELETE("/customers/:id", CustomerHandler.DeleteCustomer)

	router.GET("/employeds", EmployedHandler.GetEmployeds)
	router.GET("/employeds/:id", EmployedHandler.GetSingleEmployed)
	router.PUT("/employeds/:id", EmployedHandler.PutEmployed)
	router.POST("/employeds", EmployedHandler.PostEmployed)
	router.DELETE("/employeds/:id", EmployedHandler.DeleteEmployed)

	router.GET("/employees", EmployeeHandler.GetEmployees)
	router.GET("/employees/:id", EmployeeHandler.GetSingleEmployee)
	router.PUT("/employees/:id", EmployeeHandler.PutEmployee)
	router.POST("/employees", EmployeeHandler.PostEmployee)
	router.DELETE("/employees/:id", EmployeeHandler.DeleteEmployee)

	router.GET("/favorites", FavoriteHandler.GetFavorites)
	router.GET("/favorites/:id", FavoriteHandler.GetSingleFavorite)
	router.PUT("/favorites/:id", FavoriteHandler.PutFavorite)
	router.POST("/favorites", FavoriteHandler.PostFavorite)
	router.DELETE("/favorites/:id", FavoriteHandler.DeleteFavorite)

	router.GET("/foods", FoodHandler.GetFoods)
	router.GET("/foods/:id", FoodHandler.GetSingleFood)
	router.PUT("/foods/:id", FoodHandler.PutFood)
	router.POST("/foods", FoodHandler.PostFood)
	router.DELETE("/foods/:id", FoodHandler.DeleteFood)

	router.GET("/managers", ManagerHandler.GetManagers)
	router.GET("/managers/:id", ManagerHandler.GetSingleManager)
	router.PUT("/managers/:id", ManagerHandler.PutManager)
	router.POST("/managers", ManagerHandler.PostManager)
	router.DELETE("/managers/:id", ManagerHandler.DeleteManager)

	router.GET("/orders", OrderHandler.GetOrders)
	router.GET("/orders/:id", OrderHandler.GetSingleOrder)
	router.PUT("/orders/:id", OrderHandler.PutOrder)
	router.POST("/orders", OrderHandler.PostOrder)
	router.DELETE("/orders/:id", OrderHandler.DeleteOrder)
	
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
	
}