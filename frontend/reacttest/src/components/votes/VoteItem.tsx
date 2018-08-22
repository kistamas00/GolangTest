import * as React from 'react';
import {IVoteItem} from "./IVoteItem";

const VoteItem = (props: IVoteItem) => (
  <table className="table table-bordered table-hover">
    <thead className="thead-light">
      <tr>
        <th>
          <a href={'/letsVote/'+props.vote.Id}>{props.vote.Question} ({props.vote.Id})</a>
          <span className="float-right">
            {
              editVoteShowModal(props) &&
              <button className="fas fa-pencil-alt text-warning btn btn-light"
                      onClick={editVoteShowModal(props)} />
            }
            {
              deleteVote(props) &&
              <button className="fas fa-trash-alt text-danger btn btn-light" style={{marginLeft: 10}}
                      onClick={deleteVote(props)} />
            }
          </span>
        </th>
      </tr>
    </thead>
    <tbody>
    {
      props.vote.Options.map((o, i) => (
        <tr key={i} onClick={increaseVote(props, i)}>
          <td>{o} ({props.vote.Votes[i]})</td>
        </tr>
      ))
    }
    </tbody>
  </table>
);

const editVoteShowModal = (v: IVoteItem) => {

  if (!v.onEditVoteShowModal) {
    return;
  }

  return () => {
    if (v.onEditVoteShowModal) {
      v.onEditVoteShowModal(v.vote.Id);
    }
  }
};

const deleteVote = (v: IVoteItem) => {

  if (!v.onDeleteVote) {
    return;
  }

  return () => {
    if (v.onDeleteVote) {
      v.onDeleteVote(v.vote.Id);
    }
  }
};

const increaseVote = (v: IVoteItem, i: number) => {
  return () => {
    if (v.onIncreaseVote) {
      v.onIncreaseVote(v.vote.Id, i);
    }
  }
};

export default VoteItem;