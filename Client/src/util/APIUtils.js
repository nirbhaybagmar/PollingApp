import { API_BASE_URL } from "../constants";

const request = (options) => {
  const headers = new Headers({
    "Content-Type": "application/json",
  });

  const defaults = { headers: headers };
  options = Object.assign({}, defaults, options);

  return fetch(options.url, options).then((response) =>
    response.json().then((json) => {
      if (!response.ok) {
        return Promise.reject(json);
      }
      return json;
    })
  );
};

export function getAllPolls() {
  return request({
    url: API_BASE_URL + "/polls",
    method: "GET",
  });
}

export function createPoll(pollData) {
  return request({
    url: API_BASE_URL + "/polls",
    method: "POST",
    body: JSON.stringify(pollData),
  });
}

export function castVote(voteData) {
  return request({
    url: API_BASE_URL + "/polls/vote",
    method: "POST",
    body: JSON.stringify(voteData),
  });
}

export function login(loginRequest) {
  return request({
    url: API_BASE_URL + "/auth/login",
    method: "POST",
    body: JSON.stringify(loginRequest),
  });
}

export function signup(signupRequest) {
  console.log(signupRequest);
  return request({
    url: API_BASE_URL + "/auth/signup",
    method: "POST",
    body: JSON.stringify(signupRequest),
  });
}

export function getCurrentUser() {
  var userId = parseInt(localStorage.getItem("user_id"), 10);
  return request({
    url: API_BASE_URL + "/user/?id=" + userId,
    method: "GET",
  });
}

export function getUserProfile(username) {
  return request({
    url: API_BASE_URL + "/users/" + username,
    method: "GET",
  });
}

export function getUserCreatedPolls(userID) {
  return request({
    url: API_BASE_URL + "/polls/?id=" + userID,
    method: "GET",
  });
}

export function getUserVotedPolls(userId) {
  return request({
    url: API_BASE_URL + "/votes/?id=" + userId,
    method: "GET",
  });
}

export function getVotes() {
  return request({
    url: API_BASE_URL + "/votes",
    method: "GET",
  });
}
