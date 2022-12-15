import { Article } from "./Article";

export type User = {
  userId: string;
  userName: string;
  articles: Article[];
  birthday: string;
  affiliation: string;
};

export type Login = {
  email: string;
  password: string;
};
