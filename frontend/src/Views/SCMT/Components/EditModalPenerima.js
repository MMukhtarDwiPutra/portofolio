import React, {Component, useEffect, StrictMode, useState  }  from 'react';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import "../../../Assets/css/PopUp_Edit_Penerimaan.css"

const EditModal = ({ showModal, handleCloseModal, data, setCurrentItem, handleSaveChanges, setMessage, fetchDataPenerima }) => {
  const [showPopup1, setShowPopup1] = useState(false);
  const [showPopup2, setShowPopup2] = useState(false);

  const [dataForm, setDataForm] = useState({
    tanggal_pengiriman: data?.tanggal_pengiriman || "",
    tanggal_sampai: data?.tanggal_sampai || "",
    ido_gd: data?.ido_gd || "",
    sn_mac_barcode_file: data?.sn_mac_barcode || null,
  });

  const handleFileChange = (e) => {
    setDataForm({ ...dataForm, sn_mac_barcode_file: e.target.files[0] });
    console.log(dataForm)
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setDataForm((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  useEffect(() => {
    if (data) {
      // Update form data when the data prop changes
      setDataForm({
        tanggal_pengiriman: data?.tanggal_pengiriman || "",
        tanggal_sampai: data?.tanggal_sampai || "",
        ido_gd: data?.ido_gd || "",
        sn_mac_barcode_file: data?.sn_mac_barcode || null,
      });
    }
  }, [data]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const dataTmp = new FormData();
    dataTmp.append('tanggal_pengiriman', dataForm.tanggal_pengiriman);
    dataTmp.append('tanggal_sampai', dataForm.tanggal_sampai);
    dataTmp.append('ido_gd', dataForm.ido_gd);
    dataTmp.append('sn_mac_barcode_file', dataForm.sn_mac_barcode_file);

    try {
      // Log FormData for debugging
      for (let pair of dataTmp.entries()) {
        console.log(`${pair[0]}: ${pair[1]}`);
      }
      const response = await fetch(`http://localhost:8080/api/edit_on_delivery_by_id/ont/${data.id}`, {
        method: 'PUT',
        body: dataTmp
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      const result = await response.json();
      setMessage(result.data.message);
      handleCloseModal()
      handleShowPopup2()
      closeModalPopup1()
      fetchDataPenerima()
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const closeModalPopup1 = () => {
    setShowPopup1(false)
  }

  const handleShowPopup1 = () => {
    setShowPopup1(!showPopup1);
  };

  const handleShowPopup2 = () => {
     setShowPopup2(!showPopup2)
     setShowPopup1(!showPopup1) // Hide the first popup when showing the second
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

  const handleModalClose = () => {
    setShowPopup1(false); // Close popup1 when modal is closed
    setShowPopup2(false); // Optionally, close popup2 as well
    handleCloseModal(); // Call the parent component's close handler
  };

  return (
    <>
    {data && (
      <Modal show={showModal} onHide={handleModalClose} size="lg">
        <Modal.Header closeButton>
          <Modal.Title>Edit Data Penerimaan</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          {data && (
            <form onSubmit={handleSubmit} name='form_edit' method='POST' acceptCharset='UTF-8' encType='multipart/form-data' id='form_edit'>
                      <div className="modal-body" id="modal-hapus-body">
                          <div className="row">
                              <div className="col-4">
                                  <label>Tanggal Pengiriman</label>
                                   <input className="form-control" onChange={handleChange} defaultValue={data.tanggal_pengiriman} name="tanggal_pengiriman" id="tanggal_pengiriman" type="date" />
                              </div>
                              <div className="col-4">
                                  <label>Tanggal Sampai</label>
                                  <input className="form-control" defaultValue={data.tanggal_sampai} onChange={handleChange} name="tanggal_sampai" id="tanggal_sampai" type="date" />
                              </div>
                          </div>
                          <div className="row mt-3">
                              <div className="col-5">
                                  <label>Nomor IDO/GD</label>
                                  <input className="form-control" defaultValue={data.ido_gd} onChange={handleChange} name="ido_gd" id="ido_gd" type="text" />
                              </div>
                              <div className="col-5">
                                  <label>Upload SN; MAC Hasil Barcode</label>
                                  <input className="form-control" name="sn_mac_barcode_file" id="sn_mac_barcode_file" onChange={handleFileChange} type="file" />
                              </div>
                              <div className="col-2">
                                  <label>Template SN</label>
                                  <a className="btn btn-secondary" href="http://localhost:8080/api/download_template_serial_number_ont"
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
                                      <button type="submit" className="btn btn-primary" style={{height: "40px"}}>
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
      )}
    </>
  );
};

export default EditModal;