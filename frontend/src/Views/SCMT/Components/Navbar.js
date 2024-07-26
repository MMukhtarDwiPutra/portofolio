import React from "react"

import Image from 'react-bootstrap/Image';
import Logo from '../../../Assets/img/minitok.jpeg'
import Bell from '../../../Assets/img/bell.png'

export default function Navbar(){
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
			            <div className="icon" onClick="toggleNotifi()">
			                <div className="d-flex justify-content-end">
			                    <a className="active btn btn-danger align-middle my-auto">
			              			<img src={Bell} alt=""  style={{color: "white", height: "40px"}} /><span id="notifesHeader"></span>
			                        
			                    </a>
			                </div>
			            </div>
			        </div>

			        <div className="col-md-2" style={{paddingTop: "30px", paddingBottom: "5px"}}>
			            <div className="d-flex justify-content-end">
			                <a className="active btn btn-danger align-middle my-auto" style={{color: "white", height: "40px"}}
			                    href="">Logout</a>
			            </div>
			        </div>
			    </div>
			</div>

		</>
	)
}