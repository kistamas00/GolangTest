import * as React from "react";
import {connect} from "react-redux";
import {INCREASE_VOTE} from "../../store/Store";
import {IVote} from "../votes/IVote";
import VoteItem from "../votes/VoteItem";

const PublicPage = (props: any) => (
  <div className="container">
    {
      props.vote ?
        <VoteItem vote={props.vote} onIncreaseVote={props.onIncreaseVote} /> :
        'Can\'t find vote \''+props.match.params.id+'\'!'
    }
  </div>
);

const mapStateToProps = (state: any, props: any) => {
  const vote = state.filter((v: IVote) => (v.Id === props.match.params.id));
  if (vote.length === 1) {
    return {vote: vote[0]};
  } else {
    return {};
  }
};

const mapDispatchToProps = (dispatch: any) => ({
  onIncreaseVote : (id: string, index: number) => {
    const action = {type: INCREASE_VOTE, voteId: id, voteIndex: index};
    dispatch(action);
  }
});

export default connect(mapStateToProps, mapDispatchToProps)(PublicPage);