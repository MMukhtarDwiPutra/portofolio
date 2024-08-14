import Sidebar from "./Components/Sidebar"
import Navbar from "./Components/Navbar"
import "../../Assets/js/Sidebar_SCMT.js"
import "../../Assets/css/HomeSCMT.css"
import 'bootstrap/dist/css/bootstrap.min.css';  
import React, {Component, useEffect, StrictMode, useState  } from 'react'
import { BrowserRouter as Router, Route, Routes, Link, useNavigate } from "react-router-dom";
import { useParams } from 'react-router-dom';

const HomeSCMT = () => {
    const { lokasi_wh } = useParams();
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [history, setHistory] = useState([]);
    const [login, setLogin] = useState(false)
    const [waktuUpdate, setWaktuUpdate] = useState('')
    const fetchDataUser = async () =>{
        try{
            let response;
            response = await fetch(`http://localhost:8080/api/user`, { 
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                })

            const result = await response.json();
            if(result.data.username){
                setLogin(true)
            }
        }catch(error){
        }
    }
    useEffect(() => {
        fetchDataUser();
    }, []);

    const navigate = useNavigate();

    const fetchData = async () => {
        try {
            let response;
            if (!lokasi_wh) {
                response = await fetch(`http://localhost:8080/api/get_rekap_delivery_treg`, {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });
            } else {
                response = await fetch(`http://localhost:8080/api/get_rekap_delivery_treg/witel/${lokasi_wh}`,{ 
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });
            }

            if (response.ok) {
                const result = await response.json();
                if(!lokasi_wh){
                    setData(result.data.treg);
                }else{
                    setData(result.data.response);
                }
                setWaktuUpdate(result.data.last_update)
            }
        } catch (error) {
            setError(error.message);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchData();
    }, [lokasi_wh]);

    useEffect(() => {
        if (lokasi_wh) {
            setHistory((prevHistory) => {
                const newHistory = [...prevHistory];
                // Only add the location if it's not already the latest entry
                if (newHistory[newHistory.length - 1] !== lokasi_wh) {
                    newHistory.push(lokasi_wh);
                }
                return newHistory;
            });
        }
    }, [lokasi_wh]);

    useEffect(() => {
    }, [history])

    const handleLinkClick = (event, locationWh) => {
        event.preventDefault();
        navigate(`/scmt/rekap_delivery/witel/${locationWh}`);
    };

    const handleLinkBackClick = (event) => {
        event.preventDefault();
        if (history.length > 1) {
            // Remove the current location from history
            const newHistory = [...history];
            newHistory.pop();
            const previousLocation = newHistory[newHistory.length - 1];
            setHistory(newHistory);
            navigate(`/scmt/rekap_delivery/witel/${previousLocation}`);
        } else {
            setHistory([])
            navigate(`/scmt/rekap_delivery/`);
        }
    };

    const [exportData, setExportData] = useState('all')
    const handleExportChange = async (e) =>{
        setExportData(e.target.value)
    }

    const [exportDataAll, setExportDataAll] = useState('all')
    const handleExportAllChange = async (e) =>{
        setExportDataAll(e.target.value)
    }

    const handleOnClickExport = async () => {
        if (!lokasi_wh) {
            const link = document.createElement('a');
            link.href = `http://localhost:8080/api/export_data_tmp_rekap_page/${exportData}/treg_only`;
            document.body.appendChild(link);
            link.click();
            link.remove();
        } else {
            const link = document.createElement('a');
            link.href = `http://localhost:8080/api/export_data_tmp_rekap_page/${exportData}/${lokasi_wh}`;
            console.log(`http://localhost:8080/api/export_data_tmp_rekap_page/${exportData}/${lokasi_wh}`)
            document.body.appendChild(link);
            link.click();
            link.remove();
        }
    }

    const handleOnClickExportAll = async () => {
        // await fetch(`http://localhost:8080/api/export_data_tmp_rekap_page/${exportDataAll}/all`, {
        //     method:'GET'
        // });
        // navigate(`http://localhost:8080/api/export_data_tmp_rekap_page/${exportDataAll}/all`)
        try {
            // const response = await fetch(`http://localhost:8080/api/export_data_tmp_rekap_page/${exportDataAll}/all`, {
            //     method: 'GET',
            //     headers: {
            //         'Content-Type': 'application/json',
            //     },
            // });

            const link = document.createElement('a');
            link.href = `http://localhost:8080/api/export_data_tmp_rekap_page/${exportDataAll}/all`;
            // link.setAttribute('download', 'data_export.xlsx'); // Specify the file name
            document.body.appendChild(link);
            link.click();
            link.remove();
        } catch (error) {
            console.error('Error during export:', error);
        }
    }

    useEffect(() => {
        console.log(exportDataAll)
    }, [exportDataAll])

    let grandTotalRetailStock = 0;
    let grandTotalPremiumStock = 0;
    let grandTotalRetailGapStock = 0;
    let grandTotalPremiumGapStock = 0;
    let grandTotalRetailNeed = 0;
    let grandTotalPremiumNeed = 0;
    let grandTotalRetail = 0;
    let grandTotalPremium = 0;
    let grandTotalOnDeliveryRetail = 0;
    let grandTotalOnDeliveryPremium = 0;

    const formatter = new Intl.NumberFormat('id-ID');

	return(
		<>
		<div className="wrapper d-flex align-items-stretch">
			<Sidebar/>
	        <div id="content" style={{margin: "0 auto", boxSizing: "border-box"}}>
	            <div className="container-fluid" style={{width: "105.5%"}}>
	            	<Navbar/>
	                <div className="card mb-3 mt-3">
	                    <div className="card-body mb-2">
	                            <div className="container-fluid p-0" style={{width: "100%"}}>
	                                <div className="export-filter mb-1 row">
	                                    <div className="col-md-6 order-md-1 order-2">
	                                        
	                                    </div>
	                                    <div className="col-md-6 order-md-2 order-1 d-flex justify-content-md-end">
	                                        
	                                    </div>
	                                </div>
	                            </div>

                                {login && (
                                <div className="container-fluid p-0" style={{width: "100%"}}>
                                    <div className="export-filter mb-1 row">
                                        <div className="col-md-6 order-md-1 order-2">
                                            <form className="mt-3" id="form_export">
                                                <div className="row">
                                                    <div className="col-4">
                                                        <select id="exportFilter" onChange={handleExportChange} className="form-control" style={{marginLeft:"0"}}>
                                                            <option value="all">All</option>
                                                            <option value="merah">Merah</option>
                                                            <option value="kuning">Kuning</option>
                                                        </select>
                                                    </div>
                                                    <div className="col-6" style={{marginLeft:"-20px"}}>
                                                        <button onClick={handleOnClickExport} className="btn btn-secondary mb-2 col-6" value="false" type="button">Export Data</button>
                                                    </div>
                                                </div>
                                            </form>

                                            <span className="ml-1">{waktuUpdate && (`Last update: ${waktuUpdate}`)} </span>
                                        </div>
                                        <div className="col-md-6 order-1 d-flex justify-content-end">
                                            <form className="mt-3" id="form_export_all">
                                                <div className="row">
                                                    <div className="col-6">
                                                        <select id="exportFilterAll" onChange={handleExportAllChange} className="form-control">
                                                            <option value="all">All</option>
                                                            <option value="merah">Merah</option>
                                                            <option value="kuning">Kuning</option>
                                                        </select>
                                                    </div>
                                                    <div className="col-6">
                                                        <button onClick={handleOnClickExportAll} className="btn btn-secondary mb-2" value="false" type="button">Export All Data</button>
                                                    </div>
                                                </div>
                                            </form>
                                        </div>
                                    </div>
                                </div>
                                )}


	                            <div className="category-filter">
	                            <div className="table-responsive">
	                                <div className="category-filter mb-3">
                                        {history.length != 0 && (<a className="btn btn-secondary mb-2 ml-1" onClick={handleLinkBackClick}>Back</a>)}
                                      	<div id="tableBiasa">
                                            <table className="table table-bordered" id="filterTable" width="100%" style={{fontSize:"14px"}}>
                                                <thead className="thead-grey">
                                                	<tr className="text-center" style={{verticalAlign: "middle"}}>
                                                        <th rowSpan="2" className="" style={{minWidth: "250px", textAlign: "center"}}>
                                                            Warehouse
                                                        </th>
                                                        <th rowSpan="2" hidden>regional</th>
                                                        <th rowSpan="2" hidden>witel</th>
                                                        <th rowSpan="2"
                                                            style={{textAlign: "center", verticalAlign: "middle"}}
                                                            style={{maxWidth:"70px"}} hidden>Minimum Qty</th>
                                                        <th colSpan="2">Stock SCMT</th>
                                                        <th colSpan="2">GAP Stock</th>
                                                        <th colSpan="2">Kebutuhan</th>
                                                        <th colSpan="2">Minimum Stock Requirement</th>
                                                        <th colSpan="2">On Delivery</th>
                                                    </tr>
                                                    <tr className="text-center">
                                                    	<th style={{textAlign: "center", verticalAlign: "middle"}}>Total
                                                            Retail</th>
                                                        <th style={{textAlign: "center", verticalAlign: "middle"}}>Total
                                                            Premium</th>
                                                        <th
                                                            style={{maxWidth: "100px", textAlign: "center", verticalAlign: "middle"}}>
                                                            Total Retail</th>
                                                        <th
                                                            style={{maxWidth: "100px", textAlign: "center", verticalAlign: "middle"}}>
                                                            Total Premium</th>

                                                        <th
                                                            style={{maxWidth: "100px", textAlign: "center", verticalAlign: "middle"}}>
                                                            Total Retail</th>
                                                        <th
                                                            style={{maxWidth: "100px", textAlign: "center", verticalAlign: "middle"}}>
                                                            Total Premium</th>

                                                        <th
                                                            style={{maxWidth: "100px", textAlign: "center", verticalAlign: "middle"}}>
                                                            Total Retail</th>
                                                        <th
                                                            style={{maxWidth: "100px", textAlign: "center", verticalAlign: "middle"}}>
                                                            Total Premium</th>

                                                        <th
                                                            style={{maxWidth: "100px", textAlign: "center", verticalAlign: "middle"}}>
                                                            Total Retail</th>
                                                        <th style={{maxWidth: "100px", textAlign: "center", verticalAlign: "middle"}}>
                                                            Total Premium</th>
                                                    </tr>
                                                 </thead>
                                                <tbody>
                                                {Array.isArray(data) && 
                                                    data.map((item, index) => {
                                                        grandTotalRetailStock += item.total_retail_stock;
                                                        grandTotalPremiumStock += item.total_premium_stock;
                                                        grandTotalRetailGapStock += item.total_retail_stock - item.total_retail + item.on_delivery_total_retail;
                                                        grandTotalPremiumGapStock += item.total_premium_stock - item.total_premium + item.on_delivery_total_premium;
                                                        grandTotalRetail += item.total_retail;
                                                        grandTotalPremium += item.total_premium;
                                                        grandTotalOnDeliveryRetail += item.on_delivery_total_retail;
                                                        grandTotalOnDeliveryPremium += item.on_delivery_total_premium;
                                                        if(item.qty_kirim_retail_zte + item.qty_kirim_retail_fh + item.qty_kirim_retail_hw + item.qty_kirim_retail_alu != 0){
                                                            grandTotalRetailNeed += item.qty_kirim_retail_zte + item.qty_kirim_retail_fh + item.qty_kirim_retail_hw + item.qty_kirim_retail_alu
                                                        }
                                                        if(item.qty_kirim_premium_zte + item.qty_kirim_premium_fh + item.qty_kirim_premium_hw != 0){
                                                            grandTotalPremiumNeed += item.qty_kirim_premium_zte + item.qty_kirim_premium_fh + item.qty_kirim_premium_hw
                                                        }

                                                        let classGapRetailTag = '';

                                                        if(item.total_retail_stock - item.total_retail + item.on_delivery_total_retail < -(item.total_retail * 0.75)){
                                                            classGapRetailTag = 'td-red'
                                                        }else if(item.total_retail_stock - item.total_retail + item.on_delivery_total_retail < 0){
                                                            classGapRetailTag ='td-yellow'
                                                        }else{
                                                            classGapRetailTag = 'td-green'
                                                        }

                                                        let classGapPremiumTag = '';
                                                        if(item.total_premium_stock - item.total_premium + item.on_delivery_total_premium < -(item.total_premium * 0.75)){
                                                            classGapPremiumTag = 'td-red'
                                                        }else if(item.total_premium_stock - item.total_premium + item.on_delivery_total_premium < 0){
                                                            classGapPremiumTag = 'td-yellow'
                                                        }else{
                                                            classGapPremiumTag = 'td-green'
                                                        }

                                                        var clickWH = false
                                                        if(item.lokasi_wh.includes("TREG") || item.lokasi_wh.includes("WITEL")){
                                                            clickWH = true
                                                        }
                                                        return(
                                                            <tr key={index} className="font-grey">
                                                                <td className="thead-grey bolding-font">
                                                                    {clickWH ? (
                                                                    <a
                                                                      href=""
                                                                      onClick={(e) => handleLinkClick(e, item.lokasi_wh)}
                                                                      className="nav-link active">{item.lokasi_wh}
                                                                    </a>
                                                                    ) : (
                                                                        item.lokasi_wh
                                                                    )}
                                                                </td>
                                                                <td>{formatter.format(item.total_retail_stock)}</td>
                                                                <td>{formatter.format(item.total_premium_stock)}</td>
                                                                <td className={classGapRetailTag}>{formatter.format(item.total_retail_stock - item.total_retail + item.on_delivery_total_retail)}</td>
                                                                <td className={classGapPremiumTag}>{formatter.format(item.total_premium_stock - item.total_premium + item.on_delivery_total_premium)}</td>
                                                                <td>{item.qty_kirim_retail_zte + item.qty_kirim_retail_fh + item.qty_kirim_retail_hw + item.qty_kirim_retail_alu != 0 && (formatter.format(item.qty_kirim_retail_zte + item.qty_kirim_retail_fh + item.qty_kirim_retail_hw + item.qty_kirim_retail_alu))}
                                                                </td>
                                                                <td>{item.qty_kirim_premium_zte + item.qty_kirim_premium_fh + item.qty_kirim_premium_hw != 0 && (formatter.format(item.qty_kirim_premium_zte + item.qty_kirim_premium_fh + item.qty_kirim_premium_hw))}
                                                                </td>
                                                                <td>{formatter.format(item.total_retail)}</td>
                                                                <td>{formatter.format(item.total_premium)}</td>
                                                                <td>{formatter.format(item.on_delivery_total_retail)}</td>
                                                                <td>{formatter.format(item.on_delivery_total_premium)}</td>
                                                            </tr>
                                                        );
                                                })}
                                                <tr className="thead-grey bolding-font">
                                                <td className="text-center"> Total </td>
                                                <td>{formatter.format(grandTotalRetailStock)}</td>
                                                <td>{formatter.format(grandTotalPremiumStock)}</td>
                                                <td>{formatter.format(grandTotalRetailGapStock)}</td>
                                                <td>{formatter.format(grandTotalPremiumGapStock)}</td>
                                                <td>{formatter.format(grandTotalRetailNeed)}</td>
                                                <td>{formatter.format(grandTotalPremiumNeed)}</td>
                                                <td>{formatter.format(grandTotalRetail)}</td>
                                                <td>{formatter.format(grandTotalPremium)}</td>
                                                <td>{formatter.format(grandTotalOnDeliveryRetail)}</td>
                                                <td>{formatter.format(grandTotalOnDeliveryPremium)}</td>
                                                </tr>
                                                </tbody>
                                            </table>
                                        </div>
                                    </div>
                                </div>
                                </div>
	                    </div>
	                </div>
	            </div>
	        </div>
	    </div>
	   	</>
	)
}

export default HomeSCMT;