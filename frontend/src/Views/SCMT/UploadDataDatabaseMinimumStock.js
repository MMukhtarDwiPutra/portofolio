import Sidebar from "./Components/Sidebar"
import Navbar from "./Components/Navbar"
import React, {Component, useEffect, StrictMode, useState, useRef  } from 'react'

const UploadDataDatabaseMinimumStock = () =>{
	const [fileGudang, setFileGudang] = useState(null);
	const [message, setMessage] = useState('');
  	const fileInputRef = useRef(null); // Create a reference for the file input

	const handleFileGudangChange = (e) => {
		setFileGudang(e.target.files[0]);
	};

	const handleUploadGudang = async (e) => {
	    e.preventDefault();

	    if (!fileGudang) {
	      setMessage('No file selected');
	      return;
	    }
		const formData = new FormData();
	    formData.append('file', fileGudang);

	    try {
	      const response = await fetch('http://localhost:8080/api/upload_gudang', {
            credentials: 'include',
	        method: 'POST',
	        body: formData,
	      });

	      if (!response.ok) {
	        setMessage('File upload failed');
	      }else{
		      const responseData = await response.json();
			  setMessage("Success Upload Replace Data Gudang!");
	      }
	    } catch (error) {
	      console.error('Error:', error);
	    } finally{
	    	setFileGudang(null); // Reset state
      		fileInputRef.current.value = '';
	    }
	  };
	return(
		<>
			<div className="wrapper d-flex align-items-stretch">
			<Sidebar/>
			 <div id="content" style={{margin: "0 auto", boxSizing: "border-box"}}>
            	<div className="container-fluid" style={{width: "105.5%"}}>
            	<Navbar/>
            		<div className="card">
            		{message && (
	                        <div className="alert alert-success alert-dismissible fade show ml-3 mr-3 mb-3 mt-5" role="alert">
	                            <strong style={{fontSize:"15px", fontWeight:"bold"}}>{message}</strong>
	                        </div>
	                    )}
		                    <form className="mt-3"  style={{paddingLeft: "15px", paddingRight: "15px", paddingBottom: "15px"}}
		                        method="POST" onSubmit={handleUploadGudang} 
		                        encType="multipart/form-data">
		                        <div className="form-group row mt-1">
		                            <label>Upload untuk update database minimum stock:</label>
		                        </div>
		                        <div className="row mt-3">
		                            <div className="col-sm-12 col-md-4 mb-3">
		                                <input className="form-control" ref={fileInputRef} onChange={handleFileGudangChange} type="file" style={{height: "45px"}} name="file_database"
		                                    required/>
		                            </div>
		                            <div className="col-sm-12 col-md-8">
		                                <div className="row">
		                                    <div className="col-md-6 mb-2">
		                                        <button type="submit" className="btn btn-primary "
		                                           style={{height: "40px"}}>Upload</button>
		                                        <a href="http://localhost:8080/api/download_template_minimum_stock"
		                                            className="btn btn-secondary ml-1">Download Template</a>

		                                    </div>
		                                    <div className="col-md-3"></div>
		                                    <div className="col-md-3 d-flex justify-content-end" style={{height: "40px"}}>
		                                        <a className="btn btn-warning" href="http://localhost:8080/api/export_minimum_stock_database">Export
		                                            Database</a>
		                                    </div>

		                                </div>
		                            </div>
		                        </div>
		                    </form>
		                </div>
            		</div>
            	</div>
			</div>
		</>
	)
}

export default UploadDataDatabaseMinimumStock