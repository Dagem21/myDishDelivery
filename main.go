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
	dbconn, err := gorm.Open("mysql","epiz_25367355:pSaO04QFb5u@tcp(sql304.epizy.com)/epiz_25367355_mydishdelivery")

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

	router.GET("/v1/comments", CommentHandler.GetComments)
	router.GET("/v1/comments/:id", CommentHandler.GetSingleComment)
	router.PUT("/v1/comments/:id", CommentHandler.PutComment)
	router.POST("/v1/comments", CommentHandler.PostComment)
	router.DELETE("/v1/comments/:id", CommentHandler.DeleteComment)

	router.GET("/v1/companies", CompanyHandler.GetCompanies)
	router.GET("/v1/companies/:id", CompanyHandler.GetSingleCompany)
	router.PUT("/v1/companies/:id", CompanyHandler.PutCompany)
	router.POST("/v1/companies", CompanyHandler.PostCompany)
	router.DELETE("/v1/companies/:id", CompanyHandler.DeleteCompany)

	router.GET("/v1/customers", CustomerHandler.GetCustomers)
	router.GET("/v1/customers/:id", CustomerHandler.GetSingleCustomer)
	router.PUT("/v1/customers/:id", CustomerHandler.PutCustomer)
	router.POST("/v1/customers", CustomerHandler.PostCustomer)
	router.DELETE("/v1/customers/:id", CustomerHandler.DeleteCustomer)

	router.GET("/v1/employeds", EmployedHandler.GetEmployeds)
	router.GET("/v1/employeds/:id", EmployedHandler.GetSingleEmployed)
	router.PUT("/v1/employeds/:id", EmployedHandler.PutEmployed)
	router.POST("/v1/employeds", EmployedHandler.PostEmployed)
	router.DELETE("/v1/employeds/:id", EmployedHandler.DeleteEmployed)

	router.GET("/v1/employees", EmployeeHandler.GetEmployees)
	router.GET("/v1/employees/:id", EmployeeHandler.GetSingleEmployee)
	router.PUT("/v1/employees/:id", EmployeeHandler.PutEmployee)
	router.POST("/v1/employees", EmployeeHandler.PostEmployee)
	router.DELETE("/v1/employees/:id", EmployeeHandler.DeleteEmployee)

	router.GET("/v1/favorites", FavoriteHandler.GetFavorites)
	router.GET("/v1/favorites/:id", FavoriteHandler.GetSingleFavorite)
	router.PUT("/v1/favorites/:id", FavoriteHandler.PutFavorite)
	router.POST("/v1/favorites", FavoriteHandler.PostFavorite)
	router.DELETE("/v1/favorites/:id", FavoriteHandler.DeleteFavorite)

	router.GET("/v1/foods", FoodHandler.GetFoods)
	router.GET("/v1/foods/:id", FoodHandler.GetSingleFood)
	router.PUT("/v1/foods/:id", FoodHandler.PutFood)
	router.POST("/v1/foods", FoodHandler.PostFood)
	router.DELETE("/v1/foods/:id", FoodHandler.DeleteFood)

	router.GET("/v1/managers", ManagerHandler.GetManagers)
	router.GET("/v1/managers/:id", ManagerHandler.GetSingleManager)
	router.PUT("/v1/managers/:id", ManagerHandler.PutManager)
	router.POST("/v1/managers", ManagerHandler.PostManager)
	router.DELETE("/v1/managers/:id", ManagerHandler.DeleteManager)

	router.GET("/v1/orders", OrderHandler.GetOrders)
	router.GET("/v1/orders/:id", OrderHandler.GetSingleOrder)
	router.PUT("/v1/orders/:id", OrderHandler.PutOrder)
	router.POST("/v1/orders", OrderHandler.PostOrder)
	router.DELETE("/v1/orders/:id", OrderHandler.DeleteOrder)

	http.ListenAndServe(":"+os.Getenv("PORT"), router)

}