import * as React from 'react';
import {connect} from 'react-redux';
import {default as Store, DELETE_VOTE, EDIT_VOTE, NEW_VOTE} from "../../store/Store";
import Modal from "../modal/Modal";
import NavBar from '../navbar/Navbar';
import {IVote} from "../votes/IVote";
import VoteItem from '../votes/VoteItem';

class AdminPage extends React.Component<any, any> {
  constructor(props: any) {
    super(props);

    this.state = {editVote: {}};

    this.onEditVoteShowModal = this.onEditVoteShowModal.bind(this);
    this.onDeleteVote = this.onDeleteVote.bind(this);
  }

  public render() {
    return (
      <div>
        <NavBar target="newItemModalCenter"/>
        <Modal modalId="newItemModalCenter" vote={this.state.editVote as IVote} onSaveVote={this.props.onSaveVote} />
        <div className="container">
          {
            this.props.votes.map((v: any, i: number) => {
              const key = v.Id ? v.Id : i;
              return (
                <VoteItem key={key} vote={v} onEditVoteShowModal={this.onEditVoteShowModal} onDeleteVote={this.onDeleteVote}/>
              )
            })
          }
        </div>
      </div>
    )
  }

  private onEditVoteShowModal(id: string) {
    const vote = Store.getState().filter((v: IVote) => (v.Id === id));
    if (vote.length === 1) {
      this.setState({editVote: vote[0]});
    }
  }

  private onDeleteVote(id: string) {
    this.setState({editVote: {}});
    const action = { type: DELETE_VOTE, voteId: id };
    Store.dispatch(action);
  }
};

const mapStateToProps = (state: any) => ({
  votes: state
});

const mapDispatchToProps = (dispatch: any) => ({
  onSaveVote: (v: IVote) => {
    if (v.Id) {
      const action = { type: EDIT_VOTE, vote: v };
      dispatch(action);
    } else {
      const action = {type: NEW_VOTE, vote: v};
      dispatch(action);
    }
  }
});

export default connect(mapStateToProps, mapDispatchToProps)(AdminPage);