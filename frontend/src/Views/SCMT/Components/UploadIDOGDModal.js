import React, {Component, useEffect, StrictMode, useState  }  from 'react';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

const UploadIDOGDModal = ({ showModal, handleCloseModal, handleClick }) => {
  return (
    <>
    <Modal show={showModal} onHide={handleCloseModal} size="lg">
      <Modal.Header closeButton>
        <Modal.Title>Delete All Data Penerimaan</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <form name='form_edit_ido_gd_bulk' method='POST' accept-charset='UTF-8' enctype='multipart/form-data' action="{{URL('edit_ido_gd_bulk')}}" id='form_edit_ido_gd_bulk'>
            <div class="modal-body" id="modal-hapus-body">
                <div class="row mt-3">
                    <div class="col-5">
                        <label>Upload IDO GD Bulk</label>
                        <input type="file" class="form-control" name="file_edit_ido_gd_bulk"
                            id="file_edit_ido_gd_bulk" style={{height: "45px"}}
                            accept="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel"/>
                            <span id="span_time_sn_added"></span>
                    </div>
                </div>
            </div>
        </form>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleCloseModal}>
          Cancel
        </Button>
        <Button variant="primary" onClick={handleClick}>
          Upload IDO GD
        </Button>
      </Modal.Footer>
    </Modal>
    </>
  );
};

export default UploadIDOGDModal;