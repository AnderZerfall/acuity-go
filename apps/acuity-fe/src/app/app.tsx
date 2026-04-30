import { Route, Routes } from 'react-router-dom';
import { MainScreen } from './screens/main/MainScreen';

export function App() {
  return (
    <Routes>
      <Route path="/" element={<MainScreen />} index />
    </Routes>
  );
}

export default App;
