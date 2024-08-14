import React, {Component, useEffect, StrictMode, useState  } from 'react'
import "../../../Assets/css/Sidebar_SCMT.css"
import jQuery from 'jquery'
import {Helmet} from "react-helmet"
import useScript from '../../../Assets/js/UseScript';
import "../../../Assets/js/Sidebar_SCMT.js"
import 'jquery/dist/jquery.min.js'
import 'bootstrap/dist/js/bootstrap.min.js'
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";

export default function Sidebar(){
	useEffect(() => {
    const link = document.createElement('link');
    link.rel = 'stylesheet';
    link.href = 'https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css';
    document.head.appendChild(link);

    return () => {
      document.head.removeChild(link);
    };
  }, []);

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

  const [login, setLogin] = useState(false)
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
				                    <Link className="nav-link active" to="/scmt/rekap_delivery">Rekap Minimum Stock ONT</Link>
				                </li>
				                <li>
					                <Link className="nav-link active" to="/scmt/report_delivery_ont">Report Delivery ONT</Link>
				                </li>
				            </ul>
				        </li>


				        
				        {login && (
				        <li key="3" className="active">
				            <a href="#" data-toggle="collapse" onClick={() => toggleCollapse(3)} aria-expanded="false" className="dropdown-toggle">Input Data</a>
				            <ul className={collapseActive[3] ? 'collapse list-unstyled components' : 'list-unstyled components'} id="SubMenu3">
				                <li>
				                 <Link className="nav-link active" to="/scmt/upload_file_data_stock">Upload File Stock</Link>
				                </li>
				                <li>
				                	<Link className="nav-link active" to="/scmt/upload_file_pengiriman">Update File Delivery</Link>
				                </li>
				                <li>
				                	<Link className="nav-link active" to="/scmt/upload_file_minimum_stock">Upload File Minimum Stock</Link>
				                </li> 
				            </ul>
				        </li>
				        )}
				        {login && (
				        <li className="justify-content: end;">
				            <a className="nav-link active" style={{color:"#fff"}} href="/scmt/edit_profile">Edit Profile</a>
				        </li>
				        )}
				    </ul>     
				</div>
			</nav>
		</>
	)
}