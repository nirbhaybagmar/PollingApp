import React, { Component } from "react";
import "./Poll.css";
import { Avatar, Icon } from "antd";
import { Link } from "react-router-dom";
import { getAvatarColor } from "../util/Colors";
import { formatDateTime } from "../util/Helpers";

import { Radio, Button } from "antd";
const RadioGroup = Radio.Group;

class Poll extends Component {
  calculatePercentage = (choice) => {
    if (this.props.totalVotes[this.props.poll.id]) {
      if (this.props.totalVotes[this.props.poll.id] === 0) {
        return 0;
      }
      if (this.props.totalVotesEachOption[choice]) {
        return (
          (this.props.totalVotesEachOption[choice] * 100) /
          this.props.totalVotes[this.props.poll.id]
        );
      } else {
        return 0;
      }
    } else {
      return 0;
    }
  };

  getTotalVotes = () => {
    var result = 0;
    this.props.poll.edges.answer.forEach((option) => {
      if (option.numOfVotes) {
        result += option.numOfVotes;
      }
    });
    return result;
  };

  isPollVoted = () => {
    var result = false;
    this.props.votes.forEach((vote) => {
      if (
        vote.user_id === this.props.user.id &&
        this.props.poll.id === vote.question_id
      ) {
        result = true;
      }
    });
    return result;
  };

  isSelected = (choice) => {
    var result = false;
    this.props.votes.forEach((vote) => {
      if (
        vote.user_id === this.props.user.id &&
        this.props.poll.id === vote.question_id &&
        choice.id === vote.answer_id
      ) {
        result = true;
      }
    });
    return result;
  };

  getWinningChoice = () => {
    var maxVote = 0;
    var result;
    this.props.poll.edges.answer.forEach((option) => {
      if (option.numOfVotes) {
        if (option.numOfVotes > maxVote) {
          result = option.id;
          maxVote = option.numOfVotes;
        }
      }
    });
    return result;
  };

  getTimeRemaining = (poll) => {
    const expirationTime = new Date(poll.pollExpiry).getTime();
    const currentTime = new Date().getTime();

    var difference_ms = expirationTime - currentTime;
    var seconds = Math.floor((difference_ms / 1000) % 60);
    var minutes = Math.floor((difference_ms / 1000 / 60) % 60);
    var hours = Math.floor((difference_ms / (1000 * 60 * 60)) % 24);
    var days = Math.floor(difference_ms / (1000 * 60 * 60 * 24));

    let timeRemaining;

    if (days > 0) {
      timeRemaining = days + " days left";
    } else if (hours > 0) {
      timeRemaining = hours + " hours left";
    } else if (minutes > 0) {
      timeRemaining = minutes + " minutes left";
    } else if (seconds > 0) {
      timeRemaining = seconds + " seconds left";
    } else {
      timeRemaining = "less than a second left";
    }

    return timeRemaining;
  };

  render() {
    const pollChoices = [];
    var isPollExpired =
      new Date(this.props.poll.pollExpiry).getTime() < new Date().getTime();
    if (this.isPollVoted() || isPollExpired) {
      const winningChoice = isPollExpired ? this.getWinningChoice() : null;

      this.props.poll.edges.answer.forEach((choice) => {
        pollChoices.push(
          <CompletedOrVotedPollChoice
            key={choice.id}
            choice={choice.body}
            isWinner={winningChoice && choice.id === winningChoice}
            isSelected={this.isSelected(choice)}
            percentVote={this.calculatePercentage(choice.id)}
          />
        );
      });
    } else {
      this.props.poll.edges.answer.forEach((choice) => {
        pollChoices.push(
          <Radio
            className="poll-choice-radio"
            key={choice.id}
            value={choice.id}
          >
            {choice.body}
          </Radio>
        );
      });
    }
    return (
      <div className="poll-content">
        <div className="poll-header">
          <div className="poll-creator-info">
            <Link
              className="creator-link"
              to={`/users/${this.props.poll.edges.user.name}`}
            >
              <Avatar
                className="poll-creator-avatar"
                style={{
                  backgroundColor: getAvatarColor(
                    this.props.poll.edges.user.name[0]
                  ),
                }}
              >
                {this.props.poll.edges.user.name[0].toUpperCase()}
              </Avatar>
              <span className="poll-creator-name">
                {this.props.poll.edges.user.name}
              </span>
              <span className="poll-creator-username">
                @{this.props.poll.edges.user.name}
              </span>
              <span className="poll-creation-date">
                {formatDateTime(this.props.poll.created_at)}
              </span>
            </Link>
          </div>
          <div className="poll-question">{this.props.poll.body}</div>
        </div>
        <div className="poll-choices">
          <RadioGroup
            className="poll-choice-radio-group"
            onChange={this.props.handleVoteChange}
            value={this.props.currentVote}
          >
            {pollChoices}
          </RadioGroup>
        </div>
        <div className="poll-footer">
          {!(this.isPollVoted() || isPollExpired) ? (
            <Button
              className="vote-button"
              disabled={!this.props.currentVote}
              onClick={this.props.handleVoteSubmit}
            >
              Vote
            </Button>
          ) : null}
          <span className="total-votes">{this.getTotalVotes()} votes</span>
          <span className="separator">•</span>
          <span className="time-left">
            {isPollExpired
              ? "Final results"
              : this.getTimeRemaining(this.props.poll)}
          </span>
        </div>
      </div>
    );
  }
}

function CompletedOrVotedPollChoice(props) {
  return (
    <div className="cv-poll-choice">
      <span className="cv-poll-choice-details">
        <span className="cv-choice-percentage">
          {Math.round(props.percentVote * 100) / 100}%
        </span>
        <span className="cv-choice-text">{props.choice}</span>
        {props.isSelected ? (
          <Icon className="selected-choice-icon" type="check-circle-o" />
        ) : null}
      </span>
      <span
        className={
          props.isWinner
            ? "cv-choice-percent-chart winner"
            : "cv-choice-percent-chart"
        }
        style={{ width: props.percentVote + "%" }}
      ></span>
    </div>
  );
}

export default Poll;
