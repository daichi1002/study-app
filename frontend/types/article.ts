export type Article = {
  articleId: string | null;
  title: string;
  content: string;
  tags: ArticleTag[];
  updatedAt: string | null;
};

export type ArticleTag = {
  tagId: number;
  tagName: string;
};
