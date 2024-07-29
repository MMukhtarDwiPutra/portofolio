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
	                                <table className="table table-bordered w-100" id="dataTable-pengiriman">
	                                    <thead className="text-center">
	                                        <tr>
	                                            <th rowSpan="3" className="text-center align-middle">No</th>
	                                            {jenisAkun === "Admin" && ( 
	                                                <th rowSpan="3" className="text-center align-middle">Action</th>
	                                            )}
	                                            <th rowSpan="2" className="align-middle text-center">Type</th>
	                                            <th rowSpan="2" className="align-middle text-center">Qty</th>
	                                            <th colSpan="2" className="text-center">Pengirim</th>
	                                            <th colSpan="3" className="text-center">Penerima</th>
	                                            <th rowSpan="2" className="align-middle text-center">Tanggal Pengiriman</th>
	                                            <th rowSpan="2" className="align-middle text-center">Tanggal Sampai</th>
	                                            <th rowSpan="2" className="align-middle text-center">Batch</th>
	                                            <th rowSpan="2" className="align-middle text-center">Edit</th>
	                                        </tr>
	                                        <tr>
	                                            <th className="text-center">Alamat</th>
	                                            <th className="text-center">PIC</th>
	                                            <th className="text-center">Alamat</th>
	                                            <th className="text-center" style={{width: "180px"}}>Warehouse</th>
	                                            <th className="text-center">PIC</th>
	                                            <th className="text-center" hidden>regional</th>
	                                        </tr>
	                                    </thead>
	                                    <tbody>
                                        {dataPenerima.map((item, index) => (
                                              <tr key={index}>
                                                <td>{index+1}</td>
                                                <td>Action</td>
                                                <td>{item.type}</td>
                                                <td>{item.qty}</td>
                                                <td>{item.alamat_pengirim}</td>
                                                <td>{item.pic_pengirim}</td>
                                                <td>{item.alamat_penerima}</td>
                                                <td>{item.warehouse_penerima}</td>
                                                <td>{item.pic_penerima}</td>
                                                <td>{item.tanggal_pengiriman}</td>
                                                <td>{item.tanggal_sampai}</td>
                                                <td>{item.batch}</td>
                                                <td>Edit</td>
                                              </tr>
                                            ))}
	                                    </tbody>
	                                </table>
	                            </div>
	                            ) : (
	                            <div className="text-center mt-4" style={{backgroundColor: "gray", fontSize: "20px"}}>
	                                <span style={{color:"white"}}>Data report delivery ONT sedang dimaintance, mohon menunggu.</span>
	                            </div>
	                            )}
	                        </div>