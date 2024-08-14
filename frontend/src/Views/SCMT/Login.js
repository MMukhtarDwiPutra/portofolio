import React, {Component, useEffect, StrictMode, useState  } from 'react'
import Logo from '../../Assets/img/minitok.jpeg'
import Image from 'react-bootstrap/Image';
import {useNavigate } from "react-router-dom";
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";

const Login = () => {
	const [data, setData] = useState({
	    username: '',
	    password: '',
	  });

	const [redirect, setRedirect] = useState(false)
	const [message, setMessage] = useState('');

	const handleChange = (e) => {
	    const { name, value } = e.target;
	    setData((prevData) => ({
	      ...prevData,
	      [name]: value,
	    }));
	};

	const handleSubmit = async (e) => {
	    e.preventDefault();

	    try {
	      const response = await fetch('http://localhost:8080/scmt/login', {
	        method: 'POST',
	        headers: {
	          'Content-Type': 'application/json',
	        },
	        credentials: "include",
	        body: JSON.stringify(data),
	      });

	      const result = await response.json();
	      if (result.data.message != "Login sukses!") {
	        setMessage(result.data.message);
	        setData({
			  username: '',
			  password: ''
			});
	      }else{
	          console.log('Success:', result);
	      	  setRedirect(true)
	      }
	    } catch (error) {
	      console.error('Error:', error);
	    }
	};

	const navigate = useNavigate();

    useEffect(() => {
        if (redirect) {
            navigate('/scmt/rekap_delivery');
        }
    }, [redirect, navigate]);
	return(
		<div className="wrapper-page mt-5">
            <div className="row h-100">
                <div className="col-sm-12 d-flex justify-content-center">
                        <div className="card mx-auto my-auto" style={{width:"500px"}}>
                        {message && (
			                        <div className="alert alert-success alert-dismissible fade show ml-3 mr-3 mt-3" role="alert">
			                            <strong style={{fontSize:"15px", fontWeight:"bold"}}>{message}</strong>
			                        </div>
			                    )}
                            <div className="card-body">
                            	<div className="row d-flex justify-content-end">
                            		<div className="col-3">
                            		<Link className="btn btn-secondary" to="/scmt/rekap_delivery">Dashboard</Link>
                            		</div>
                            	</div>
                                <h4 className="text-center font-size-20" style={{color: "#EE1C25"}}><b><Image src={Logo} className="img-fluid" alt="Logo" width="120px" height="120px"/></b></h4>
                                <div className="text-center">
                                    <span>
                                        <span className="font-weight-bold" style={{fontSize:"17px"}}>Minimum Stock</span><br/>
                                        The Ultimate Novelty Tools As Solutions
                                    </span>
                                </div>

                                <hr/>

            
                                <div className="p-3">
                                    <form className="form-horizontal mt-3" onSubmit={handleSubmit} method="POST" encType="multipart/form-data">
                                        <div className="form-group mb-3 row">
                                            <div className="col-12">
                                                <input className="form-control" type="text" name="username" onChange={handleChange} value={data.username} required="" placeholder="Username"/>
                                            </div>
                                        </div>
            
                                        <div className="form-group mb-3 row">
                                            <div className="col-12">
                                                <input className="form-control" type="password" name="password" required="" onChange={handleChange} value={data.password} placeholder="Password"/>
                                            </div>
                                        </div>
            
                                        <div className="form-group mb-3 text-center row mt-3 pt-1">
                                            <div className="col-12">
                                                <button className="btn btn-info w-100 waves-effect waves-light" id="btn_login" type="submit">Log In</button>
                                            </div>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </div>
                </div>
            </div>
        </div>
	)
} 

export default Login;