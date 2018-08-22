import { createStore } from 'redux';
import {IVote} from "../components/votes/IVote";

const initialState = [{"Id":"5b7bf53dc54d5a9c407fe6a6","Question":"First test question","Options":["First option with 1 vote","Second option with 0 vote"],"Votes":[1,0]},{"Id":"5b7bf53dc54d5a9c407fe6a7","Question":"Second test question","Options":["First option with 0 vote","Second option with 1 vote","Third option with 8 vote"],"Votes":[0,1,8]}];

const reducer = (state = initialState, action: any) => {
  switch (action.type) {
    case NEW_VOTE:
      const newState = JSON.parse(JSON.stringify(state));
      action.vote.Votes = Array(action.vote.Options.length).fill(0);
      newState.push(action.vote);
      return newState;
    case EDIT_VOTE:
      return state.map((v: IVote) => (
        v.Id === action.vote.Id ? action.vote : JSON.parse(JSON.stringify(v))
      ));
    case DELETE_VOTE:
      return state.filter((v: IVote) => (v.Id !== action.voteId));
    case INCREASE_VOTE:
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

const updateStoreFromDB = () => {
  fetch('http://localhost:8080/admin/votes').then(res => {
    alert(JSON.stringify(res));
  })
};

const initializeStore = () => {

  updateStoreFromDB();

  return createStore(reducer);
};

const Store = initializeStore();

export const NEW_VOTE = "NEW_VOTE";
export const EDIT_VOTE = "EDIT_VOTE";
export const DELETE_VOTE = "DELETE_VOTE";
export const INCREASE_VOTE = "INCREASE_VOTE";
export default Store;