import { Article } from "./Article";

export type User = {
  userId: string;
  userName: string;
  articles: Article[];
  birthday: string;
  affiliation: string;
  languages: Languages[];
};

export type Login = {
  email: string;
  password: string;
};

export type Languages = {
  name: string;
  ratio: number;
};
