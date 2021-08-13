import './App.scss';
import AppContainer from './components/AppContainer/AppContainer.component';
import React from 'react';
import HomePage from './components/HomePage/HomePage.component';
import { Route, BrowserRouter as Router } from 'react-router-dom';
import DetailPageContainer from './components/DetailPageContainer/DetailPageContainer.component';
function App() {
  return (
    <>
    <AppContainer title="MicroFinance Helper"></AppContainer>
      <Router>  
          <Route exact path="/" component={HomePage} />
          <Route path="/home" component={HomePage} /> 
          <Route path="/detail" component={DetailPageContainer}></Route>
      </Router>
    </>
  );
}  

export default App;
