import React from "react";
import { Routes, Route, BrowserRouter } from "react-router-dom";
import { Header } from "./components/header";
import { Cars } from "./pages/cars";
import { Main } from "./pages/main";
import { Order } from "./pages/order";
import { MyOrders } from "./pages/my-orders";
import { Footer } from "./footer";
import { ThemeProvider } from "@emotion/react";
import { MainTheme } from "./theme";

export const App = () => {
  return (
    <>
      <ThemeProvider theme={MainTheme}>
        <BrowserRouter>
          <Header />
          <Routes>
            <Route path="/" element={<Main />} />
            <Route path="/cars" element={<Cars />} />
            <Route path="/order" element={<Order />} />
            <Route path="/my-orders" element={<MyOrders />} />
          </Routes>
        </BrowserRouter>
        <Footer />
      </ThemeProvider>
    </>
  );
};
