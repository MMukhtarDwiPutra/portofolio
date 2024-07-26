import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { Button} from 'react-bootstrap';
import image1 from '../../Assets/img/intelikwan.png'
import image2 from '../../Assets/img/CTA.png'

export default function LandingPage() {
    return (
        <Container>
          <Row>
              <Col className="text-center">
                <h5>Selamat Datang di Era Digital!</h5>
                <a>Kami hadir untuk bikin belajar teknologi jadi lebih gampang dan seru.<br />
                Bisa paham AI dan dunia digital itu seru, kayak ngobrol santai. <br />
                Gak pake ribet, kita bantu kamu kenal lebih dekat dengan dunia digital yang asik. <br /> <br />
                Yuk, mulai petualangan digitalmu dengan santai!
                </a>
              </Col>
          </Row>
          <Row className="justify-content-md-center">
              <Col className="col-4 text-center pt-3">
                <a href="#dapatkan_sekarang" class="form-control" style={{background: "#F4A261", fontSize:"24px", paddingBottom:"10px"}}>
                  <b> Dapatkan Sekarang! </b>
                </a>
              </Col>
          </Row>
          <Row className="justify-content-md-center">
              <Col className="col-4 text-center pt-3">
                <img src={image1} className="img-fluid"></img>  
              </Col>
          </Row>
          <Row>
              <Col className="text-center mt-5">
                <h5>Era Digital, Bukan Cuma Seru tapi Juga Penuh Tantangan!</h5>
                <a>AI lagi hits banget, bikin kepo tapi bingung mulai dari mana? Istilah teknis bikin pusing? Tenang, kamu gak sendirian. Kita ada solusi buat kamu yang mau ngerti AI tanpa ribet!Temukan cara seru dan mudah untuk memahami dasar-dasarnya tanpa pusing!
                </a>
              </Col>
          </Row>

          <Row>
              <Col className="col-4 mt-5">
                <h5 className="text-center">Kenapa Intelikwan Bisa Jadi Pilihanmu?</h5>
                <a>✅Mudah Dipahami: Dapatkan penjelasan AI yang ringan dan mudah dicerna, cocok untuk pemula. <br/>
                    ✅Bahasa Santai: Nikmati pembelajaran dengan gaya bahasa yang santai, membuat proses belajar serasa ngobrol dengan teman. <br/>
                    ✅Gerbang ke Era Digital: Langkah awal yang sempurna untuk siapa saja yang ingin mengerti dan menguasai dunia digital yang terus berkembang.
                </a>
              </Col>

              <Col className="col-4 mt-5">
                <h5 className="text-center">Siap untuk Upgrade Keahlian Digitalmu?</h5>
                <a>Investasikan Rp350.000 dan buka lembaran baru petualanganmu di dunia AI bersama INTELIKWAN—panduan santai buat kamu yang mulai dari dasar. Ebook ini dirancang khusus buat yang pengen kenal AI dari dasar banget.
                </a>
              </Col>

              <Col className="col-4 mt-5">
                <h5 className="text-center">Klaim Potongan Hargamu, Jangan Sampai Kehabisan!</h5>
                <a>Penawaran Spesial! Investasi Rp109.000 untuk kamu yang siap menjelajah dunia AI bersama INTELIKWAN. Jangan lewatkan kesempatan emas ini, tekan 'Dapatkan Sekarang' dan mulailah transformasi menjadi pribadi yang baru. Investasi pintar untuk masa depan lebih baik dimulai sekarang!
                </a>
              </Col>
          </Row>

          <Row className="justify-content-md-center mt-5">
              <Col className="col-6 text-center pt-3">
                <Row>
                  <img src={image2} className="img-fluid"></img>
                </Row>
                <Row className="mt-5" style={{fontSize:"20px"}}>
                  <div className="text-center">
                      <h6>Berapa Investasinya?</h6>
                      <h6 style={{color:"black"}}><s>Rp 350.000</s></h6>
                      <h1>Rp 109.000</h1>
                  </div>
                  <a href="#dapatkan_sekarang" class="form-control" style={{background: "#F4A261", fontSize:"24px", paddingBottom:"10px"}}>
                  <b> Dapatkan Sekarang! </b>
                </a>
                </Row>
              </Col>
          </Row>
        </Container>
    )
}

