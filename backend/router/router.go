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

	gudangService := service.NewGudangService(gudangRepository)
	userService := service.NewUserService(userRepository)	
	dataTmpService := service.NewDataTmpService(dataTmpRepository, gudangRepository, penerimaRepository)
	penerimaService := service.NewPenerimaService(penerimaRepository, gudangRepository, fiturRepository)

	gudangController := controller.NewGudangController(gudangService)
	userController := controller.NewUserController(userService)
	dataTmpController := controller.NewDataTmpController(dataTmpService)
	penerimaController := controller.NewPenerimaController(penerimaService)

	api := muxRouter.PathPrefix("/api").Subrouter()

	api.HandleFunc("/user", userController.GetUser).Methods("GET")

	api.HandleFunc("/upload_gudang", gudangController.UploadGudangBulk).Methods("POST")

	api.HandleFunc("/upload_data_tmp", dataTmpController.UploadDataTmp).Methods("POST")
	api.HandleFunc("/get_all_data_tmp", dataTmpController.GetAllDataTmp).Methods("GET")
	api.HandleFunc("/insert_data_tmp", dataTmpController.InsertDataTmp).Methods("POST")
	api.HandleFunc("/count_retail/{merk}", dataTmpController.CountRetailPerWitel).Methods("GET")
	api.HandleFunc("/count_premium/{merk}", dataTmpController.CountPremiumPerWitel).Methods("GET")
	api.HandleFunc("/count_stb/{merk}", dataTmpController.CountSTBPerWitel).Methods("GET")
	api.HandleFunc("/count_ap/{merk}", dataTmpController.CountAPPerWitel).Methods("GET")
	muxRouter.HandleFunc("/api/get_rekap_delivery_treg", dataTmpController.GetRekapDeliveryTREG).Methods("GET")
	api.HandleFunc("/get_rekap_delivery_treg/witel/{lokasi_wh}", dataTmpController.GetRekapDeliveryWitel).Methods("GET")
	api.HandleFunc("/export_data_tmp", dataTmpController.ExportDataTmp).Methods("GET")
	api.HandleFunc("/export_minimum_stock_database", dataTmpController.ExportMinimumStockDatabase).Methods("GET")
	api.HandleFunc("/download_template_minimum_stock", dataTmpController.DownloadTemplateMinimumStock).Methods("GET")
	api.HandleFunc("/download_template_data_tmp", dataTmpController.DownloadTemplateDataTmp).Methods("GET")
	api.HandleFunc("/export_data_tmp_rekap_page/{jenis_warna}/{jenis_export}", dataTmpController.ExportDataTmpRekapPage).Methods("GET")

	muxRouter.HandleFunc("/api/get_pengiriman_ont", penerimaController.GetPengirimanONT).Methods("GET")
	api.HandleFunc("/export_all_penerima", penerimaController.ExportAllPenerima).Methods("GET")
	api.HandleFunc("/export_all_penerima_ont", penerimaController.ExportAllPenerimaONT).Methods("GET")
	api.HandleFunc("/download_all_sn_ont", penerimaController.DownloadAllSNONT).Methods("GET")
	api.HandleFunc("/download_all_sn_ont_exist", penerimaController.DownloadAllSNONTExist).Methods("GET")
	api.HandleFunc("/download_template_penerima", penerimaController.DownloadTemplatePenerima).Methods("GET")
	api.HandleFunc("/download_template_serial_number_ont", penerimaController.DownloadTemplateSerialNumberONT).Methods("GET")
	api.HandleFunc("/tambah_penerima", penerimaController.AddPenerima).Methods("POST")
	api.HandleFunc("/delete_penerima/{id}", penerimaController.DeletePenerimaById).Methods("DELETE")
	api.HandleFunc("/delete_all_penerima", penerimaController.DeleteAllPenerimaONT).Methods("DELETE")
	api.HandleFunc("/download_serial_number/{id}", penerimaController.DownloadSerialNumber).Methods("GET")
	api.HandleFunc("/tambah_penerima_bulk/{jenis_upload}", penerimaController.UploadPenerimaan).Methods("POST")
	api.HandleFunc("/edit_on_delivery_by_id/{jenis_delivery}/{id}", penerimaController.EditOnDeliveryById).Methods("PUT")

	muxRouter.HandleFunc("/scmt/register_user", userController.Register).Methods("POST")
	muxRouter.HandleFunc("/scmt/login", userController.Login).Methods("POST")
	muxRouter.HandleFunc("/scmt/logout", userController.Logout).Methods("GET")
	api.HandleFunc("/scmt/change_password/{id}", userController.ChangePassword).Methods("PUT")
	api.HandleFunc("/scmt/change_data_user/{id}", userController.ChangeDataUser).Methods("PUT")

	api.Use(middleware.JWTMiddleware)

	muxRouter.Use(exception.ErrorHandler)

	return muxRouter
}