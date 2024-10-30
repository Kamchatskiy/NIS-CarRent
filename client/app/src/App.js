import React from "react";
import { Routes, Route, BrowserRouter } from "react-router-dom";
import { Header } from "./components/header";
import { Cars } from "./pages/cars";
import { Main } from "./pages/main";
import { Order } from "./pages/order";
import { MyOrders } from "./pages/my-orders";
import { Footer } from "./components/footer";
import { ThemeProvider } from "@emotion/react";
import { MainTheme } from "./theme";
import "./style.css";
import { Register } from "./pages/register";

export const App = () => {
  return (
    <>
      <ThemeProvider theme={MainTheme}>
        <BrowserRouter>
          <Header />
          <Routes>
            <Route path="/" element={<Main />} />
            <Route path="/cars" element={<Cars />} />
            <Route path="/register" element={<Register />} />
            <Route path="/order" element={<Order />} />
            <Route path="/my-orders" element={<MyOrders />} />
          </Routes>
        </BrowserRouter>
        <Footer />
      </ThemeProvider>
    </>
  );
};
