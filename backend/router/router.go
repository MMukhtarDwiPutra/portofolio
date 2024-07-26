package router

import(
	"github.com/gorilla/mux"
	"portofolio.com/api/scmt/controller"
	"portofolio.com/service/scmt"
	"portofolio.com/repository/scmt"
	"portofolio.com/api/helper"
	"portofolio.com/api/exception"
	_ "github.com/go-sql-driver/mysql"
)

func AddRouter(muxRouter *mux.Router) *mux.Router{
	db := helper.NewDB();

	dataTmpRepository := repository.NewDataTmpSCMTRepository(db)
	dataTmpService := service.NewDataTmpService(dataTmpRepository)

	dataTmpController := controller.NewDataTmpController(dataTmpService)

	muxRouter.HandleFunc("/api/get_all_data_tmp", dataTmpController.GetAllDataTmp).Methods("GET")
	muxRouter.HandleFunc("/api/insert_data_tmp", dataTmpController.InsertDataTmp).Methods("POST")
	muxRouter.HandleFunc("/api/count_retail/{merk}", dataTmpController.CountRetailPerWitel).Methods("GET")
	muxRouter.HandleFunc("/api/count_premium/{merk}", dataTmpController.CountPremiumPerWitel).Methods("GET")
	muxRouter.HandleFunc("/api/count_stb/{merk}", dataTmpController.CountSTBPerWitel).Methods("GET")
	muxRouter.HandleFunc("/api/count_ap/{merk}", dataTmpController.CountAPPerWitel).Methods("GET")

	muxRouter.Use(exception.ErrorHandler)

	return muxRouter
}