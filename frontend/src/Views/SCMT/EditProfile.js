import Sidebar from "./Components/Sidebar"
import Navbar from "./Components/Navbar"
import React, {Component, useEffect, StrictMode, useState  } from 'react'

const EditProfile = () =>{
	const [message, setMessage] = useState('');
	const [dataUser, setDataUser] = useState([])
	const [dataForm, setDataForm] = useState({
	    fullname: dataUser?.fullname || "",
	    password_lama: "",
	    password_baru: "",
	    confirm_password_baru: "",
	});

	const fetchDataUser = async () =>{
        try{
            let response;
            response = await fetch(`http://localhost:8080/api/user`, { 
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                })

            const result = await response.json();
            if(result.data){
            	setDataUser(result.data)
            }else{
            	setDataUser([])
            }
        }catch(error){
        }
    }

    const handleChange = (e) => {
	    const { name, value } = e.target;
	    setDataUser((prevData) => ({
	      ...prevData,
	      [name]: value,
	    }));
	  };

    useEffect(() => {
        fetchDataUser();
    }, []);

    const changePasswordSubmitHandler = async(e) =>{
    	e.preventDefault();
    	if(dataUser.password_baru != dataUser.confirm_password_baru){
    		setMessage("Password baru yang dimasukan tidak sama dengan confirm password baru!")
    	}else{
	    	const dataTmp = new FormData();
		    dataTmp.append('password_lama', dataUser.password_lama);
		    dataTmp.append('password_baru', dataUser.password_baru);
    		try{
    			let response;
    			response = await fetch(`http://localhost:8080/api/scmt/change_password/${dataUser.id}`, {
                    credentials: 'include',
	                method: 'PUT',
	        		body: dataTmp
    			})

    			const result = await response.json();
    			setMessage(result.data.message)
    		}catch(error){
    			setMessage(error)
    		}
    	}
    }

    const changeDataSubmitHandler = async(e) =>{
    	e.preventDefault();
    	const dataTmp = new FormData();
	    dataTmp.append('fullname', dataUser.fullname);
	    console.log(dataTmp)

		try{
			let response;
			response = await fetch(`http://localhost:8080/api/scmt/change_data_user/${dataUser.id}`, {
                credentials: 'include',
                method: 'PUT',
        		body: dataTmp
			})

			const result = await response.json();
			if(result.data){
				setMessage(result.data.message)
			}
		}catch(error){
			setMessage(error)
		}
    }

	return(
		<>
		<div className="wrapper d-flex align-items-stretch">
	        <Sidebar/>
	        <div id="content" style={{margin: "0 auto",  boxSizing: "border-box"}}>
	            <div className="container-fluid" style={{width: "105.5%"}}>
	                <Navbar/>
	                <div className="container-fluid">
	                    {message && (
	                        <div className="alert alert-success alert-dismissible fade show ml-3 mr-3 mb-3 mt-5" role="alert">
	                            <strong style={{fontSize:"15px", fontWeight:"bold"}}>{message}</strong>
	                        </div>
	                    )}
	                    <section className="container my-2 p-2">
	                        <form onSubmit={changeDataSubmitHandler} method='PUT' encType='multipart/form-data' className="row g-3 p-3">
	                            <div className="col-md-12">
	                                <label className="form-label">Fullname</label>
	                                <input type="text" className="form-control" onChange={handleChange} id="fullname" defaultValue={dataUser.fullname} name="fullname" required/>
	                            </div>

	                            <div className="col-md-12 mt-3">
	                                <button type="submit" className="btn btn-primary float-right">Change Fullname</button>
	                            </div>
	                        </form>
	                        <form onSubmit={changePasswordSubmitHandler} className="row g-3 p-3">
	                            <div className="col-md-12 mt-2">
	                                <label className="form-label">Password Lama</label>
	                                <input onChange={handleChange} type="password" className="form-control" id="password_lama" name="password_lama"/>
	                            </div>
	                            <div className="col-md-6 mt-2">
	                                <label className="form-label">Password Baru</label>
	                                <input onChange={handleChange} type="password" className="form-control" id="password_baru" name="password_baru"/>
	                            </div>
	                            <div className="col-md-6 mt-2">
	                                <label className="form-label">Confirm Password Baru</label>
	                                <input onChange={handleChange} type="password" className="form-control" id="confirm_password_baru" name="confirm_password_baru"/>
	                            </div>

	                            <div className="col-md-12 mt-3 ml-auto">
	                                <button type="submit" className="btn btn-primary float-right">Change Password</button>
	                            </div>
	                        </form>
	                    </section>
	                </div>
	            </div>
	        </div>
    	</div>
    </>
    )
}

export default EditProfile;