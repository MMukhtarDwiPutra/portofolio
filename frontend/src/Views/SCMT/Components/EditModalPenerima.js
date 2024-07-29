import React, {Component, useEffect, StrictMode, useState  }  from 'react';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import "../../../Assets/css/PopUp_Edit_Penerimaan.css"

const EditModal = ({ showModal, handleCloseModal, data, setCurrentItem, handleSaveChanges }) => {
  const [showPopup1, setShowPopup1] = useState(false);
  const [showPopup2, setShowPopup2] = useState(false);

  const handleShowPopup1 = () => {
    setShowPopup1(!showPopup1);
  };

  const handleShowPopup2 = () => {
     setShowPopup2(!showPopup2)
     setShowPopup1(!showPopup1); // Hide the first popup when showing the second
  };

  const defaultStylePopUp1 = {
    display: "none",
  };

  const clickedStylePopUp1 = {
    display: "block" // Spread default styles
  };

  const defaultStylePopUp2 = {
    display: "none",
  };

  const clickedStylePopUp2 = {
    display: "block" // Spread default styles
  };

  return (
    <>
    <Modal show={showModal} onHide={handleCloseModal} size="lg">
      <Modal.Header closeButton>
        <Modal.Title>Edit Data Penerimaan</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        {data && (
          <form name='form_edit' method='POST' acceptCharset='UTF-8' encType='multipart/form-data' id='form_edit'>
                    <div className="modal-body" id="modal-hapus-body">
                        <div className="row">
                            <div className="col-4">
                                <label>Tanggal Pengiriman</label>
                                 <input className="form-control" value={data.tanggal_pengiriman} name="tanggal_pengiriman" id="tanggal_pengiriman" type="date" />
                            </div>
                            <div className="col-4">
                                <label>Tanggal Sampai</label>
                                <input className="form-control" value={data.tanggal_sampai} name="tanggal_sampai" id="tanggal_sampai" type="date" />
                            </div>
                        </div>
                        <div className="row mt-3">
                            <div className="col-5">
                                <label>Nomor IDO/GD</label>
                                <input className="form-control" value={data.ido_gd} name="nomor_ido_gd" id="nomor_ido_gd" type="text" />
                            </div>
                            <div className="col-5">
                                <label>Upload SN; MAC Hasil Barcode</label>
                                <input className="form-control" name="edit_sn_mac_barcode" id="edit_sn_mac_barcode" type="file" />
                            </div>
                            <div className="col-2">
                                <label>Template SN</label>
                                <a className="btn btn-secondary" href="{{ URL('/download_template_sn/ont') }}"
                                    value="false">Template</a>
                            </div>
                        </div>
                        {showPopup1 && (
                          <div className="popup_box" style={showPopup1 ? clickedStylePopUp1 : defaultStylePopUp1}>
                              <div className="row">
                                <div className="col-1">
                                  <i className="fas fa-exclamation"></i>
                                </div>
                                <div className="col-11">
                                  <label>Anda yakin SN;MAC yang diupload TYPE XXXXX sejumlah XX ke WH
                                      XXX ???</label>
                                </div>
                              </div>
                              <div className="row d-flex justify-content-end mt-3">
                                <div className="col-auto">
                                    <a href="#" className="btn2 btn btn-primary" onClick={handleShowPopup2}>Yakin</a>
                                    <a href="#" className="btn1 btn btn-secondary ml-1" onClick={handleShowPopup1}>Tidak Yakin</a>
                                </div>
                              </div>
                          </div>
                        )}

                        {showPopup2 && (
                          <div className="popup_box2" style={showPopup2 ? clickedStylePopUp2 : defaultStylePopUp2}>
                              <div className="row">
                                <div className="col-1">
                                  <i className="fas fa-exclamation"></i>
                                </div>
                                <div className="col-11">
                                  <label>Semua SN;MAC yang diupload jadi acuan di SCMT, jadi jangan salah upload.</label>
                                </div>
                              </div>
                            
                              <div className="row d-flex justify-content-end mt-3">
                                <div className="col-auto">
                                    <button type="submit" className="btn btn-primary" style={{height: "40px"}} onClick={handleSaveChanges}>
                                        Teruskan Upload
                                    </button>
                                    <a href="#" className="btn3 btn btn-secondary ml-1" onClick={handleShowPopup2}>Batalkan Upload</a>
                                </div>
                              </div>
                        </div>
                        )}
                    </div>
                </form>
        )}
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleCloseModal}>
          Close
        </Button>
        <Button variant="primary click" onClick={handleShowPopup1}>
          Save Changes
        </Button>
      </Modal.Footer>
    </Modal>
    </>
  );
};

export default EditModal;