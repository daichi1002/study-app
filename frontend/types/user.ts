import { Article } from "./Article";

export type User = {
  userId: string;
  userName: string;
  articles: Article[];
};
