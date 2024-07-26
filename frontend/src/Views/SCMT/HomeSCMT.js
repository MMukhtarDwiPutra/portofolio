import React from 'react'
import Sidebar from "./Components/Sidebar"
import Navbar from "./Components/Navbar"
import "../../Assets/js/Sidebar_SCMT.js"

export default function HomeSCMT(){
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
                                      	<div id="tableBiasa">
                                            <table className="table table-bordered" id="filterTable" width="100%">
                                                <thead>
                                                	<tr className="text-center" style={{color: "black", background:"gray", color:"white"}}>
                                                        <th rowSpan="2" className="first-col sticky-col" style={{minWidth: "250px", textAlign: "center", verticalAlign: "middle"}}>
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
                                                        <th colSpan="2">Minimum Stock Requirement Retail</th>
                                                        <th colSpan="2">On Delivery</th>
                                                    </tr>
                                                    <tr className="text-center" style={{background:"gray", color:"white"}}>
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