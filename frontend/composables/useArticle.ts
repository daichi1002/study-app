import type { Ref } from "vue";
import { Article, ArticleTag } from "~/types/article";

export const setTitle = (article: Ref<Article>) => {
  return (title: string | undefined) => {
    if (title == undefined) {
      return;
    }
    article.value.title = title;
  };
};

export const setTags = (article: Ref<Article>) => {
  return (tag: ArticleTag) => {
    article.value.tags.push(tag);
  };
};

export const setContent = (article: Ref<Article>) => {
  return (content: string | undefined) => {
    if (content == undefined) {
      return;
    }
    article.value.content = content;
  };
};

export const setArticle = (article: Ref<Article>) => {
  return (resArticle: Article) => {
    article.value = resArticle;
  };
};

export const resetArticle = (article: Ref<Article>) => () => {
  article.value = reactive<Article>({
    articleId: null,
    title: "",
    content: "",
    tags: [],
    updatedAt: null,
  });
};

export const useArticle = () => {
  const article = useState("article", () =>
    reactive<Article>({
      articleId: null,
      title: "",
      content: "",
      tags: [],
      updatedAt: null,
    })
  );
  return {
    article: readonly(article),
    setTitle: setTitle(article),
    setTags: setTags(article),
    setContent: setContent(article),
    setArticle: setArticle(article),
    resetArticle: resetArticle(article),
  };
};
