import {IVote} from "../votes/IVote";

export interface IModal {
  modalId: string,
  vote: IVote,
  onSaveVote: (onNewVote: IVote) => void
}
