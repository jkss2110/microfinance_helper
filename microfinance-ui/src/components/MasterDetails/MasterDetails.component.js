import React from "react";
import "./MasterDetails.scss";
import {
  Container,
  Divider,
  Typography,
  TextField,
  Button,
} from "@material-ui/core";
import { Icon } from "@material-ui/core";
import { green } from "@material-ui/core/colors";
export default class MasterDetails extends React.Component {
  render() {
    return (
      <>
        <Container className="masterDetail-root">
          <Container className="masterDetail-header">
            <Typography class="masterDetail-title" variant="h5" component="h2">
              Anupa Suresh
            </Typography>
            <Typography class="masterDetail-status" variant="h6" component="h5">
              <Icon style={{ color: green[500] }}>add_circle</Icon>
              Active
            </Typography>
          </Container>
          <Divider></Divider>

          <div className="masterDetail-loanbalance">
            <Typography variant="body" component="textSecondary">
              Loan Amount taken is 1200
            </Typography>
            <Typography variant="body1" component="textSecondary">
              Loan Amount taken is 1200
            </Typography>
          </div>
          <div className="masterDetail-info">
            <span>
              <p> INR 100 EMI/month</p>
              <p> Next due payment after 10 days</p>
            </span>
            <span className="masterDetail-info2">
              <p>Next Payment date 11/04/2022</p>
              <p> Score reduction if not paid on time is 10</p>
            </span>
          </div>
          <Divider></Divider>
          <div className="masterDetail-newdate">
            <Typography variant="body1" component="textSecondary">
              Do you like to propose a new payment date :
            </Typography>
            <TextField
              id="date"
              type="date"
              className="masterDetail-datepicker"
              defaultValue="2022-04-11"
            ></TextField>
            <Button
              color="primary"
              endIcon={<Icon>send</Icon>}
              className="masterDetail-sendBtn"
            >
              Send
            </Button>
          </div>
          <TextField
            className="masterDetail-Feedback"
            id="standard-secondary"
            label="Feeds to the Borrower"
            color="secondary"
          />
          <span className="masterDetail-feedbackBtn">
          <Button variant="contained">Clear</Button>
          <Button variant="contained" color="primary">
            Send Feed
          </Button>
          </span>
        </Container>
      </>
    );
  }
}
