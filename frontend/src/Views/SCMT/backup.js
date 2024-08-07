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












<li key="1" className={collapseActive[1] ? 'active' : ''}>
    <a href="#" data-toggle="collapse" onClick={() => toggleCollapse(1)} aria-expanded="false" className="dropdown-toggle">Minimum Stock STB</a>
    <ul className={collapseActive[1] ? 'collapse list-unstyled components' : 'list-unstyled components'} id="SubMenu1">
        <li>
            <a className="nav-link active" href="{{url('/stb/rekap_delivery_stb')}}">Rekap Minimum Stock STB</a>
        </li>
        <li>
            <a className="nav-link active" href="{{url('/pengiriman_stb')}}">Report Delivery STB</a>
        </li>
    </ul>
</li>

<li key="2" className="active">
    <a href="#" data-toggle="collapse" aria-expanded="false" onClick={() => toggleCollapse(2)} className="dropdown-toggle">Minimum Stock AP</a>
    <ul className={collapseActive[2] ? 'collapse list-unstyled components' : 'list-unstyled components'} id="SubMenu2">
        <li>
            <a className="nav-link active" href="{{url('/rekap_delivery_ap')}}">Rekap Minimum Stock AP</a>
        </li>
        <li>
            <a className="nav-link active" href="{{url('/pengiriman_ap')}}">Report Delivery AP</a>
        </li>
    </ul>
</li>

<li>
	<a className="nav-link active" href="{{url('/input_sn_mac_vendor')}}">Upload File SN Vendor</a>
</li>

<li>
    <a className="nav-link active" style={{color:"#fff"}} href="{{url('/request_outbond')}}">Request Outbond</a>
</li>