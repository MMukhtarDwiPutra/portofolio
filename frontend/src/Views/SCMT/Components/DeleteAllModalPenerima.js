import React, {Component, useEffect, StrictMode, useState  }  from 'react';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import "../../../Assets/css/PopUp_Edit_Penerimaan.css"

const DeleteAllModal = ({ showModal, handleCloseModal, handleClick }) => {
  return (
    <>
    <Modal show={showModal} onHide={handleCloseModal} size="lg">
      <Modal.Header closeButton>
        <Modal.Title>Delete All Data Penerimaan</Modal.Title>
      </Modal.Header>
      <Modal.Body>
          <div class="modal-body">Apakah anda yakin ingin menghapus semua data Pengiriman?</div>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleCloseModal}>
          Cancel
        </Button>
        <Button variant="danger" onClick={handleClick}>
          Hapus
        </Button>
      </Modal.Footer>
    </Modal>
    </>
  );
};

export default DeleteAllModal;