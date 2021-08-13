import MasterDetails from "../MasterDetails/MasterDetails.component";
import TimelineContainer from "../TimelineContainer/TimelineContainer.component";
import { Grid,Paper } from "@material-ui/core";
import './DetailPageContainer.scss';
export default function DetailPageContainer(props) {
  return (
      <Grid container className="detailPage-root" justifyContent="left" spacing={5}>
          <Grid key="1" item className="masterDetail-grid">
            <Paper className="masterDetail-paper">
              <MasterDetails></MasterDetails>
            </Paper>
          </Grid>
          <Grid key="2" item>
              <TimelineContainer></TimelineContainer>
          </Grid>
      </Grid>
  );
}
