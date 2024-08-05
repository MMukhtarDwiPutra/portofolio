import Sidebar from "./Components/Sidebar"
import Navbar from "./Components/Navbar"
import "../../Assets/js/Sidebar_SCMT.js"
import "../../Assets/css/HomeSCMT.css"
import 'bootstrap/dist/css/bootstrap.min.css';  
import React, {Component, useEffect, StrictMode, useState  } from 'react'
import { BrowserRouter as Router, Route, Routes, Link, useNavigate } from "react-router-dom";
import { useParams } from 'react-router-dom';

const HomeSCMT = () => {
    // const { lokasi_wh } = useParams();

    // const [data, setData] = useState([]);
    // const [loading, setLoading] = useState(true);
    // const [error, setError] = useState(null);

    // const fetchData = async () => {
    //     // Simulate an API call
    //     try {
    //         // console.log("HIT Get Fetch Data Penerima")
    //         let response
    //         if(lokasi_wh === undefined){
    //             const response = await fetch(`http://localhost:8080/api/get_rekap_delivery_treg`); // Replace with your API endpoint
    //             const result = await response.json();
    //             setData(result.data)
    //         }else{
    //             if(lokasi_wh.includes("TREG")){
    //                 const response = await fetch(`http://localhost:8080/api/get_rekap_delivery_treg/witel/${lokasi_wh}`); // Replace with your API endpoint
    //                 const result = await response.json();
    //                 setData(result.data)
    //             }else if(lokasi_wh.includes("WITEL")){
    //                 const response = await fetch(`http://localhost:8080/api/get_rekap_delivery_treg/witel/${lokasi_wh}`); // Replace with your API endpoint
    //                 const result = await response.json();
    //                 setData(result.data)
    //             }
    //         }
    //         if(response){
    //             setPrevLokasiWH(lokasi_wh)
    //             console.log(lokasi_wh)
    //         }
    //     }catch (error){
    //         setError(error.message)
    //     }finally {
    //         setLoading(false);
    //     }
    // };

    // useEffect(() => {
    //     fetchData()
    // }, [])

    // useEffect(() => {
    //     setData(data)
    // },[data])

    // const [prevLokasiWH, setPrevLokasiWH] = useState([]);

    // const navigate = useNavigate();
    //     const handleLinkClick = (event, locationWh) => {
    //         event.preventDefault(); // Prevent the default link behavior
    //         // navigate(`/scmt/rekap_delivery/witel/${locationWh}`, { replace: true });
    //         navigate(`/scmt/rekap_delivery/witel/${locationWh}`);
    //         fetchData()
    //         // window.location.reload();
    //         console.log(locationWh)
    //   };

    // const handleLinkBackClick = (event) => {
    //     // event.preventDefault(); // Prevent the default link behavior
    //     // navigate(`/scmt/rekap_delivery/${prevLokasiWH}`, { replace: true });
    //     fetchData();
    //     navigate(-1); // Go back to the previous page
    //     // window.location.reload();
    //   };

    const { lokasi_wh } = useParams();
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [history, setHistory] = useState([]);

    const navigate = useNavigate();

    const fetchData = async () => {
        try {
            let response;
            if (!lokasi_wh) {
                response = await fetch(`http://localhost:8080/api/get_rekap_delivery_treg`);
            } else {
                response = await fetch(`http://localhost:8080/api/get_rekap_delivery_treg/witel/${lokasi_wh}`);
            }

            if (response.ok) {
                const result = await response.json();
                setData(result.data);
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
        console.log(history)
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
                                                {data && 
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