import { Route, Routes } from "react-router-dom";
import Layout from "./components/layout/layout";
import Dashboard from "./pages/dashboard/dashboard";
import Product from "./pages/product/product";
import Admin from "./pages/admin/admin";
import Login from "./pages/login/login";
import NotFound from "./pages/notfound/notfound";

export default function App() {
    return (
        <Routes>
            <Route element={<Layout />}>
                <Route index element={<Dashboard />} />
                <Route path="products" element={<Product />} />
                <Route path="admins" element={<Admin />} />
            </Route>
            <Route>
                <Route path="login" element={<Login />} />
                <Route path="*" element={<NotFound />} />
            </Route>
        </Routes>
    )
}