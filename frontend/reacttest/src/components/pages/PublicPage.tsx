import * as React from "react";
import {connect} from "react-redux";
import {INCREASE_VOTE, updateVoteFromDB} from "../../store/Store";
import {IVote} from "../votes/IVote";
import VoteItem from "../votes/VoteItem";

class PublicPage extends React.Component<any, any> {
  constructor(props: any) {
    super(props);

    updateVoteFromDB(props.match.params.id);
  }

  public render() {
    return (
      <div className="container">
        {
          this.props.vote ?
            <VoteItem vote={this.props.vote} onIncreaseVote={this.props.onIncreaseVote} /> :
            'Can\'t find vote \''+this.props.match.params.id+'\'!'
        }
      </div>
    )
  }
}

const mapStateToProps = (state: any, props: any) => {
  if (state === null) {
    return {};
  }
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