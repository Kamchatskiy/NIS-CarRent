import * as React from "react";
import Box from "@mui/material/Box";
import SendIcon from "@mui/icons-material/Send";
import {
  Button,
  Stack,
  CircularProgress,
  Typography,
  Alert,
} from "@mui/material";
import { CustomTextField } from "./components/textfield";

export const Register = () => {
  const [name, setName] = React.useState("");
  const [surname, setSurname] = React.useState("");
  const [email, setEmail] = React.useState("");
  const [phone_number, setPhoneNumber] = React.useState("");
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState(null);
  const [success, setSuccess] = React.useState(null); // New state for success message

  const handleRegister = async () => {
    const data = {
      name,
      surname,
      email,
      phone_number,
    };

    setLoading(true);
    setError(null);
    setSuccess(null); // Reset success message

    try {
      const response = await fetch("http://127.0.0.1:8080/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        // If the response is successful
        setSuccess("Registration successful!"); // Set success message
      } else {
        const result = await response.json();
        setError(result.message || "An error occurred"); // Set error message
      }
    } catch (err) {
      setError("An error occurred, please try again"); // Handle network errors
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
            An error occurred, please try again
          </Typography>
        </Box>
      )}

      {success && (
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            height: 400,
          }}
        >
          <Alert severity="success">{success}</Alert>{" "}
          {/* Display success alert */}
        </Box>
      )}

      {!loading && !error && !success && (
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
