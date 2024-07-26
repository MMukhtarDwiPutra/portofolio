import logo from './logo.svg';
import HeaderPage from "./Components/Header";
import FooterPage from "./Components/Footer";
import 'bootstrap/dist/css/bootstrap.min.css';  
import Routes from './Routes/Route';

function App() {
  return (
    <>
    <div className="App">
      <header className="App-header">
        <HeaderPage/>
        <Routes />
      </header>
      <FooterPage/>
    </div>
    </>
  );
}

export default App;
