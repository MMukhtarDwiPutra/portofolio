package router

import(
	"github.com/gorilla/mux"
	"portofolio.com/api/scmt/controller"
	"portofolio.com/api/middleware"
	"portofolio.com/service/scmt"
	"portofolio.com/repository/scmt"
	"portofolio.com/api/helper"
	"portofolio.com/api/exception"
	_ "github.com/go-sql-driver/mysql"
)

func AddRouter(muxRouter *mux.Router) *mux.Router{
	db := helper.NewDB();

	dataTmpRepository := repository.NewDataTmpSCMTRepository(db)
	gudangRepository := repository.NewGudangRepository(db)
	penerimaRepository := repository.NewPenerimaRepository(db)
	fiturRepository := repository.NewFiturRepository(db)
	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepository)	
	dataTmpService := service.NewDataTmpService(dataTmpRepository, gudangRepository, penerimaRepository)
	penerimaService := service.NewPenerimaService(penerimaRepository, gudangRepository, fiturRepository)

	userController := controller.NewUserController(userService)
	dataTmpController := controller.NewDataTmpController(dataTmpService)
	penerimaController := controller.NewPenerimaController(penerimaService)

	muxRouter.HandleFunc("/api/get_all_data_tmp", dataTmpController.GetAllDataTmp).Methods("GET")
	muxRouter.HandleFunc("/api/insert_data_tmp", dataTmpController.InsertDataTmp).Methods("POST")
	muxRouter.HandleFunc("/api/count_retail/{merk}", dataTmpController.CountRetailPerWitel).Methods("GET")
	muxRouter.HandleFunc("/api/count_premium/{merk}", dataTmpController.CountPremiumPerWitel).Methods("GET")
	muxRouter.HandleFunc("/api/count_stb/{merk}", dataTmpController.CountSTBPerWitel).Methods("GET")
	muxRouter.HandleFunc("/api/count_ap/{merk}", dataTmpController.CountAPPerWitel).Methods("GET")
	muxRouter.HandleFunc("/api/get_rekap_delivery_treg", dataTmpController.GetRekapDeliveryTREG).Methods("GET")
	muxRouter.HandleFunc("/api/get_rekap_delivery_treg/witel/{lokasi_wh}", dataTmpController.GetRekapDeliveryWitel).Methods("GET")


	api := muxRouter.PathPrefix("/api").Subrouter()
	api.HandleFunc("/get_pengiriman_ont", penerimaController.GetPengirimanONT).Methods("GET")
	api.Use(middleware.JWTMiddleware)

	muxRouter.HandleFunc("/api/export_all_penerima", penerimaController.ExportAllPenerima).Methods("GET")
	muxRouter.HandleFunc("/api/download_all_sn_ont", penerimaController.DownloadAllSNONT).Methods("GET")
	muxRouter.HandleFunc("/api/download_all_sn_ont_exist", penerimaController.DownloadAllSNONTExist).Methods("GET")
	muxRouter.HandleFunc("/api/download_template_penerima", penerimaController.DownloadTemplatePenerima).Methods("GET")
	muxRouter.HandleFunc("/api/download_template_serial_number_ont", penerimaController.DownloadTemplateSerialNumberONT).Methods("GET")
	muxRouter.HandleFunc("/api/tambah_penerima", penerimaController.AddPenerima).Methods("POST")
	muxRouter.HandleFunc("/api/delete_penerima/{id}", penerimaController.DeletePenerimaById).Methods("DELETE")
	muxRouter.HandleFunc("/api/delete_all_penerima", penerimaController.DeleteAllPenerima).Methods("DELETE")
	muxRouter.HandleFunc("/api/download_serial_number/{id}", penerimaController.DownloadSerialNumber).Methods("GET")
	muxRouter.HandleFunc("/api/tambah_penerima_bulk", penerimaController.UploadPenerimaan).Methods("POST")
	muxRouter.HandleFunc("/api/edit_on_delivery_by_id/{jenis_delivery}/{id}", penerimaController.EditOnDeliveryById).Methods("PUT")

	muxRouter.HandleFunc("/api/register_user", userController.Register).Methods("POST")
	muxRouter.HandleFunc("/api/login", userController.Login).Methods("POST")
	muxRouter.HandleFunc("/api/logout", userController.Logout).Methods("GET")

	muxRouter.Use(exception.ErrorHandler)

	return muxRouter
}