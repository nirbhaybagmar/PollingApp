import React, { Component } from "react";
import {
  getAllPolls,
  getUserCreatedPolls,
  getUserVotedPolls,
  getVotes,
} from "../util/APIUtils";
import Poll from "./Poll";
import { castVote } from "../util/APIUtils";
import LoadingIndicator from "../common/LoadingIndicator";
import { Button, Icon, notification } from "antd";
import { POLL_LIST_SIZE } from "../constants";
import { withRouter } from "react-router-dom";
import "./PollList.css";

class PollList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      polls: [],
      votes: [],
      totalVotes: {},
      totalVotesEachOption: {},
      last: true,
      currentVotes: [],
      isLoading: false,
    };
    this.loadPollList = this.loadPollList.bind(this);
    this.handleLoadMore = this.handleLoadMore.bind(this);
  }

  loadVotesList() {
    getVotes()
      .then((response) => {
        this.setState(
          {
            votes: response,
            isLoading: false,
          },
          () => {
            this.sortVotesByQuestion();
          }
        );
      })
      .catch((err) => {
        this.setState({
          isLoading: false,
        });
      });
  }

  sortVotesByQuestion() {
    var totalVotes = {};
    var totalVotesEachOption = {};
    this.state.votes.forEach((vote) => {
      if (totalVotes[vote.question_id]) {
        totalVotes[vote.question_id] += 1;
      } else {
        totalVotes[vote.question_id] = 1;
      }
      if (totalVotesEachOption[vote.answer_id]) {
        totalVotesEachOption[vote.answer_id] += 1;
      } else {
        totalVotesEachOption[vote.answer_id] = 1;
      }
    });

    this.setState({
      totalVotes,
      totalVotesEachOption,
    });
  }

  loadPollList(page = 0, size = POLL_LIST_SIZE) {
    let promise;
    if (this.props.username) {
      if (this.props.type === "USER_CREATED_POLLS") {
        promise = getUserCreatedPolls(this.props.currentUser.id, page, size);
      } else if (this.props.type === "USER_VOTED_POLLS") {
        promise = getUserVotedPolls(this.props.userId, page, size);
      }
    } else {
      promise = getAllPolls();
    }

    if (!promise) {
      return;
    }

    this.setState({
      isLoading: true,
    });

    promise
      .then((response) => {
        this.setState({
          polls: response,
          isLoading: false,
        });
        this.loadVotesList();
      })
      .catch((error) => {
        this.setState({
          isLoading: false,
        });
      });
  }

  componentDidMount() {
    this.loadPollList();
  }

  componentDidUpdate(nextProps) {
    if (this.props.isAuthenticated !== nextProps.isAuthenticated) {
      this.loadPollList();
    }
  }

  handleLoadMore() {
    this.loadPollList(this.state.page + 1);
  }

  handleVoteChange(event, pollIndex) {
    const currentVotes = this.state.currentVotes.slice();
    currentVotes[pollIndex] = event.target.value;
    this.setState({
      currentVotes: currentVotes,
    });
  }

  handleVoteSubmit(event, pollIndex) {
    event.preventDefault();
    if (!this.props.isAuthenticated) {
      this.props.history.push("/login");
      notification.info({
        message: "Polling App",
        description: "Please login to vote.",
      });

      return;
    }

    const poll = this.state.polls[pollIndex];
    const selectedChoice = this.state.currentVotes[pollIndex];

    const voteData = {
      question_id: poll.id,
      answer_id: selectedChoice,
      user_id: parseInt(localStorage.getItem("user_id"), 10),
    };
    castVote(voteData)
      .then((response) => {
        notification.info({
          message: "Vote Registered Successfully",
        });
        this.setState({
          polls: this.state.polls,
        });
        this.loadPollList();
      })
      .catch((error) => {
        if (error.status === 401) {
          this.props.handleLogout(
            "/login",
            "error",
            "You have been logged out. Please login to vote"
          );
        } else {
          notification.error({
            message: "Polling App",
            description:
              error.message || "Sorry! Something went wrong. Please try again!",
          });
        }
      });
  }

  render() {
    const pollViews = [];
    this.state.polls.forEach((poll, pollIndex) => {
      pollViews.push(
        <Poll
          key={poll.id}
          poll={poll}
          user={this.props.currentUser}
          votes={this.state.votes}
          totalVotes={this.state.totalVotes}
          totalVotesEachOption={this.state.totalVotesEachOption}
          currentVote={this.state.currentVotes[pollIndex]}
          handleVoteChange={(event) => this.handleVoteChange(event, pollIndex)}
          handleVoteSubmit={(event) => this.handleVoteSubmit(event, pollIndex)}
        />
      );
    });

    return (
      <div className="polls-container">
        {pollViews}
        {this.state.polls.length === 0 ? (
          <div className="no-polls-found">
            <span>No Polls Found.</span>
          </div>
        ) : null}
        {!this.state.isLoading && !this.state.last ? (
          <div className="load-more-polls">
            <Button
              type="dashed"
              onClick={this.handleLoadMore}
              disabled={this.state.isLoading}
            >
              <Icon type="plus" /> Load more
            </Button>
          </div>
        ) : null}
        {this.state.isLoading ? <LoadingIndicator /> : null}
      </div>
    );
  }
}

export default withRouter(PollList);
