import * as React from "react";
import Box from "@mui/material/Box";
import { Button, TextField } from "@mui/material";
import ListAltIcon from "@mui/icons-material/ListAlt";
import { DataGrid } from "@mui/x-data-grid";

const columns = [
  {
    field: "car_id",
    type: "number",
    headerName: "Car ID",
    flex: 1,
    align: "center",
    headerAlign: "center",
  },
  {
    field: "price",
    headerName: "Price",
    type: "number",
    flex: 1,
    align: "center",
    headerAlign: "center",
  },
  {
    field: "start_date",
    headerName: "Start Date",
    type: "date",
    flex: 1,
    align: "center",
    headerAlign: "center",
  },
  {
    field: "end_date",
    headerName: "End Date",
    type: "date",
    flex: 1,
    align: "center",
    headerAlign: "center",
  },
];

const paginationModel = { page: 0, pageSize: 10 };

export const MyOrders = () => {
  const [email, setEmail] = React.useState("");
  const [rows, setRows] = React.useState([]);

  const handleEmailChange = (event) => {
    setEmail(event.target.value);
  };

  const handleButtonClick = async () => {
    const response = await fetch(
      `http://127.0.0.1:8080/orders?email=${encodeURIComponent(email)}`
    );
    const data = await response.json();
    const formattedRents = data.rents.map((rent) => ({
      ...rent,
      start_date: new Date(rent.start_date),
      end_date: new Date(rent.end_date),
    }));
    setRows(formattedRents);
  };

  return (
    <>
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          justifyContent: "center",
          height: "10%",
          paddingTop: 2,
        }}
      >
        <TextField
          required
          variant="standard"
          label="Email"
          value={email}
          onChange={handleEmailChange}
        />
        <Button
          variant="contained"
          size="large"
          sx={{ width: "30%", height: "50%", marginLeft: 2 }}
          endIcon={<ListAltIcon />}
          onClick={handleButtonClick}
        >
          My Orders
        </Button>
      </Box>

      <DataGrid
        columns={columns}
        rows={rows}
        initialState={{ pagination: { paginationModel } }}
        pageSizeOptions={[10, 50, 100]}
        sx={{ border: 0 }}
      />
    </>
  );
};
