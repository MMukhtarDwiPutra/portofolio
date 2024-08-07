import Sidebar from "./Components/Sidebar"
import Navbar from "./Components/Navbar"
import React, {Component, useEffect, StrictMode, useState, useRef  } from 'react'

const UploadDataPengiriman = () =>{
    const [filePenerima, setFilePenerima] = useState(null);
	const [message, setMessage] = useState('');
  	const fileInputRef = useRef(null); // Create a reference for the file input

	const handleFilePenerimaChange = (e) => {
		setFilePenerima(e.target.files[0]);
	};

	const handleUploadPenerima = async (e) => {
	    e.preventDefault();

	    if (!filePenerima) {
	      console.log('No file selected');
	      return;
	    }
		const formData = new FormData();
	    formData.append('file', filePenerima);

	    try {
	      const response = await fetch('http://localhost:8080/api/tambah_penerima_bulk/replace', {
            credentials: 'include',
	        method: 'POST',
	        body: formData,
	      });

	      if (!response.ok) {
	        throw new Error('File upload failed');
	      }else{
		      const responseData = await response.json();
			  setMessage("Success Upload Replace Data Penerimaan!");
	      }
	    } catch (error) {
	      console.error('Error:', error);
	    } finally{
	    	setFilePenerima(null); // Reset state
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
            		<div className="mb-3">
		               <div className="card mt-3 mr-3">
		               {message && (
	                        <div className="alert alert-success alert-dismissible fade show ml-3 mr-3 mb-3 mt-5" role="alert">
	                            <strong style={{fontSize:"15px", fontWeight:"bold"}}>{message}</strong>
	                        </div>
	                    )}
		                        <form onSubmit={handleUploadPenerima} className="mt-3" style={{paddingLeft: "15px", paddingRight: "15px", paddingBottom: "15px"}}
		                        method="POST" encType="multipart/form-data">
		                            <div className="form-group row">
		                                <label className="ml-2">Masukan file pengiriman untuk replace data di database
		                                    pengiriman:</label>
		                            </div>
		                            <div className="row mt-3">
		                                <div className="col-sm-12 col-md-4 mb-3">
		                                    <input className="form-control" type="file" style={{height: "45px"}}
		                                        name="file_penerima" ref={fileInputRef} onChange={handleFilePenerimaChange} required/>
		                                </div>
		                                <div className="col-sm-12 col-md-8">
		                                    <div className="row">
		                                        <div className="col-md-6 mb-1">
		                                            <button type="submit" className="btn btn-primary "
		                                                style={{height:"40px"}}>Upload</button>
		                                            <a href="{{ URL('download_template_penerima') }}"
		                                                className="btn btn-secondary ml-1">Download
		                                                Template</a>

		                                        </div>
		                                        <div className="col-md-1"></div>
		                                        <div className="col-md-5 d-flex justify-content-end">
		                                            <a href="http://localhost:8080/api/export_all_penerima" className="btn btn-secondary mr-3"
		                                                style={{height: "40px"}}>Export All Data Pengiriman</a>
		                                        </div>

		                                    </div>
		                                </div>
		                            </div>
		                        </form>
		                    </div>
		                </div>
            		</div>
            	</div>
			</div>
		</>
	)
}

export default UploadDataPengiriman