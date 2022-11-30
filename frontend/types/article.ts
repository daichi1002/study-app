export type Article = {
  articleId: string;
  userName: string;
  title: string;
  content: string;
  tag: string;
  updatedDate: string;
};

export type RequestArticle = {
  title: string;
  content: string;
  tags: ArticleTag[];
  userId: string;
};

type ArticleTag = {
  tagId: number;
};
