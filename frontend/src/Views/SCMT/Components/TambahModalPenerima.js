import React, {Component, useEffect, StrictMode, useState  }  from 'react';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import "../../../Assets/css/Form_Tambah_Penerima.css"
import "../../../Assets/css/ZF_Form.css"

const TambahModalPenerima = ({ showModal, handleCloseModal, handleSaveChanges }) => {
  return (
    <>
    <Modal show={showModal} onHide={handleCloseModal} size="lg">
      <Modal.Header closeButton>
        <Modal.Title>Delete All Data Penerimaan</Modal.Title>
      </Modal.Header>
      <Modal.Body>
          <div class="modal-body" id="isi_qr">
                        <div class="zf-templateWrapper">
                            <ul class="zf-tempHeadBdr">
                                <li class="zf-tempHeadContBdr">
                                    <h2 class="zf-frmTitle"><em>Tambah Pengiriman</em></h2>
                                    <p class="zf-frmDesc"></p>
                                    <div class="zf-clearBoth"></div>
                                </li>
                            </ul>
                            <table>
                                <tr style={{borderBottom : "1px solid #ddd"}}>
                                    <td>
                                        <ul>
                                            <li class="zf-tempFrmWrapper zf-small">
                                                <label class="zf-labelName" style={{color: "black"}}>Type</label>
                                                <div class="zf-tempContDiv zf-twoType">
                                                    <select style={{border: "1px solid black"}}
                                                        class=" zf-form-sBox col-12 col-md-7 me-2 custom-select custom-select-sm form-control"
                                                        name="type" checktype="c1" required>
                                                        <option selected="true" value="">-Pilih Type-
                                                        </option>
                                                        <option value="ONT_ZTE_F670L">ONT_ZTE_F670L</option>
                                                        <option value="ONT_NOKIA_G240WL">ONT_NOKIA_G240WL</option>
                                                        <option value="ONT_NOKIA_G-2425G-A">ONT_NOKIA_G-2425G-A
                                                        </option>
                                                        <option value="ONT_FIBERHOME_HG6145D2">
                                                            ONT_FIBERHOME_HG6145D2</option>
                                                        <option value="ONT_FIBERHOME_HG6145F">ONT_FIBERHOME_HG6145F
                                                        </option>
                                                        <option value="ONT_HUAWEI_HG8145V5">ONT_HUAWEI_HG8145V5
                                                        </option>
                                                        <option value="ONT_ZTE_F670 V2.0">ONT_ZTE_F670 V2.0
                                                        </option>
                                                        <option value="ONT_ZTE_F670">ONT_ZTE_F670</option>
                                                        <option value="ONT_FIBERHOME_HG6245N">ONT_FIBERHOME_HG6245N
                                                        </option>
                                                        <option value="ONT_HUAWEI HG8245W5-6T">ONT_HUAWEI
                                                            HG8245W5-6T</option>
                                                        <option value="ONT_HW_HG8245W5-6T">ONT_HW_HG8245W5-6T
                                                        </option>
                                                    </select>
                                                </div>
                                            </li>
                                        </ul>
                                    </td>
                                    <td>
                                        <ul>
                                            <li class="zf-tempFrmWrapper zf-name zf-namemedium">
                                              <div class="zf-tempContDiv zf-twoType">
                                                <div class="zf-nameWrapper">
                                                  <span>
                                                    <label class="zf-labelName" style={{color: "black"}}>Jumlah</label>
                                                  </span>
                                                  <input style={{border: "1px solid black"}}
                                                    type="number" name="qty" value="" maxlength="18"
                                                    placeholder="" required />
                                                  </div>
                                                </div>
                                            </li>
                                        </ul>
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <ul>
                                            <li class="zf-tempFrmWrapper zf-small">
                                                <label class="zf-labelName" style={{color: "black"}}>Pengirim: </label>
                                            </li>
                                        </ul>
                                    </td>
                                    <td>
                                        <ul>
                                            <li class="zf-tempFrmWrapper zf-name zf-namemedium">
                                                <div class="zf-tempContDiv zf-twoType">
                                                    <div class="zf-nameWrapper">
                                                        <span>
                                                            <label class="zf-labelName" style={{color: "black"}}>PIC</label>
                                                        </span>
                                                        <input style={{border: "1px solid black"}} type="text"
                                                            maxlength="255" fieldType="7" placeholder=""
                                                            name="pic_pengirim" required />
                                                    </div>
                                                </div>
                                                <div class="zf-clearBoth"></div>
                                            </li>
                                        </ul>
                                    </td>
                                </tr>
                                <tr style={{borderBottom: "1px solid #ddd"}}>
                                    <td></td>
                                    <td>
                                        <li class="zf-tempFrmWrapper zf-name zf-namemedium">
                                            <div class="zf-tempContDiv zf-twoType">
                                                <div class="zf-nameWrapper">
                                                    <span>
                                                        <label class="zf-labelName" style={{color: "black"}}>Alamat</label>
                                                        <textarea style={{border: "1px solid black"}} checktype="c1" maxlength="65535" fieldType="7" name="alamat_pengirim"
                                                            placeholder="" required></textarea>
                                                    </span>
                                                </div>
                                            </div>
                                            <div class="zf-clearBoth"></div>
                                        </li>
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <ul>
                                            <li class="zf-tempFrmWrapper zf-name zf-namemedium">                                                
                                                <label class="zf-labelName" style={{color: "black"}}>Penerima: </label>
                                            </li>
                                        </ul>
                                    </td>
                                    <td>
                                        <li class="zf-tempFrmWrapper zf-name zf-namemedium">
                                            <div class="zf-tempContDiv zf-twoType">
                                                <div class="zf-nameWrapper">
                                                    <span>
                                                        <label class="zf-labelName" style={{color: "black"}}>PIC</label>
                                                        <input style={{border: "1px solid black"}} type="text"
                                                            maxlength="255" name="pic_penerima" fieldType="7"
                                                            placeholder="" required />
                                                    </span>
                                                </div>
                                            </div>
                                            <div class="zf-clearBoth"></div>
                                        </li>
                                    </td>
                                </tr>
                                <tr>
                                    <td></td>
                                    <td style={{paddingLeft: "20px"}}>
                                        <ul>
                                            <li class="zf-tempFrmWrapper zf-name zf-namemedium">
                                                <label class="zf-labelName" style={{color: "black"}}>Warehouse</label>
                                            </li>
                                        </ul>
                                        <select style={{border: "1px solid black"}}>
                                            class="zf-form-sBox col-12 col-md-8 me-5 custom-select custom-select-sm form-control"
                                            name="warehouse_penerima" checktype="c1" required>
                                            <option selected="true" value="">-Pilih Warehouse-
                                            </option>
                                        </select>
                                    </td>
                                </tr>
                                <tr style={{borderBottom: "1px solid #ddd"}}>
                                    <td></td>
                                    <td>
                                        <li class="zf-tempFrmWrapper zf-name zf-namemedium">
                                            <div class="zf-tempContDiv zf-twoType">
                                                <div class="zf-nameWrapper">
                                                    <span>
                                                        <label class="zf-labelName" style={{color: "black"}}>Alamat</label>
                                                        <textarea style={{border: "1px solid black"}} checktype="c1" maxlength="65535" name="alamat_penerima" fieldType="7"
                                                            placeholder="" required></textarea>
                                                    </span>
                                                </div>
                                            </div>
                                            <div class="zf-clearBoth"></div>
                                        </li>
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <li class="zf-tempFrmWrapper zf-name zf-namemedium">
                                            <label class="zf-labelName" style={{color: "black"}}>Waktu
                                            Pengiriman:</label>
                                        </li>
                                    </td>
                                    <td>
                                        <div class="zf-subContWrap zf-topAlign">
                                            <li class="zf-tempFrmWrapper zf-name zf-namemedium">
                                                <div class="zf-tempContDiv zf-twoType">
                                                    <div class="zf-nameWrapper">

                                                        <span>
                                                            <label class="zf-labelName" style={{color: "black"}}>Tanggal Kirim</label>
                                                            <input type="date" maxlength="255"
                                                                name="tanggal_pengiriman" fieldType="7"
                                                                placeholder="" />
                                                        </span>
                                                        <span>
                                                            <label class="zf-labelName" style={{color: "black"}}>Tanggal Sampai</label>
                                                            <input type="date" maxlength="255"
                                                                name="tanggal_sampai" fieldType="7" placeholder="" />
                                                        </span>
                                                    </div>
                                                </div>
                                                <div class="zf-clearBoth"></div>
                                            </li>
                                        </div>
                                    </td>
                                </tr>
                            </table>
                        </div>
                        </div>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleCloseModal}>
          Cancel
        </Button>
        <Button variant="primary">
          Tambah Data
        </Button>
      </Modal.Footer>
    </Modal>
    </>
  );
};

export default TambahModalPenerima;