import React from 'react'
import ReactDOM from 'react-dom/client'
import './main.css'
import '@mantine/core/styles.css';
import { createTheme, MantineProvider } from '@mantine/core';
import App from './app';
import { BrowserRouter } from "react-router-dom";

const theme = createTheme({
});

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <MantineProvider theme={theme}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </MantineProvider>
  </React.StrictMode>,
)