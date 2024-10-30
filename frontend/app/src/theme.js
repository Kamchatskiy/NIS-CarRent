import { createTheme } from "@mui/material/styles";

export const mainThemeOptions = {
  palette: {
    mode: "dark",
    primary: {
      main: "#00fff6",
    },
    secondary: {
      main: "#00adff",
    },
    background: {
      default: "#000000",
      paper: "#000000",
    },
    text: {
      primary: "#00fff7",
      secondary: "#00adff",
      hint: "#085aa5",
      disabled: "#d8fffc",
    },
    warning: {
      main: "#FCA311",
    },
    error: {
      main: "#D00000",
    },
    info: {
      main: "#065A82",
    },
    success: {
      main: "#72FF0E",
    },
    divider: "#5F7470",
  },
};

export const MainTheme = createTheme(mainThemeOptions);
