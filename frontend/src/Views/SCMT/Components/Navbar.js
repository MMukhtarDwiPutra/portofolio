import Image from 'react-bootstrap/Image';
import Logo from '../../../Assets/img/minitok.jpeg'
import Bell from '../../../Assets/img/bell.png'
import React, {Component, useEffect, StrictMode, useState  } from 'react'
import {useNavigate } from "react-router-dom";

export default function Navbar(){
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

	const [redirect, setRedirect] = useState(false)
  	const handleLogout = async (e) => {
	    e.preventDefault();

	    try {
	      const response = await fetch('http://localhost:8080/scmt/logout', {
	        method: 'GET',
	        headers: {
	          'Content-Type': 'application/json',
	        },
	        credentials: "include",
	      });

	      setRedirect(true)
	    } catch (error) {
	      console.error('Error:', error);
	    }
	};

	const navigate = useNavigate();

    useEffect(() => {
        if (redirect) {
            navigate('/scmt/rekap_delivery');
            window.location.reload();
        }
    }, [redirect, navigate]);

	return(
		<>
			<div className="container-fluid" style={{width: "100%"}}>
			    <div className="row">
			        <div className="col-md-8">
			            <div className="row">
			                <div className="col-md-4 col-12">
			                    <a><Image src={Logo} className="img-fluid" alt="Logo" style={{width: "120px", height: "auto", maxWidth: "100%"}}/>
			                    </a>
			                </div>
			                <div className="col-md-8 col-12">
			                    <div className="d-flex flex-column align-items-center text-center">
			                        <div className="font-weight-bold">
			                            <span className="d-md-none" style={{fontSize: "20px"}}>Minimum Stock</span><br/>
			                            <span className="d-none d-md-block"
			                                style={{fontSize: "25px", overflow: "hidden", whiteSpace: "nowrap", textOverflow: "ellipsis"}}>Minimum
			                                Stock</span>

			                            <span className="d-md-none" style={{fontSize: "12px"}}>The Ultimate Novelty Tools As
			                                Solutions</span>
			                            <span className="d-none d-md-block" style={{fontSize: "20px"}}>The Ultimate Novelty Tools As
			                                Solutions</span>
			                        </div>
			                    </div>
			                </div>
			            </div>
			        </div>

			        <div className="col-md-2" style={{paddingTop: "30px", paddingBottom: "5px"}}>
			            <div className="icon" >
			                <div className="d-flex justify-content-end">
			                    <a className="active btn btn-danger align-middle my-auto">
			              			<img src={Bell} alt=""  style={{color: "white", height: "40px"}} /><span id="notifesHeader"></span>
			                        
			                    </a>
			                </div>
			            </div>
			        </div>

			        <div className="col-md-2" style={{paddingTop: "30px", paddingBottom: "5px"}}>
			            <div className="d-flex justify-content-end">
			            {login ? (
			                <a className="active btn btn-danger align-middle my-auto" style={{color: "white", height: "40px"}}
			                    onClick={handleLogout}>Logout</a>
			                ) : (
			                <a className="btn btn-primary" href="/scmt/login">Login</a>
			                )
			            }
			            </div>
			        </div>
			    </div>
			</div>

		</>
	)
}