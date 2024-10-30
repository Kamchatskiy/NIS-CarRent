import React from "react";
import { ListItem, ListItemAvatar, Avatar, ListItemText } from "@mui/material";

export const Person = ({ name, source }) => {
  return (
    <>
      <ListItem
        alignItems="center"
        sx={{ display: "flex", justifyContent: "center" }}
      >
        <ListItemAvatar>
          <Avatar alt={name} src={source} />
        </ListItemAvatar>
        <ListItemText
          primary={name}
          sx={{ textAlign: "center", flexGrow: 1, color: "text.primary" }}
        />
      </ListItem>
    </>
  );
};
