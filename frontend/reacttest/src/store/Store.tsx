import { createStore } from 'redux';
import {IVote} from "../components/votes/IVote";

const saveVoteInDB = (vote: IVote) => {
  let m;
  let u;
  if (vote.Id) {
    m = 'PUT';
    u = '/admin/votes/'+vote.Id;
  } else {
    m = 'POST';
    u = '/admin/votes';
  }

  fetch(u, {
    body: $.param(vote, true),
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    },
    method: m
  }).then(res => res.json())
    .then(json => {
      if (!vote.Id) {
        updateVoteFromDB();
      }
    })
    .catch(error => error);
};

const deleteVoteInDB = (id: string) => {
  fetch('/admin/votes/'+id, {
    method: 'DELETE'
  }).then(res => res.json())
    .then(json => json)
    .catch(error => error);
};

const increaseVoteInDB = (id: string, index: number) => {
  fetch('/votes/'+id+'/inc/'+index, {
    method: 'PUT'
  }).then(res => res.json())
    .then(json => json)
    .catch(error => error);
};

const reducer = (state = [], action: any) => {

  if (!state) {
    state = [];
  }

  switch (action.type) {
    case UPDATE_VOTE:
      return action.data;
    case NEW_VOTE:
      saveVoteInDB(action.vote);
      const newState = JSON.parse(JSON.stringify(state));
      action.vote.Votes = Array(action.vote.Options.length).fill(0);
      newState.push(action.vote);
      return newState;
    case EDIT_VOTE:
      saveVoteInDB(action.vote);
      return state.map((v: IVote) => (
        v.Id === action.vote.Id ? action.vote : JSON.parse(JSON.stringify(v))
      ));
    case DELETE_VOTE:
      deleteVoteInDB(action.voteId);
      return state.filter((v: IVote) => (v.Id !== action.voteId));
    case INCREASE_VOTE:
      increaseVoteInDB(action.voteId, action.voteIndex);
      return state.map((v: IVote) => {
        if (v.Id === action.voteId) {
          const newVotes = v.Votes.slice();
          newVotes[action.voteIndex]++;
          const newVote = JSON.parse(JSON.stringify(v));
          newVote.Votes = newVotes;
          return newVote
        } else {
          return JSON.parse(JSON.stringify(v));
        }
      });
    default:
      return state;
  }
};

const Store = createStore(reducer);

const UPDATE_VOTE = "UPDATE_VOTE";

export const updateVoteFromDB = (id?: string) => {
  const url = id ? ('/votes/'+id) : '/admin/votes';
  fetch(url, {
    method: 'GET'
  }).then(res => res.json())
    .then(json => {
      Store.dispatch({type: UPDATE_VOTE, data: json})
    })
    .catch(error => error);
};

export const NEW_VOTE = "NEW_VOTE";
export const EDIT_VOTE = "EDIT_VOTE";
export const DELETE_VOTE = "DELETE_VOTE";
export const INCREASE_VOTE = "INCREASE_VOTE";
export default Store;