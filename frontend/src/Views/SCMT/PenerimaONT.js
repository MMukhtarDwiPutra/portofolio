import Sidebar from "./Components/Sidebar"
import Navbar from "./Components/Navbar"
import React, {Component, useEffect, StrictMode, useState, useRef  } from 'react'
import DataTable from 'react-data-table-component';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import EditModal from './Components/EditModalPenerima';
import DeleteAllModal from './Components/DeleteAllModalPenerima';
import TambahModal from './Components/TambahModalPenerima';
import UploadIDOGDModal from './Components/UploadIDOGDModal';
import DeleteByIdModal from './Components/DeleteById';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';

const PenerimaONT = () => {
	const [showModalEdit, setShowModalEdit] = useState(false);
	const [currentItemEdit, setCurrentItemEdit] = useState(null);
	const fileInputRef = useRef(null); // Create a reference for the file input
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

	const handleEditClick = (id) => {
		const item = dataPenerima.find((item) => item.id === id);
		setCurrentItemEdit(item);
		console.log(item)
		setShowModalEdit(true);
	};

	const handleCloseModalEdit = () => {
		setShowModalEdit(false);
		setCurrentItemEdit(null);
	};

	const handleSaveChangesEdit = () => {
		console.log('Saving changes for:', currentItemEdit);
		// Implement save logic here
		handleCloseModalEdit();
	};

	const [uniqueBatch, setUniqueBatch] = useState([]);

	const [showModalDeleteByID, setShowModalDeleteByID] = useState(false);
	const [currentItemDeleteByID, setCurrentItemDeleteByID] = useState(null);

	const handleDeleteByID = (id) => {
		const item = dataPenerima.find((item) => item.id === id);
		setCurrentItemDeleteByID(item);
		setShowModalDeleteByID(true);
	};

	const handleCloseDeleteByID = () => {
		setShowModalDeleteByID(false);
		setCurrentItemDeleteByID(null);
	};

	const handleDeleteByIdClick = async () => {
		try {
		  	const response = await fetch(`http://localhost:8080/api/delete_penerima/${currentItemDeleteByID.id}`, {
			    method: 'DELETE',
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
		  });

		  if (!response.ok) {
		    throw new Error('Network response was not ok');
		  }else{
			  const result = await response.json();

			  setMessage('Success Delete');
			  handleCloseDeleteByID()
			  await fetchDataPenerima()
		  }
		} catch (error) {
		  console.error('Error:', error);
		}
	};

	const [showModalDeleteAll, setShowModalDeleteAll] = useState(false);

	const handleDeleteAll = () => {
		setShowModalDeleteAll(true);
	};

	const handleCloseModalDeleteAll = () => {
		setShowModalDeleteAll(false);
	};

	const handleDeleteAllClick = async () => {
		try {
		  	const response = await fetch(`http://localhost:8080/api/delete_all_penerima`, {
			    method: 'DELETE',
			    headers: {
			      'Content-Type': 'application/json',
			    },
			    credentials: 'include',
		  });

		  if (!response.ok) {
		    throw new Error('Network response was not ok');
		  }else{
			  const result = await response.json();
			  setMessage('Success Delete All');
			  handleCloseModalDeleteAll()
			  await fetchDataPenerima()
		  }

		} catch (error) {
		  console.error('Error:', error);
		}
	};

	const [showModalUploadIDOGD, setShowModalUploadIDOGD] = useState(false);

	const handleUploadIDOGD = () => {
		setShowModalUploadIDOGD(true);
	};

	const handleCloseModalUploadIDOGD = () => {
		setShowModalUploadIDOGD(false);
	};

	const handleUploadIDOGDClick = () => {
		//Implements upload ido gd API
		handleCloseModalDeleteAll();
	};

	const [showModalTambah, setShowModalTambah] = useState(false);

	const handleTambahModal = () => {
		setShowModalTambah(true);
	};

	const handleCloseModalTambah = () => {
		setShowModalTambah(false);
	};

	const handleTambahClick = () => {
		//Implements tambah click API
		fetchDataPenerima();
		handleCloseModalTambah();
	};

	const [message, setMessage] = useState('');
	const [jenisAkun, setJenisAkun] = useState('');

	const fetchJenisAkun = () => {
	  // Simulate getting an account type (e.g., from a session)
	  const sessionJenisAkun = 'Admin'; // This would be dynamic in a real app
	  setJenisAkun(sessionJenisAkun);
	};

	const [asal, setAsal] = useState('');

	const fetchSessionAsal = () => {
	  // Simulate getting an account type (e.g., from a session)
	  const sessionAsal = 'DID'; // This would be dynamic in a real app
	  setAsal(sessionAsal);
	};

	const [statusFillingDisable, setStatusFillingDisable] = useState('');

	const fetchStatusFillingDisable = () => {
	  // Simulate getting an account type (e.g., from a session)
	  const sessionStatusFillingDisable = 'OFF'; // This would be dynamic in a real app
	  setStatusFillingDisable(sessionStatusFillingDisable);
	};
	
	const [dataWarehouse, setDataWarehouse] = useState([]);
	const [dataPenerima, setDataPenerima] = useState([]);
	const [data, setData] = useState([]);
	const [loading, setLoading] = useState(true);
  	const [error, setError] = useState(null);
  	const [lastUpdate, setLastUpdate] = useState(null)

	const fetchDataPenerima = async () => {
		// Simulate an API call
		try {
			// console.log("HIT Get Fetch Data Penerima")
			const response = await fetch('http://localhost:8080/api/get_pengiriman_ont', {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
            });; // Replace with your API endpoint

			const result = await response.json();
			if (result.data.penerima != null) {
			  setDataPenerima(result["data"].penerima);
			  setDataWarehouse(result["data"].warehouse)
			  setData(result["data"].penerima)
			  setLastUpdate(result["data"].last_update)
			}else{
			  	setDataPenerima([]);
			  	setData([])
			}
		}catch (error){
			setError(error.message)
		}finally {
			setLoading(false);
		}
	};


	useEffect(() => {
		// Fetch data from an API or define it directly
		try {
			fetchDataPenerima()
			fetchStatusFillingDisable();
			fetchJenisAkun();
			fetchSessionAsal();
		}catch (error){
			console.log("Error fetching API get pengiriman ONT")
		}

	}, []);

	useEffect(() => {
	    // Process unique batches once dataPenerima is updated
	    const batchs = dataPenerima.map(item => item.batch.toLowerCase());
	    const uniqueBatchArray = [...new Set(batchs)];
	    setUniqueBatch(uniqueBatchArray);
	  }, [dataPenerima]); // Run this effect whenever dataPenerima changes

	const [filePenerima, setFilePenerima] = useState(null);

	const handleFilePenerimaChange = (e) => {
		setFilePenerima(e.target.files[0]);
	};

	const [filterBatch, setFilterBatch] = useState("");
	const [filterTREG, setFilterTREG] = useState("");

	useEffect(() => {
	    filterData()
	  }, [filterBatch, filterTREG]);

	const handleFilterBatchChange = async (e) => {
		setFilterBatch(e.target.value);
	};

	const handleFilterTREGChange = async (e) => {
		setFilterTREG(e.target.value);
	};

	const filterData = () => {
		const dataTmp = dataPenerima.filter((item) =>
			(item.batch.toLowerCase().includes(filterBatch.toLowerCase()) && item.regional.toLowerCase().includes(filterTREG.toLowerCase()))
		);
		
		setData(dataTmp)
	}

	const formData = new FormData();
    formData.append('file', filePenerima);

	const handleUploadPenerima = async (e) => {
	    e.preventDefault();

	    if (!filePenerima) {
	      console.log('No file selected');
	      return;
	    }

	    try {
	      const response = await fetch('http://localhost:8080/api/tambah_penerima_bulk/add_only', {
            credentials: 'include',
	        method: 'POST',
	        body: formData,
	      });

	      if (!response.ok) {
		    await fetchDataPenerima()
	        throw new Error('File upload failed');
	      }else{
		      const responseData = await response.json();
			  setMessage("Success Upload Data Penerimaan!");
		  	  await fetchDataPenerima()
	      }
	    } catch (error) {
	      console.error('Error:', error);
	    } finally{
	    	setFilePenerima(null); // Reset state
      		fileInputRef.current.value = '';
	    }
	  };

	 const columns = [
	    {
	      name: 'No',
	      selector: (row, index) => index + 1,
	      center: 'true',
	      width: '50px',
	    },
	    ...(login 
	    ? [
		    {
		      name: 'Action',
		      selector: (row) => row.id,
		      cell: (row) => {
		        let btn = (
		          <button
		            className="btn btn-danger"
		            onClick={() => handleDeleteByID(row.id)}
		          >
		            Delete
		          </button>
		        );

		        return btn;
		      },
		      center: 'true',
		      width: '120px',
	    	},
	      ]
	    : []),
	    {
	      name: 'Type',
	      selector: (row, index) => row.type,
	      center: 'true',
	      width: '150px',
	    },{
	      name: 'Qty',
	      selector: (row, index) => row.qty,
	      center: 'true',
	      width: '70px',
	    },{
	      name: 'Alamat Pengirim',
	      selector: (row, index) => row.alamat_pengirim,
	      center: 'true',
	      width: '170px',
	    },{
	      name: 'PIC Pengirim',
	      selector: (row, index) => row.pic_pengirim,
	      center: 'true',
	      width: '150px',
	    },{
	      name: 'Alamat Penerima',
	      selector: (row, index) => row.alamat_penerima,
	      width: '400px',
	    },{
	      name: 'Warehouse Penerima',
	      selector: (row, index) => row.warehouse_penerima,
	      center: 'true',
	      width: '250px',
	    },{
	      name: 'PIC Penerima',
	      selector: (row, index) => row.pic_penerima,
	      center: 'true',
	      width: '200px',
	    },{
	      name: 'Tanggal Pengiriman',
	      selector: (row, index) => row.tanggal_pengiriman,
	      center: 'true',
	      width: '120px',
	    },{
	      name: 'Tanggal Sampai',
	      selector: (row, index) => row.tanggal_sampai,
	      center: 'true',
	      width: '120px',
	    },{
	      name: 'Batch',
	      selector: (row, index) => row.batch,
	      center: 'true',
	      width: '120px',
	    },
	    ...(login 
	     ? [
	    {
	      name: 'Edit',
	      selector: (row) => row.id,
	      cell: (row) => {
	        const href = `http://localhost:8080/api/download_serial_number/${row.id}`; // URL for download
	        let btn = (
	          <button
	            className="btn btn-warning"
	            onClick={() => handleEditClick(row.id)}
	            data-target="#editModalById"
	            data-toggle="modal"
	          >
	            Edit
	          </button>
	        );

	        if (row.sn_mac_barcode !== '') {
	          btn = (
	            <>
	              {btn}
	              <a className="btn btn-secondary ml-1" href={href}>
	                Download SN
	              </a>
	            </>
	          );
	        }

	        return btn;
	      },
	      className: 'text-center first-col sticky-col',
	      center: 'true',
	      width: '220px',
	    }
	    ]: []),
	]

	return(
		<>
			<div className="wrapper d-flex align-items-stretch">
	        <Sidebar/>
	        <div id="content" style={{margin: "0 auto", boxSizing: "border-box"}}>
	            <div className="container-fluid" style={{width: "105.5%"}}>
					<Navbar/>

					{login && (
                    <div className="card mt-5">
                        <form onSubmit={handleUploadPenerima} className="mt-3" style={{paddingLeft: "15px", paddingRight: "15px", paddingBottom: "15px"}}
                            method="POST" encType="multipart/form-data">
                            <div className="form-group row">
                                <label className="ml-4">Masukan file pengiriman untuk diupload:</label>
                            </div>
                            <div className="row mt-3">
                                <div className="col-sm-12 col-md-4 mb-3">
                                    <input className="form-control" ref={fileInputRef} type="file" style={{height: "45px"}} name="file_penerima" onChange={handleFilePenerimaChange} required></input>
                                </div>
                                <div className="col-sm-12 col-md-8">
                                    <div className="row">
                                        <div className="col-md-6 mb-2">
                                                <button type="submit" className="btn btn-primary">Upload</button>
                                            <a href="http://localhost:8080/api/download_template_penerima"
                                                className="btn btn-secondary ml-1">Download Template</a>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </form>
                    </div>
                    )}

	                <div className="card mb-3 mt-3">
	                    {message && (
	                        <div className="alert alert-success alert-dismissible fade show ml-3 mr-3 mb-3 mt-5" role="alert">
	                            <strong style={{fontSize:"15px", fontWeight:"bold"}}>{message}</strong>
	                        </div>
	                    )}

	                    <div className="card-body mb-2">

	                    	{login && (
		                        <div className="row">
		                            <div className="col-md-6 col-lg-5">
		                                <a className="btn btn-secondary mb-2" href="http://localhost:8080/api/export_all_penerima_ont">Export</a>
		                                <a className="btn btn-secondary mb-2 ml-1" href="http://localhost:8080/api/download_all_sn_ont_exist">Download All SN Exist</a>
		                                <a className="btn btn-secondary mb-2 ml-1" href="http://localhost:8080/api/download_all_sn_ont">Download All SN</a>
		                            </div>

		                            {jenisAkun === "Admin" ? ( 
		                                <div className="col-md-6 col-lg-7 col-xl-7">
		                                    <div className="d-flex justify-content-end align-items-center">
		                                        <a className="btn btn-primary btn-sm mx-1" data-toggle="modal"
		                                            data-target="#tambahModal" onClick={() => handleTambahModal()}>Tambah</a>
		                                        <a className="btn btn-danger btn-sm mx-1" data-toggle="modal"
		                                            data-target="#hapusModal" onClick={() => handleDeleteAll()}>Delete All</a>
		                                    </div>
		                                </div>
		                            ) : null}
		                        </div>
	                        )}
	                        {login && asal === "DID" && ( 
		                        <div className="row mb-3">
		                            <div className="col-md-6 col-lg-5">
		                                <a className="btn btn-primary mb-2" data-toggle="modal" data-target="#editIdoGDBulk" onClick={() => handleUploadIDOGD()} >Upload IDO GD Bulk</a>
		                            </div>
		                        </div>
		                    )}

	                        <div className="row d-flex justify-content-end mb-1 mr-2">
                                <div className="col-auto">
                                    <div className="row">
                                    	<Form.Control id="batch-filter" as="select" onChange={handleFilterBatchChange}>
	                                        <option value="">All</option>
	                                        {uniqueBatch && uniqueBatch.map((item, index) => (
	                                        	<option key={index}>{item}</option>
	                                        ))}
	                                    </Form.Control>
	                                </div>
                                    <div className="row mt-1">
                                    	<Form.Control id="treg-filter" as="select" onChange={handleFilterTREGChange}>
		                                    <option value="">All TREG</option>
		                                    <option value="TREG 1">TREG 1</option>
		                                    <option value="TREG 2">TREG 2</option>
		                                    <option value="TREG 3">TREG 3</option>
		                                    <option value="TREG 4">TREG 4</option>
		                                    <option value="TREG 5">TREG 5</option>
		                                    <option value="TREG 6">TREG 6</option>
		                                    <option value="TREG 7">TREG 7</option>
	                                    </Form.Control>
		                            </div>
	                            </div>
	                        </div>
	                        

	                        {lastUpdate && (<span>Last update: {lastUpdate} </span>)}
	                        <div className="table-responsive mt-2">
	                            {statusFillingDisable === "OFF" ? ( 
	                            <div className="table-responsive">
	                                <DataTable 
	                               		columns={columns}
									      data={data}
									      pagination
									      highlightOnHover
									      striped
									      responsive
	                                		customStyles={{
									        tableWrapper: {
									          style: {
									            display: 'grid',
									            gridTemplateColumns: `repeat(${columns.length}, 1fr)`,
									            gridAutoRows: 'minmax(50px, auto)',
									          },
									        },
									      }}
	                                	pagination
	                                />
	                            </div>
	                            ) : (
	                            <div className="text-center mt-4" style={{backgroundColor: "gray", fontSize: "20px"}}>
	                                <span style={{color:"white"}}>Data report delivery ONT sedang dimaintance, mohon menunggu.</span>
	                            </div>
	                            )}
	                        </div>
	                    </div>
	                </div>
	            </div>
	        </div>
	    </div>

	    <EditModal
			showModal={showModalEdit}
			handleCloseModal={handleCloseModalEdit}
			data={currentItemEdit}
			setCurrentItem={setCurrentItemEdit}
			handleSaveChanges={handleSaveChangesEdit}
			setMessage={setMessage}
			fetchDataPenerima={fetchDataPenerima}
		/>
		<DeleteAllModal
			showModal={showModalDeleteAll}
			handleCloseModal={handleCloseModalDeleteAll}
			handleClick={handleDeleteAllClick}
		/>
		<TambahModal
	   		showModal={showModalTambah}
			handleCloseModal={handleCloseModalTambah}
			handleTambahClick={handleTambahClick}
			dataWarehouse={dataWarehouse}
			setMessage={setMessage}
			fetchDataPenerima={fetchDataPenerima}
	   	/>
	   	<UploadIDOGDModal
	   		showModal={showModalUploadIDOGD}
			handleCloseModal={handleCloseModalUploadIDOGD}
			handleTambahClick={handleUploadIDOGDClick}
	   	/>
	   	<DeleteByIdModal
	   		showModal={showModalDeleteByID}
			handleCloseModal={handleCloseDeleteByID}
			handleClick={handleDeleteByIdClick}
			data={currentItemDeleteByID}
	   	/>
	   	</>
	)
}


export default PenerimaONT;