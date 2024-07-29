import Sidebar from "./Components/Sidebar"
import Navbar from "./Components/Navbar"
import React, {Component, useEffect, StrictMode, useState  } from 'react'
import DataTable from 'react-data-table-component';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import EditModal from './Components/EditModalPenerima';
import DeleteAllModal from './Components/DeleteAllModalPenerima';
import TambahModal from './Components/TambahModalPenerima';
import UploadIDOGDModal from './Components/UploadIDOGDModal';
import DeleteByIdModal from './Components/DeleteById';

const PenerimaONT = () => {
	const [showModalEdit, setShowModalEdit] = useState(false);
	const [currentItemEdit, setCurrentItemEdit] = useState(null);

	const handleEditClick = (id) => {
		const item = dataPenerima.find((item) => item.id === id);
		setCurrentItemEdit(item);
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

	const handleDeleteByIdClick = () => {
		// Implement delete by id API
		handleCloseDeleteByID();
	};

	const [showModalDeleteAll, setShowModalDeleteAll] = useState(false);

	const handleDeleteAll = () => {
		console.log("show up delete all")
		setShowModalDeleteAll(true);
	};

	const handleCloseModalDeleteAll = () => {
		setShowModalDeleteAll(false);
	};

	const handleDeleteAllClick = () => {
		//Implements delete all here API
		handleCloseModalDeleteAll();
	};

	const [showModalUploadIDOGD, setShowModalUploadIDOGD] = useState(false);

	const handleUploadIDOGD = () => {
		console.log("show up ido gd upload")
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
		console.log("show up tambah")
		setShowModalTambah(true);
	};

	const handleCloseModalTambah = () => {
		setShowModalTambah(false);
	};

	const handleTambahClick = () => {
		//Implements tambah click API
		handleCloseModalEdit();
	};

	const [message, setMessage] = useState('');

	// Simulate fetching message from a session or API
	useEffect(() => {
		// Example: Simulate a message from a session or some external source
		const fetchMessage = () => {
		  // Simulate getting a message
		  const sessionMessage = 'This is a session message!';
		  setMessage(sessionMessage);
		}; 

		fetchMessage();
	}, []);

	const [jenisAkun, setJenisAkun] = useState('');

	// Simulate fetching account type from a session or API
	useEffect(() => {
		// Example: Simulate fetching the account type from a session or API
		const fetchJenisAkun = () => {
		  // Simulate getting an account type (e.g., from a session)
		  const sessionJenisAkun = 'Admin'; // This would be dynamic in a real app
		  setJenisAkun(sessionJenisAkun);
		};

		fetchJenisAkun();
	}, []);

	const [asal, setAsal] = useState('');

	// Simulate fetching account type from a session or API
	useEffect(() => {
		// Example: Simulate fetching the account type from a session or API
		const fetchSessionAsal = () => {
		  // Simulate getting an account type (e.g., from a session)
		  const sessionAsal = 'DID'; // This would be dynamic in a real app
		  setAsal(sessionAsal);
		};

		fetchSessionAsal();
	}, []);

	const [statusFillingDisable, setStatusFillingDisable] = useState('');

	// Simulate fetching account type from a session or API
	useEffect(() => {
		// Example: Simulate fetching the account type from a session or API
		const fetchStatusFillingDisable = () => {
		  // Simulate getting an account type (e.g., from a session)
		  const sessionStatusFillingDisable = 'OFF'; // This would be dynamic in a real app
		  setStatusFillingDisable(sessionStatusFillingDisable);
		};

		fetchStatusFillingDisable();
	}, []);

	const [dataPenerima, setDataPenerima] = useState([]);

	useEffect(() => {
		// Fetch data from an API or define it directly
		const fetchDataPenerima = async () => {
		  // Simulate an API call
		  const response = await fetch('http://localhost:8080/api/testing_penerima'); // Replace with your API endpoint
		  const result = await response.json();
		  console.log(result["data"].penerima)
		  setDataPenerima(result["data"].penerima);
		};

		fetchDataPenerima();
	}, []);

	 const columns = [
	    {
	      name: 'No',
	      selector: (row, index) => index + 1,
	      center: true,
	      width: '50px',
	    },{
	      name: 'Action',
	      selector: (row) => row.id,
	      cell: (row) => {
	        const href = `/delete_data/${row.id}`; // URL for download
	        let btn = (
	          <button
	            className="btn btn-danger"
	            onClick={() => handleDeleteByID(row.id)}
	            data-target="#editModalById"
	            data-toggle="modal"
	          >
	            Delete
	          </button>
	        );

	        return btn;
	      },
	      center: true,
	      width: '120px',
	    },{
	      name: 'Type',
	      selector: (row, index) => row.type,
	      center: true,
	      width: '150px',
	    },{
	      name: 'Qty',
	      selector: (row, index) => row.qty,
	      center: true,
	      width: '70px',
	    },{
	      name: 'Alamat Pengirim',
	      selector: (row, index) => row.alamat_pengirim,
	      center: true,
	      width: '170px',
	    },{
	      name: 'PIC Pengirim',
	      selector: (row, index) => row.pic_pengirim,
	      center: true,
	      width: '150px',
	    },{
	      name: 'Alamat Penerima',
	      selector: (row, index) => row.alamat_penerima,
	      width: '400px',
	    },{
	      name: 'Warehouse Penerima',
	      selector: (row, index) => row.warehouse_penerima,
	      center: true,
	      width: '250px',
	    },{
	      name: 'PIC Penerima',
	      selector: (row, index) => row.pic_penerima,
	      center: true,
	      width: '200px',
	    },{
	      name: 'Tanggal Pengiriman',
	      selector: (row, index) => row.tanggal_pengiriman,
	      center: true,
	      width: '120px',
	    },{
	      name: 'Tanggal Sampai',
	      selector: (row, index) => row.tanggal_sampai,
	      center: true,
	      width: '120px',
	    },{
	      name: 'Batch',
	      selector: (row, index) => row.batch,
	      center: true,
	      width: '120px',
	    },{
	      name: 'Edit',
	      selector: (row) => row.id,
	      cell: (row) => {
	        const href = `/download_serial_number/${row.id}`; // URL for download
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
	      right: true,
	      center: true,
	      width: '220px',
	    }
	]

	return(
		<>
			<div className="wrapper d-flex align-items-stretch">
	        <Sidebar/>
	        <div id="content" style={{margin: "0 auto", boxSizing: "border-box"}}>
	            <div className="container-fluid" style={{width: "105.5%"}}>
					<Navbar/>

                    <div className="card mt-5">
                        <form className="mt-3" style={{paddingLeft: "15px", paddingRight: "15px", paddingBottom: "15px"}}
                            method="POST" action="" encType="multipart/form-data">
                            <div className="form-group row">
                                <label className="ml-4">Masukan file pengiriman untuk diupload:</label>
                            </div>
                            <div className="row mt-3">
                                <div className="col-sm-12 col-md-4 mb-3">
                                    <input className="form-control" type="file" style={{height: "45px"}}
                                        name="file_penerima" required></input>
                                </div>
                                <div className="col-sm-12 col-md-8">
                                    <div className="row">
                                        <div className="col-md-6 mb-2">
                                                <button type="submit" className="btn btn-primary">Upload</button>
                                            <a href="{{ URL('download_template_penerima') }}"
                                                className="btn btn-secondary ml-1">Download Template</a>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </form>
                    </div>

	                <div className="card mb-3 mt-3">
	                    {message && (
	                        <div className="alert alert-success alert-dismissible fade show ml-3 mr-3 mb-3 mt-5" role="alert">
	                            <strong>{message}</strong>
	                        </div>
	                    )}

	                    <div className="card-body mb-2">
	                        <div className="row">
	                            <div className="col-md-6 col-lg-5">
	                                <a className="btn btn-secondary mb-2" href="{{ URL('/export_penerima/ONT') }}">Export</a>
	                                <a className="btn btn-secondary mb-2 ml-1" href="{{ URL('/download_all_serial_number/ont/exist') }}">Download All SN Exist</a>
	                                <a className="btn btn-secondary mb-2 ml-1" href="{{ URL('/download_all_serial_number/ont/all') }}">Download All SN</a>
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
	                            ) : (
	                            	<div>
	                            	</div>
	                            )}
	                        </div>

	                        {asal === "DID" && ( 
	                        <div className="row mb-3">
	                            <div className="col-md-6 col-lg-5">
	                                <a className="btn btn-primary mb-2" data-toggle="modal" data-target="#editIdoGDBulk" onClick={() => handleUploadIDOGD()} >Upload IDO GD Bulk</a>
	                            </div>
	                        </div>
	                        )}

	                        <div className="row mb-1">
	                            <div className="col-md-6 col-lg-5">
	                            </div>
	                                <div className="col-md-6 col-lg-7 col-xl-7">
	                                    <div className="d-flex justify-content-end align-items-center">
	                                        <select id="batch-filter" name="native-select" placeholder="Batch"
	                                            data-search="true" data-silent-initial-value-set="true"
	                                            className="col-12 col-md-2 custom-select custom-select-sm form-control">
	                                        </select>
	                                    </div>
	                                </div>
	                        </div>

	                        <span>Last update: </span>
	                        <div className="table-responsive mt-2">
	                                <select id="TREGFilter"
	                                    className="col-12 col-md-3 me-2 custom-select custom-select-sm form-control">

	                                    <option value="" disabled selected>TREG:</option>
	                                    <option value="">All TREG</option>
	                                    <option value="WH TR TREG1">TREG 1</option>
	                                    <option value="WH TR TREG2">TREG 2</option>
	                                    <option value="WH TR TREG3">TREG 3</option>
	                                    <option value="WH TR TREG4">TREG 4</option>
	                                    <option value="WH TR TREG5">TREG 5</option>
	                                    <option value="WH TR TREG6">TREG 6</option>
	                                    <option value="WH TR TREG7">TREG 7</option>
	                                </select>
	                            {statusFillingDisable === "OFF" ? ( 
	                            <div className="table-responsive">
	                                <DataTable 
	                               		columns={columns}
									      data={dataPenerima}
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
		/>
		<DeleteAllModal
			showModal={showModalDeleteAll}
			handleCloseModal={handleCloseModalDeleteAll}
			handleDeleteAllClick={handleDeleteAllClick}
		/>
		<TambahModal
	   		showModal={showModalTambah}
			handleCloseModal={handleCloseModalTambah}
			handleTambahClick={handleTambahClick}
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