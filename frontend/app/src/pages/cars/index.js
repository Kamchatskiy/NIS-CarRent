import * as React from "react";
import { DataGrid } from "@mui/x-data-grid";
import Paper from "@mui/material/Paper";
import { Box, CircularProgress, Typography } from "@mui/material";

const columns = [
  {
    field: "brand",
    headerName: "Brand",
    flex: 1,
    align: "center",
    headerAlign: "center",
  },
  {
    field: "model",
    headerName: "Model",
    flex: 1,
    align: "center",
    headerAlign: "center",
  },
  {
    field: "year",
    headerName: "Year",
    type: "number",
    flex: 0.5,
    align: "center",
    headerAlign: "center",
  },
  {
    field: "daily_price",
    headerName: "Daily Rent Price",
    type: "number",
    flex: 1,
    align: "center",
    headerAlign: "center",
  },
  {
    field: "insurance_price",
    headerName: "Insurance Price",
    type: "number",
    flex: 1,
    align: "center",
    headerAlign: "center",
  },
];

const paginationModel = { page: 0, pageSize: 10 };

export const Cars = () => {
  const [rows, setRows] = React.useState([]);
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState(null);

  React.useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("http://127.0.0.1:8080/cars");
        const data = await response.json();

        const transformedData = data.map((car) => ({
          id: car.id,
          brand: car.brand,
          model: car.model,
          year: car.year,
          daily_price: car.daily_price,
          insurance_price: car.insurance_price,
        }));

        setRows(transformedData);
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) {
    return (
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          height: 400,
        }}
      >
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          height: 400,
        }}
      >
        <Typography variant="h2" color="error">
          An error occurred
        </Typography>
      </Box>
    );
  }

  return (
    <Box>
      <Paper variant="outlined" sx={{ height: "100%", width: "100%" }}>
        <DataGrid
          rows={rows}
          columns={columns}
          pageSizeOptions={[10, 50, 100]}
          checkboxSelection
          sx={{ border: 0, width: "100%" }}
          initialState={{ pagination: { paginationModel } }}
        />
      </Paper>
    </Box>
  );
};
