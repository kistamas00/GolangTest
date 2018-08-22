import {IVote} from "./IVote";

export interface IVoteItem {
  vote: IVote,
  onEditVoteShowModal?: (id: string) => void,
  onDeleteVote?: (id: string) => void,
  onIncreaseVote?: (id: string, index: number) => void
}