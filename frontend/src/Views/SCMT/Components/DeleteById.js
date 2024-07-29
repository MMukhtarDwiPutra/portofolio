import React, {Component, useEffect, StrictMode, useState  }  from 'react';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

const DeleteByIdModal = ({ showModal, handleCloseModal, handleClick, data }) => {
	return (
	<>
		{data && (
	    	<Modal show={showModal} onHide={handleCloseModal} size="lg">
		      <Modal.Header closeButton>
		        <Modal.Title>Delete One Data</Modal.Title>
		      </Modal.Header>
		      <Modal.Body>
		      	<div class="" style={{fontSize: "14px"}}>Apakah anda yakin ingin menghapus data Pengiriman?<br/>
		      	Type : {data.type}<br/>
		      	Qty : {data.qty}<br/>
		      	PIC Pengirim : {data.pic_pengirim}<br/>
		      	Alamat Pengirim : {data.alamat_pengirim}<br/>
		      	PIC Penerima : {data.pic_penerima}<br/>
		      	Alamat Penerima : {data.alamat_penerima}<br/>
		      	Warehouse Penerima : {data.warehouse_penerima}<br/>
		      	Tanggal Pengiriman : {data.tanggal_pengiriman}<br/>
		      	Tanggal Sampai : {data.tanggal_sampai}<br/>
		      	Batch : {data.batch} <br/>
		      	</div>
		      </Modal.Body>
		      <Modal.Footer>
		        <Button variant="secondary" onClick={handleCloseModal}>
		          Cancel
		        </Button>
		        <Button variant="danger" onClick={handleClick}>
		          Delete
		        </Button>
		      </Modal.Footer>
		    </Modal>  
		)}
	</>
    )
}

export default DeleteByIdModal;