import React from "react";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import "./TileContainer.scss";

class TileContainer extends React.Component {
  render() {
    const loanInfo = this.props.loans;
    const routingFunction = this.props.routingFunction;
    return (
      <Card
        key={loanInfo.BorrowerName}
        className="tileContent-root"
        onClick={() => routingFunction(`/detail`)}
      >
        <CardContent>
          <Typography variant="h5" component="h2">
            {loanInfo.BorrowerName}
          </Typography>
          <Typography className="tileContent-pos" color="textSecondary">
            Loan Amount {loanInfo.LoanAmt}
          </Typography>
          <Typography variant="body2" component="p">
            Loan EMI per month {loanInfo.LoanEMI}
            <br />
          </Typography>
          <Typography variant="body2" component="p">
            Balance {loanInfo.LoanEMI}
            <br />
          </Typography>
        </CardContent>
      </Card>
    );
  }
}
export default TileContainer;
