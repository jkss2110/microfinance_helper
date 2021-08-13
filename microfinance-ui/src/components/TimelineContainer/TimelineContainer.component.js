import React from "react";
import {
  TimelineDot,
  TimelineOppositeContent,
  TimelineContent,
  TimelineConnector,
  TimelineSeparator,
  TimelineItem,
  Timeline,
} from "@material-ui/lab";

import {
  HourglassEmpty as HourglassEmptyIcon,
  AttachMoney as AttachMoneyIcon,
  MonetizationOn as MonetizationOnIcon,
} from "@material-ui/icons";
import { Typography, Paper } from "@material-ui/core";
import "./TimelineContainer.scss";

export default function TimelineContainer() {
  return (
    <Timeline align="alternate">
      <TimelineItem>
        <TimelineOppositeContent>
          <Typography variant="body2" color="textSecondary">
            10/01/2020
          </Typography>
        </TimelineOppositeContent>
        <TimelineSeparator>
          <TimelineDot>
            <HourglassEmptyIcon />
          </TimelineDot>
          <TimelineConnector />
        </TimelineSeparator>
        <TimelineContent>
          <Paper elevation={3} className="timeline-paper">
            <Typography variant="h6" component="h1">
              Loan taken
            </Typography>
            <Typography>Amount of INR 1200</Typography>
          </Paper>
        </TimelineContent>
      </TimelineItem>
      <TimelineItem>
        <TimelineOppositeContent>
          <Typography variant="body2" color="textSecondary">
            25/02/2020
          </Typography>
        </TimelineOppositeContent>
        <TimelineSeparator>
          <TimelineDot color="primary">
            <AttachMoneyIcon />
          </TimelineDot>
          <TimelineConnector />
        </TimelineSeparator>
        <TimelineContent>
          <Paper elevation={3} className="timeline-paper">
            <Typography variant="h6" component="h1">
              EMI Payment
            </Typography>
            <Typography>First Installment of INR200 paid</Typography>
          </Paper>
        </TimelineContent>
      </TimelineItem>
      <TimelineItem>
        <TimelineOppositeContent>
          <Typography variant="body2" color="textSecondary">
            25/03/2020
          </Typography>
        </TimelineOppositeContent>
        <TimelineSeparator>
          <TimelineDot color="primary" variant="outlined">
            <MonetizationOnIcon />
          </TimelineDot>
          <TimelineConnector className="timeline-secondaryTail" />
        </TimelineSeparator>
        <TimelineContent>
          <Paper elevation={3} className="timeline-paper">
            <Typography variant="h6" component="h1">
              EMI Payment
            </Typography>
            <Typography>Second Installment of INR200 paid</Typography>
          </Paper>
        </TimelineContent>
      </TimelineItem>
      <TimelineItem>
        <TimelineSeparator>
          <TimelineDot color="secondary">
            <AttachMoneyIcon />
          </TimelineDot>
        </TimelineSeparator>
        <TimelineContent>
          <Paper elevation={3} className="timeline-paper">
            <Typography variant="h6" component="h1">
              EMI Payment
            </Typography>
            <Typography>Third Installment of INR200 paid</Typography>
          </Paper>
        </TimelineContent>
      </TimelineItem>
    </Timeline>
  );
}
