import React from 'react'
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import './TileContainer.scss'


  export default class TileContainer extends React.Component {  
    constructor(props){
      super(props);
      this.state = {
        borrowerName: "Anupa Suresh",
        loanAmt: 1200,
        loanEMI: 100,
      };
    }
    render() {
      return (
        <Card className="tileContent-root">
        <CardContent>
          <Typography className="tileContent-title" color="textSecondary" gutterBottom>
            Lending Info
          </Typography>
          <Typography variant="h5" component="h2">
           {this.state.borrowerName}
          </Typography>
          <Typography className="tileContent-pos" color="textSecondary">
            Loan Amount of {this.state.loanAmt}
          </Typography>
          <Typography variant="body2" component="p">
            Loan EMI per month of {this.state.loanEMI}
            <br />
          </Typography>
        </CardContent>
        <CardActions>
          <Button size="small">More</Button>
        </CardActions>
      </Card>
      );
    };
  }