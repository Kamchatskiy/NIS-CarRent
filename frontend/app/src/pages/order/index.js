import * as React from "react";
import { DemoContainer } from "@mui/x-date-pickers/internals/demo";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import {
  Box,
  Button,
  CircularProgress,
  Stack,
  TextField,
  Typography,
  Paper,
} from "@mui/material";
import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";

export const Order = () => {
  const [email, setEmail] = React.useState("");
  const [carId, setCarId] = React.useState("");
  const [startDate, setStartDate] = React.useState(null);
  const [endDate, setEndDate] = React.useState(null);
  const [loading, setLoading] = React.useState(false);
  const [price, setPrice] = React.useState(null);

  const handleOrder = async () => {
    if (!email || !carId || !startDate || !endDate) {
      alert("Please fill in all fields.");
      return;
    }

    const orderData = {
      client_email: email,
      car_id: parseInt(carId, 10),
      start_date: startDate.toISOString(),
      end_date: endDate.toISOString(),
    };

    setLoading(true);
    setPrice(null);

    try {
      const response = await fetch("http://127.0.0.1:8080/order", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(orderData),
      });

      const responseData = await response.json();
      console.log("Order successful:", responseData);
      setPrice(responseData.price);
    } catch (error) {
      console.error("Error sending order:", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <Typography
        variant="h3"
        color="text.primary"
        textAlign="center"
        margin={3}
      >
        Notice that you should be registered first
      </Typography>

      <Box
        margin={4}
        display="flex"
        justifyContent="space-between"
        alignItems="flex-start"
      >
        <Stack spacing={4} width="40%">
          <TextField
            required
            variant="standard"
            label="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <TextField
            required
            variant="standard"
            label="Car ID"
            value={carId}
            onChange={(e) => setCarId(e.target.value)}
          />

          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <DemoContainer components={["DatePicker"]}>
              <DatePicker
                label="Choose Start Date"
                value={startDate}
                onChange={(newValue) => setStartDate(newValue)}
              />
              <DatePicker
                label="Choose End Date"
                value={endDate}
                onChange={(newValue) => setEndDate(newValue)}
              />
            </DemoContainer>
          </LocalizationProvider>
          <Button
            variant="contained"
            size="large"
            sx={{ width: "30%", height: "50%" }}
            endIcon={<ShoppingCartIcon />}
            onClick={handleOrder}
            disabled={loading}
          >
            {loading ? <CircularProgress size={24} /> : "Order"}
          </Button>
        </Stack>

        <Paper
          sx={{
            padding: 2,
            width: "50%",
            height: "300px",
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            alignItems: "center",
            boxShadow: 3,
          }}
        >
          <Typography variant="h6">Your total price is:</Typography>
          {loading ? (
            <CircularProgress />
          ) : (
            price !== null && <Typography variant="h5">{price}$</Typography>
          )}
        </Paper>
      </Box>
    </>
  );
};
