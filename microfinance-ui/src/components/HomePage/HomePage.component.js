import React from "react";
import TileContainer from "../TileContainer/TileContainer.component";
import { Chip, Divider } from "@material-ui/core";
import {scroller} from "react-scroll"
import "./HomePage.scss";

class HomePage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loanInfo: [
        {
          borrowerID: 2,
          lenderID: 1,
          borrowerName: "Anupa Suresh",
          loanAmt: 1200,
          loanEMI: 100,
        },
        {
          borrowerID: 2,
          lenderID: 3,
          borrowerName: "Jayakrishnan S",
          loanAmt: 1200,
          loanEMI: 100,
        },
        {
          borrowerID: 2,
          lenderID: 4,
          borrowerName: "Bottle",
          loanAmt: 1200,
          loanEMI: 100,
        },
        {
          borrowerID: 2,
          lenderID: 5,
          borrowerName: "Bond James",
          loanAmt: 1200,
          loanEMI: 100,
        },
      ],
    };
  }
  renderTileContent() {
    const loanInfos = this.state.loanInfo.slice();
    const tileContent = loanInfos.map((loanInfo) => {
      return (
        <TileContainer
          key={loanInfo.borrowerName}
          loans={loanInfo}
        ></TileContainer>
      );
    });
    return tileContent;
  }
  render() {
    const tiles = this.renderTileContent();
    const option ={
      duration: 1500,
      delay: 100,
      smooth: 'easeInOutQuint',
    };
    return (
      <>
        <div className="chipContainer">
          <Chip className="lenderChip" color="primary" label="Lending Info" onClick={()=> {scroller.scrollTo("lendingInfo",option);}}/>
          <Chip className="borrowerChip" label="Borrowing Info" onClick={()=> {scroller.scrollTo("borrowingInfo",option);}}/>
          <Chip className="erupiChip" label="Coupons" onClick={()=> {scroller.scrollTo("eRupiCoupon",option);}}/>
          <Chip className="analyticsChip" label="Scores and Projection" onClick={()=> {scroller.scrollTo("scoreProjection",option);}}/>
        </div>
        <h1 name="lendingInfo" className="tileContianer-header">Lending Info</h1>
        <Divider />
        <div className="tileContainer-content">{tiles}</div>
        <h1 name="borrowingInfo" className="tileContianer-header">Borrowing Info</h1>
        <Divider />
        <div className="tileContainer-content">{tiles}</div>
        <h1 name="eRupiCoupon" className="tileContianer-header">e-Rupi Coupons</h1>
        <Divider />
        <div className="tileContainer-content">{tiles}</div>
        <h1 name="scoreProjection" className="tileContianer-header">Scores and Projection</h1>
        <Divider />
      </>
    );
  }
}
export default HomePage;
