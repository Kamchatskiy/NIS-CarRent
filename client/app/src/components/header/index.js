import * as React from "react";
import { Link, useLocation } from "react-router-dom";
import Tabs from "@mui/material/Tabs";
import Tab from "@mui/material/Tab";
import Box from "@mui/material/Box";
import { Paper } from "@mui/material";

export const Header = () => {
  // eslint-disable-next-line
  const [listVisibility, setListVisibility] = React.useState(false);
  const offListVisibility = () => {
    setListVisibility(false);
  };

  const location = useLocation();
  const [activeTab, setActiveTab] = React.useState("/");

  React.useEffect(() => {
    setActiveTab(location.pathname);
  }, [location.pathname]);

  const handleChangeTab = (event, newTab) => {
    setActiveTab(newTab);
    offListVisibility();
  };

  return (
    <Paper variant="outlined" square={false} sx={{border: 3}}>
      <Box
        sx={{
          width: "100%",
        }}
      >
        <Tabs
          onChange={handleChangeTab}
          value={activeTab}
          textColor="primary"
          indicatorColor="primary"
          sx={{
            width: "100%",
          }}
        >
          <Tab
            value="/"
            label="Main Page"
            component={Link}
            to="/"
            sx={{ flex: 1 }}
          />
          <Tab
            value="/cars"
            label="Cars"
            component={Link}
            to="/cars"
            sx={{ flex: 1 }}
          />
          <Tab
            value="/order"
            label="Order"
            component={Link}
            to="/order"
            sx={{ flex: 1 }}
          />
          <Tab
            value="/my-orders"
            label="My Orders"
            component={Link}
            to="/my-orders"
            sx={{ flex: 1 }}
          />
        </Tabs>
      </Box>
    </Paper>
  );
};
