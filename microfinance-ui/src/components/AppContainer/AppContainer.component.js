import React from "react";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import './AppContainer.scss';
//import props from 'prop-types';

const AppContainer = (props) => {
    return (
      <div className="App-root">
        <AppBar className="App-bar" position="static">
          <Toolbar>
            <IconButton
              edge="start"
              className="App-menuButton"
              color="inherit"
              aria-label="menu"
            >
              <MenuIcon />
            </IconButton>
            <Typography variant="h6" className="App-title">
              {props.title}
            </Typography>
            <Button color="inherit">Login</Button>
          </Toolbar>
        </AppBar>
      </div>
    );
  }
export default AppContainer;
