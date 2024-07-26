//import react router dom
import { Routes, Route } from "react-router-dom";

//import view homepage
import LandingPage from '../Views/LandingPage/LandingPage';
import HomeSCMT from '../Views/SCMT/HomeSCMT';
import Page1 from '../Views/Page_1/page_1';
import Home from '../Views/Home/Home';

function RoutesIndex() {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/landing_page" element={<LandingPage />} />
            <Route path="/page_1" element={<Page1 />} />
            <Route path="/scmt" element={<HomeSCMT />} />
        </Routes>
    )
}

export default RoutesIndex