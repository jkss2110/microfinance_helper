import './App.scss';
import AppContainer from './components/AppContainer/AppContainer.component';
import React from 'react';
import HomePage from './components/HomePage/HomePage.component';

function App() {

  return (
    <>
    <AppContainer title="MicroFinance Helper"></AppContainer>
    <HomePage></HomePage>
    </>
  );
}

export default App;
