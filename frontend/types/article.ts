export type Article = {
  articleId: string | null;
  title: string;
  content: string;
  tags: ArticleTag[];
  userId: string;
  updatedAt: string | null;
};

export type ArticleTag = {
  tagId: number;
  tagName: string;
};
