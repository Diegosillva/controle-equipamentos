import { createBrowserRouter } from "react-router-dom"
import Cadastro from "./pages/cadastro.jsx";
import Consulta from "./pages/consulta.jsx";

const router = createBrowserRouter([
    { path: '/', element: <Consulta /> },
    { path: '/cadastro', element: <Cadastro /> }
]);
export default router;
