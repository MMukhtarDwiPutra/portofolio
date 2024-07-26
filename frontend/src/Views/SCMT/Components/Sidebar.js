import React, {Component, useEffect, StrictMode, useState  } from 'react'
import "../../../Assets/css/Sidebar_SCMT.css"
import jQuery from 'jquery'
import {Helmet} from "react-helmet"
import useScript from '../../../Assets/js/UseScript';
import "../../../Assets/js/Sidebar_SCMT.js"
import 'jquery/dist/jquery.min.js'
import 'bootstrap/dist/js/bootstrap.min.js'

export default function Sidebar(){
	const [sidebarActive, setSidebarActive] = useState(true);

  const toggleSidebar = () => {
    setSidebarActive(!sidebarActive);
  };

  const [collapseActive, setCollapseActive] = useState([true, true, true, true]);

  // Fungsi untuk toggle collapse
  const toggleCollapse = (index) => {
    console.log(`Dropdown ${index} clicked`);

    // Salin array state saat ini
    const newCollapseActive = [...collapseActive];

    // Toggle status collapse untuk dropdown yang diklik
    newCollapseActive[index] = !newCollapseActive[index];

    // Perbarui state
    setCollapseActive(newCollapseActive);
  };

	return(
		<>
		<nav id="sidebar" className={sidebarActive ? 'active' : ''}>
		    <div className="custom-menu">
		        <button type="button" id="sidebarCollapse" className="btn btn-primary" onClick={toggleSidebar}>
		            <i className="fa fa-bars"></i>
		        <span className="sr-only">Toggle Menu</span>
		        </button>
		    </div>

			<div className="p-4 pt-5">
			    <h1><a href="index.html" className="logo">Menu</a></h1>
			    <ul className="list-unstyled components">
			        <li key="0" className={collapseActive[0] ? 'active' : ''}>
			            <a href="#" data-toggle="collapse" onClick={() => toggleCollapse(0)} aria-expanded="false" className="dropdown-toggle">Minimum Stock ONT</a>
			            <ul className={collapseActive[0] ? 'collapse list-unstyled components' : 'list-unstyled components'} id="SubMenu0">
			                <li>
			                    <a className="nav-link active" href="{{url('/rekap_delivery')}}">Rekap Minimum Stock ONT</a>
			                </li>
			                <li>
			                    <a className="nav-link active" href="{{url('/pengiriman_ont')}}">Report Delivery ONT</a>
			                </li>
			            </ul>
			        </li>


			        <li key="1" className={collapseActive[1] ? 'active' : ''}>
			            <a href="#" data-toggle="collapse" onClick={() => toggleCollapse(1)} aria-expanded="false" className="dropdown-toggle">Minimum Stock STB</a>
			            <ul className={collapseActive[1] ? 'collapse list-unstyled components' : 'list-unstyled components'} id="SubMenu1">
			                <li>
			                    <a className="nav-link active" href="{{url('/stb/rekap_delivery_stb')}}">Rekap Minimum Stock STB</a>
			                </li>
			                <li>
			                    <a className="nav-link active" href="{{url('/pengiriman_stb')}}">Report Delivery STB</a>
			                </li>
			            </ul>
			        </li>

			        <li key="2" className="active">
			            <a href="#" data-toggle="collapse" aria-expanded="false" onClick={() => toggleCollapse(2)} className="dropdown-toggle">Minimum Stock AP</a>
			            <ul className={collapseActive[2] ? 'collapse list-unstyled components' : 'list-unstyled components'} id="SubMenu2">
			                <li>
			                    <a className="nav-link active" href="{{url('/rekap_delivery_ap')}}">Rekap Minimum Stock AP</a>
			                </li>
			                <li>
			                    <a className="nav-link active" href="{{url('/pengiriman_ap')}}">Report Delivery AP</a>
			                </li>
			            </ul>
			        </li>

			        <li key="3" className="active">
			            <a href="#" data-toggle="collapse" onClick={() => toggleCollapse(3)} aria-expanded="false" className="dropdown-toggle">Input Data</a>
			            <ul className={collapseActive[3] ? 'collapse list-unstyled components' : 'list-unstyled components'} id="SubMenu3">
			                <li>
			                    <a className="nav-link active" href="{{url('/input_data_stock')}}">Upload File Stock</a>
			                </li>
			                <li>
			                    <a className="nav-link active" href="{{url('/input_data_pengiriman')}}">Update File Delivery</a>
			                </li>
			                <li>
			                    <a className="nav-link active" href="{{url('/input_data_database')}}">Upload File Minimum Stock</a>
			                </li>
			                <li>
			                    <a className="nav-link active" href="{{url('/input_sn_mac_vendor')}}">Upload File SN Vendor</a>
			                </li>
			            </ul>
			        </li>

			        <li>
			            <a className="nav-link active" style={{color:"#fff"}} href="{{url('/request_outbond')}}">Request Outbond</a>
			        </li>
			        <li className="justify-content: end;">
			            <a className="nav-link active" style={{color:"#fff"}} href="{{url('/edit_profile')}}">Edit Profile</a>
			        </li>
			    </ul>     
			</div>
		</nav>
		</>
	)
}