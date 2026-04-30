import '@fontsource/poppins/400.css';
import '@fontsource/poppins/500.css';
import { StrictMode } from 'react';
import * as ReactDOM from 'react-dom/client';
import { MemoryRouter } from 'react-router-dom';
import App from './app/app';
import './styles.scss';

const container = document.getElementById('root');

if (container) {
  const root = ReactDOM.createRoot(container);
  root.render(
    <StrictMode>
      <MemoryRouter initialEntries={['/']}>
        <App />
      </MemoryRouter>
    </StrictMode>,
  );
  console.log('Acuity: React Rendered!'); // Check if this appears in Inspect Popup
} else {
  console.error('Acuity: Root element not found!');
}
