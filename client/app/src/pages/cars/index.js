import * as React from "react";
import { DataGrid } from "@mui/x-data-grid";
import Paper from "@mui/material/Paper";
import { Box, CircularProgress, Typography } from "@mui/material";

const columns = [
  { field: "Brand", headerName: "Brand", width: 130 },
  { field: "Model", headerName: "Model", width: 130 },
  {
    field: "Year",
    headerName: "Year",
    type: "number",
    width: 80,
  },
  {
    field: "Daily Rent Price",
    headerName: "Daily Price",
    type: "number",
    width: 100,
  },
  {
    field: "Insurance Price",
    headerName: "Insurance Price",
    type: "number",
    width: 100,
  },
];

const paginationModel = { page: 0, pageSize: 5 };

export const Cars = () => {
  const [rows, setRows] = React.useState([]);
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState(null);

  React.useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("/cars");
        const data = await response.json();
        setRows(data);
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
          An error occured
        </Typography>
      </Box>
    );
  }

  return (
    <Box>
      <Paper sx={{ height: 400, width: "100%" }}>
        <DataGrid
          rows={rows}
          columns={columns}
          initialState={{ pagination: { paginationModel } }}
          pageSizeOptions={[5, 10]}
          checkboxSelection
          sx={{ border: 0 }}
        />
      </Paper>
    </Box>
  );
};
