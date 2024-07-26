import React from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Container from 'react-bootstrap/Container';

function HeaderPage(){
  return(      
    <Navbar className="bg-dark">
      <Container>
        <Navbar.Brand href="/" className="text-light" style={{fontSize: "20px", fontWeight: "bold", display: "block", color: "white"}}>Portofolio</Navbar.Brand>
        <Navbar.Toggle />
      </Container>
    </Navbar>
  )
}

export default HeaderPage;