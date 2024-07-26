import React from 'react';
import Container from 'react-bootstrap/Container';
import Image from 'react-bootstrap/Image';

import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import image1 from '../../Assets/img/intelikwan.png'
import image2 from '../../Assets/img/SCMT.jpeg'

export default function Home(){
	return (
        <Container>
          <Row>
            <Col xs={6} md={4} className="text-center mt-5">
	            <a href="landing_page">
	            	<Image src={image1} fluid/>
	            </a>
	            <span className="mt-3"><b> Aidikei </b></span>
            </Col>
            <Col xs={6} md={4} className="text-center mt-5">
            	<a href="scmt">
	            	<Image src={image2} fluid />
	            </a>
	            <span className="mt-3"><b> SCMT </b></span>
	        </Col>
            <Col className="text-center mt-5">1 of 1</Col>
          </Row>
        </Container>
    )
}