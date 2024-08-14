//import react router dom
import { Routes, Route } from "react-router-dom";

//import view homepage
import LandingPage from '../Views/LandingPage/LandingPage';
import HomeSCMT from '../Views/SCMT/HomeSCMT';
import PenerimaONT from '../Views/SCMT/PenerimaONT';
import Page1 from '../Views/Page_1/page_1';
import Home from '../Views/Home/Home';
import Login from '../Views/SCMT/Login';
import UploadFileDataStock from '../Views/SCMT/UploadFileDataStock';
import UploadDataPengiriman from '../Views/SCMT/UploadDataPengiriman';
import UploadDataDatabaseMinimumStock from '../Views/SCMT/UploadDataDatabaseMinimumStock';
import EditProfile from '../Views/SCMT/EditProfile';

function RoutesIndex() {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/landing_page" element={<LandingPage />} />
            <Route path="/page_1" element={<Page1 />} />
            <Route path="/scmt/login" element={<Login />} />
            <Route path="/scmt/rekap_delivery" element={<HomeSCMT />} />
            <Route path="/scmt/rekap_delivery/witel/:lokasi_wh" element={<HomeSCMT />} />
            <Route path="/scmt/report_delivery_ont" element={<PenerimaONT />} />
            <Route path="/scmt/upload_file_data_stock" element={<UploadFileDataStock />} />
            <Route path="/scmt/upload_file_pengiriman" element={<UploadDataPengiriman />} />
            <Route path="/scmt/upload_file_minimum_stock" element={<UploadDataDatabaseMinimumStock />} />
            <Route path="/scmt/edit_profile" element={<EditProfile />} />
        </Routes>
    )
}

export default RoutesIndex