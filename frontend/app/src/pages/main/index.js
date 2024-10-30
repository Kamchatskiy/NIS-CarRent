import * as React from "react";
import Box from "@mui/material/Box";
import DirectionsCarIcon from "@mui/icons-material/DirectionsCar";
import { Divider, List, Typography } from "@mui/material";
import { Person } from "./components/person";

export const Main = () => {
  return (
    <Box
      display="flex"
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      height="70vh"
    >
      <Typography
        variant="h1"
        color="text.primary"
        display="flex"
        alignItems="center"
        sx={{ marginBottom: 3 }}
      >
        Car Rent App
        <DirectionsCarIcon
          sx={{ color: "text.primary", fontSize: "inherit", marginLeft: 1 }}
        />
      </Typography>

      <Box display="flex" alignItems="center">
        <Typography variant="h3" color="text.secondary" sx={{ marginRight: 1 }}>
          Made by
        </Typography>
        <List sx={{ display: "flex", flexDirection: "row", padding: 0 }}>
          <Divider orientation="vertical" flexItem />
          <Person name="Marat" source="/static/images/avatars/marat.jpg" />
          <Divider orientation="vertical" flexItem />
          <Person name="Vova" source="/static/images/avatars/vova.jpg" />
          <Divider orientation="vertical" flexItem />
          <Person name="Egor" source="/static/images/avatars/egor.jpg" />
          <Divider orientation="vertical" flexItem />
          <Person name="Tima" source="/static/images/avatars/tima.jpg" />
          <Divider orientation="vertical" flexItem />
        </List>
      </Box>
    </Box>
  );
};
