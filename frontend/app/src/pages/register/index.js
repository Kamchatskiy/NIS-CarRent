import * as React from "react";
import Box from "@mui/material/Box";
import SendIcon from "@mui/icons-material/Send";
import { Button, Stack, CircularProgress, Typography } from "@mui/material";
import { CustomTextField } from "./components/textfield";

export const Register = () => {
  const [name, setName] = React.useState("");
  const [surname, setSurname] = React.useState("");
  const [email, setEmail] = React.useState("");
  const [phone_number, setPhoneNumber] = React.useState("");
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState(null);

  const handleRegister = async () => {
    const data = {
      name,
      surname,
      email,
      phone_number,
    };

    setLoading(true);
    setError(null);

    try {
      const response = await fetch("YOUR_BACKEND_URL_HERE", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      // eslint-disable-next-line
      const result = await response.json();
    } catch (error) {
      setError("An error occurred, please try again");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box width="40%" margin={4}>
      {loading && (
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
      )}

      {error && (
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            height: 400,
          }}
        >
          <Typography variant="h3" color="error">
            {error}
          </Typography>
        </Box>
      )}

      {!loading && !error && (
        <Stack spacing={3}>
          <CustomTextField
            label="Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <CustomTextField
            label="Last Name"
            value={surname}
            onChange={(e) => setSurname(e.target.value)}
          />
          <CustomTextField
            label="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <CustomTextField
            label="Phone Number"
            value={phone_number}
            onChange={(e) => setPhoneNumber(e.target.value)}
          />

          <Button
            variant="contained"
            size="large"
            sx={{ width: "30%" }}
            endIcon={<SendIcon />}
            onClick={handleRegister}
          >
            Register
          </Button>
        </Stack>
      )}
    </Box>
  );
};
