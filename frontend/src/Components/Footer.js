import React from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';

function FooterPage(){
	return(
		<footer className="py-3 my-4 pt-5">
	    <ul className="nav justify-content-center border-bottom pb-3 mb-3">
	      <li className="nav-item"><a href="#" className="nav-link px-2 text-body-secondary">Home</a></li>
	      <li className="nav-item"><a href="#" className="nav-link px-2 text-body-secondary">Features</a></li>
	      <li className="nav-item"><a href="#" className="nav-link px-2 text-body-secondary">Pricing</a></li>
	      <li className="nav-item"><a href="#" className="nav-link px-2 text-body-secondary">FAQs</a></li>
	      <li className="nav-item"><a href="#" className="nav-link px-2 text-body-secondary">About</a></li>
	    </ul>
	    <p className="text-center text-body-secondary">© 2024 Company, Inc</p>
	  </footer>
	)
}

export default FooterPage;