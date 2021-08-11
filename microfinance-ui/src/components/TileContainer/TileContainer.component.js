import React from "react";
import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import "./TileContainer.scss";

export default class TileContainer extends React.Component {
  render() {
    const loanInfo = this.props.loans;
    return (
      <Card className="tileContent-root">
        <CardContent>
          <Typography variant="h5" component="h2">
            {loanInfo.borrowerName}
          </Typography>
          <Typography className="tileContent-pos" color="textSecondary">
            Loan Amount of {loanInfo.loanAmt}
          </Typography>
          <Typography variant="body2" component="p">
            Loan EMI per month of {loanInfo.loanEMI}
            <br />
          </Typography>
        </CardContent>
        <CardActions>
          <Button size="small">More</Button>
        </CardActions>
      </Card>
    );
  }
}
