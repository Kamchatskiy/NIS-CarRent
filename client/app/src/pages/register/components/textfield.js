import * as React from "react";
import { LinearProgress, Stack, TextField } from "@mui/material";

export const CustomTextField = ({ label, value, onChange }) => {
  const showProgress = value === "";

  return (
    <Stack>
      <TextField
        required
        variant="standard"
        label={label}
        value={value}
        onChange={onChange}
      />
      {showProgress && <LinearProgress color="secondary" />}
    </Stack>
  );
};
