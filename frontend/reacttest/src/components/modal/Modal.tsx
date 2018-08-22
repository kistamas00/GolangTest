import * as React from 'react';
import {IVote} from "../votes/IVote";
import {IModal} from "./IModal";

class Modal extends React.Component<IModal, IVote> {
  constructor(props: IModal) {
    super(props);

    this.state = Object.assign({Question: '', Options: ['', '', '', '']}, this.props.vote);

    this.handleChange = this.handleChange.bind(this);
    this.handleSave = this.handleSave.bind(this);
    this.handleCancel = this.handleCancel.bind(this);
  }

  public componentWillReceiveProps(nextProps: IModal) {

    if (Object.keys(nextProps.vote).length !== 0) {

      this.setState(JSON.parse(JSON.stringify(nextProps.vote)));
      $('#' + this.props.modalId).modal('show');
    }
  }

  public render() {
    return (
      <div className="modal fade" id={this.props.modalId} tabIndex={-1} role="dialog" aria-labelledby="modalCenterTitle" aria-hidden="true">
        <div className="modal-dialog modal-dialog-centered" role="document">
          <div className="modal-content">
            <div className="modal-header">
              <h5 className="modal-title" id="modalLongTitle">{this.state.Id}</h5>
            </div>
            <div className="modal-body">
              <div className="form-group">
                <label>Question:</label>
                <input type="text" className="form-control" placeholder="e.g. 2x3=?"
                  onChange={this.handleChange} name="Question" value={this.state.Question} />
                <label>Options:</label>
                {
                  this.state.Options.map((o, i) => (
                    <input key={i} type="text" className="form-control" placeholder={'Option #'+(i+1)}
                           onChange={this.handleChange} name={'Options#'+i} value={o} />
                  ))
                }
              </div>
            </div>
            <div className="modal-footer">
              <button type="button" className="btn btn-danger" data-dismiss="modal" onClick={this.handleCancel}>Cancel</button>
              <button type="button" className="btn btn-success" data-dismiss="modal" onClick={this.handleSave}>Save</button>
            </div>
          </div>
        </div>
      </div>
    )
  }

  private handleChange(event: any) {
    const newState = {};
    const name = event.target.name.split('#');
    if (name.length === 1) {
      newState[name[0]] = event.target.value;
    } else {
      newState[name[0]] = this.state[name[0]];
      newState[name[0]][name[1]] = event.target.value;
    }
    this.setState(newState);
  }

  private handleSave(event: any) {

    const newVote = JSON.parse(JSON.stringify(this.state));
    const newOptions = [];
    for (const option of newVote.Options) {
      if (option.length > 0) {
        newOptions.push(option);
      }
    }
    newVote.Options = newOptions;

    if (newVote.Question.length > 0 && newVote.Options.length >= 2) {
      this.props.onSaveVote(newVote);
    }

    const newState = {Question: '', Options: ['','','','']};
    this.setState(newState);
  }

  private handleCancel(event: any) {
    const newState = {Question: '', Options: ['','','',''], Id: ''};
    this.setState(newState);
  }
}

export default Modal;