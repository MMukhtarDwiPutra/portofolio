import React, {Component, useEffect, StrictMode, useState  }  from 'react';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import "../../../Assets/css/PopUp_Edit_Penerimaan.css"
import "../../../Assets/css/ZF_Form.css"

const TambahModalPenerima = ({ showModal, handleCloseModal, handleSaveChanges, dataWarehouse, setMessage, fetchDataPenerima}) => {

  const [data, setData] = useState({
    type: '',
    qty: '',
    alamat_pengirim: '',
    pic_pengirim: '',
    alamat_penerima: '',
    warehouse_penerima: '',
    pic_penerima: '',
    tanggal_pengiriman: '',
    tanggal_sampai: '',
    batch: '',
  });

  useEffect(() => {
    
  }, [data]); // Dependency array with 'data' ensures the effect runs on data changes


  const handleChange = (e) => {
    const { name, value } = e.target;
    setData((prevData) => ({
      ...prevData,
      [name]: value,
    }));

    console.log("tes")
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8080/api/tambah_penerima', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }else{
          const responseData = await response.json();
          console.log('Success:', responseData);
          setMessage("Success Add Data Penerimaan");
          handleCloseModal()
          await fetchDataPenerima()
      }

    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <>
    <Modal show={showModal} onHide={handleCloseModal} size="lg">
      <form onSubmit={handleSubmit}>
      <Modal.Header closeButton>
        <Modal.Title>Delete All Data Penerimaan</Modal.Title>
      </Modal.Header>
      <Modal.Body>
          <div className="modal-body" id="isi_qr">
                        <div className="zf-templateWrapper">
                            <ul className="zf-tempHeadBdr">
                                <li className="zf-tempHeadContBdr">
                                    <h2 className="zf-frmTitle"><em>Tambah Pengiriman</em></h2>
                                    <p className="zf-frmDesc"></p>
                                    <div className="zf-clearBoth"></div>
                                </li>
                            </ul>
                            <table>
                                <tbody>
                                <tr style={{borderBottom : "1px solid #ddd"}}>
                                    <td>
                                        <ul>
                                            <li className="zf-tempFrmWrapper zf-small">
                                                <label className="zf-labelName" style={{color: "black"}}>Type</label>
                                                <div className="zf-tempContDiv zf-twoType">
                                                    <select onChange={handleChange} name="type" style={{border: "1px solid black"}}
                                                        className=" zf-form-sBox col-12 col-md-7 me-2 custom-select custom-select-sm form-control"
                                                        name="type" checktype="c1" required>
                                                        <option value="">-Pilih Type-
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
                                            <li className="zf-tempFrmWrapper zf-name zf-namemedium">
                                              <div className="zf-tempContDiv zf-twoType">
                                                <div className="zf-nameWrapper">
                                                  <span>
                                                    <label className="zf-labelName" style={{color: "black"}}>Jumlah</label>
                                                  </span>
                                                  <input style={{border: "1px solid black"}}
                                                    type="number" onChange={handleChange} name="qty" maxLength="18"
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
                                            <li className="zf-tempFrmWrapper zf-small">
                                                <label className="zf-labelName" style={{color: "black"}}>Pengirim: </label>
                                            </li>
                                        </ul>
                                    </td>
                                    <td>
                                        <ul>
                                            <li className="zf-tempFrmWrapper zf-name zf-namemedium">
                                                <div className="zf-tempContDiv zf-twoType">
                                                    <div className="zf-nameWrapper">
                                                        <span>
                                                            <label className="zf-labelName" style={{color: "black"}}>PIC</label>
                                                        </span>
                                                        <input onChange={handleChange} style={{border: "1px solid black"}} type="text"
                                                            maxLength="255" fieldtype="7" placeholder=""
                                                            name="pic_pengirim" required />
                                                    </div>
                                                </div>
                                                <div className="zf-clearBoth"></div>
                                            </li>
                                        </ul>
                                    </td>
                                </tr>
                                <tr style={{borderBottom: "1px solid #ddd"}}>
                                    <td></td>
                                    <td>
                                        <li className="zf-tempFrmWrapper zf-name zf-namemedium">
                                            <div className="zf-tempContDiv zf-twoType">
                                                <div className="zf-nameWrapper">
                                                    <span>
                                                        <label className="zf-labelName" style={{color: "black"}}>Alamat</label>
                                                        <textarea name="alamat_pengirim" onChange={handleChange} style={{border: "1px solid black"}} checktype="c1" maxLength="65535" fieldtype="7" name="alamat_pengirim"
                                                            placeholder="" required></textarea>
                                                    </span>
                                                </div>
                                            </div>
                                            <div className="zf-clearBoth"></div>
                                        </li>
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <ul>
                                            <li className="zf-tempFrmWrapper zf-name zf-namemedium">                                                
                                                <label className="zf-labelName" style={{color: "black"}}>Penerima: </label>
                                            </li>
                                        </ul>
                                    </td>
                                    <td>
                                        <li className="zf-tempFrmWrapper zf-name zf-namemedium">
                                            <div className="zf-tempContDiv zf-twoType">
                                                <div className="zf-nameWrapper">
                                                    <span>
                                                        <label className="zf-labelName" style={{color: "black"}}>PIC</label>
                                                        <input style={{border: "1px solid black"}} type="text"
                                                            maxLength="255" name="pic_penerima" fieldtype="7"
                                                            placeholder="" onChange={handleChange} required />
                                                    </span>
                                                </div>
                                            </div>
                                            <div className="zf-clearBoth"></div>
                                        </li>
                                    </td>
                                </tr>
                                <tr>
                                    <td></td>
                                    <td style={{paddingLeft: "20px"}}>
                                        <ul>
                                            <li className="zf-tempFrmWrapper zf-name zf-namemedium">
                                                <label className="zf-labelName" style={{color: "black"}}>Warehouse</label>
                                            </li>
                                        </ul>
                                        <select style={{border: "1px solid black"}} className="zf-form-sBox col-12 col-md-8 me-5 custom-select custom-select-sm form-control"
                                            name="warehouse_penerima" onChange={handleChange} checktype="c1" required>
                                            <option value="">-Pilih Warehouse-</option>
                                            {dataWarehouse && dataWarehouse.map((item, index) => (
                                                <option key={index} value={item.lokasi_wh}>{item.lokasi_wh}</option>
                                            ))}
                                        </select>
                                    </td>
                                </tr>
                                <tr style={{borderBottom: "1px solid #ddd"}}>
                                    <td></td>
                                    <td>
                                        <li className="zf-tempFrmWrapper zf-name zf-namemedium">
                                            <div className="zf-tempContDiv zf-twoType">
                                                <div className="zf-nameWrapper">
                                                    <span>
                                                        <label className="zf-labelName" style={{color: "black"}}>Alamat</label>
                                                        <textarea onChange={handleChange} name="alamat_penerima" style={{border: "1px solid black"}} checktype="c1" maxLength="65535" name="alamat_penerima" fieldtype="7"
                                                            placeholder="" required></textarea>
                                                    </span>
                                                </div>
                                            </div>
                                            <div className="zf-clearBoth"></div>
                                        </li>
                                    </td>
                                </tr>
                                <tr style={{borderBottom : "1px solid #ddd"}}>
                                    <td>
                                        <li className="zf-tempFrmWrapper zf-name zf-namemedium">
                                            <label className="zf-labelName" style={{color: "black"}}>Waktu
                                            Pengiriman:</label>
                                        </li>
                                    </td>
                                    <td>
                                        <div className="zf-subContWrap zf-topAlign">
                                            <li className="zf-tempFrmWrapper zf-name zf-namemedium">
                                                <div className="zf-tempContDiv zf-twoType">
                                                    <div className="zf-nameWrapper">

                                                        <span>
                                                            <label className="zf-labelName" style={{color: "black"}}>Tanggal Kirim</label>
                                                            <input onChange={handleChange} type="date" maxLength="255"
                                                                name="tanggal_pengiriman" fieldtype="7"
                                                                placeholder="" />
                                                        </span>
                                                        <span className="ml-1">
                                                            <label className="zf-labelName" style={{color: "black"}}>Tanggal Sampai</label>
                                                            <input onChange={handleChange} type="date" maxLength="255"
                                                                name="tanggal_sampai" fieldtype="7" placeholder="" />
                                                        </span>
                                                    </div>
                                                </div>
                                                <div className="zf-clearBoth"></div>
                                            </li>
                                        </div>
                                    </td>
                                </tr>
                                <tr >
                                    <td>
                                        <ul>
                                            <li className="zf-tempFrmWrapper zf-small">
                                                <label className="zf-labelName" style={{color: "black"}}>Batch</label>
                                                <div className="zf-tempContDiv zf-twoType">
                                                    <input type="text" name="batch" onChange={handleChange}/>
                                                </div>
                                            </li>
                                        </ul>
                                    </td>
                                </tr>
                                </tbody>
                            </table>
                        </div>
                        </div>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleCloseModal}>
          Cancel
        </Button>
        <Button variant="primary" type="submit">
          Tambah Data
        </Button>
      </Modal.Footer>
      </form>
    </Modal>
    </>
  );
};

export default TambahModalPenerima;