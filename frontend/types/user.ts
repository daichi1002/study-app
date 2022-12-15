import { Article } from "./Article";

export type User = {
  userId: string;
  userName: string;
  articles: Article[];
};

export type Login = {
  email: string;
  password: string;
};
