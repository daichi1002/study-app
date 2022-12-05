export type Article = {
  articleId: string | null;
  title: string;
  content: string;
  tags: ArticleTag[];
  updatedAt: Date | null;
};

export type ArticleTag = {
  tagId: number;
  tagName: string;
};
