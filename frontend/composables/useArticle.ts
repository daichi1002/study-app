import type { Ref } from "vue";
import { Article, ArticleTag } from "~/types/article";

export const useArticle = () => {
  const article = useState<Article>("article", () =>
    reactive<Article>({
      articleId: null,
      title: "",
      content: "",
      tags: [],
      userId: "",
      updatedAt: null,
    })
  );

  const setTitle = (title: string | undefined) => {
    if (title == undefined) {
      return;
    }
    article.value.title = title;
  };

  const setTags = (t: ArticleTag) => {
    // TODO リファクタ
    if (article.value.tags.some((a) => a.tagId == t.tagId)) {
      article.value.tags = article.value.tags.filter(
        (item) => item.tagId !== t.tagId
      );
    } else {
      article.value.tags.push(t);
    }
  };

  const setContent = (content: string | undefined) => {
    if (content == undefined) {
      return;
    }
    article.value.content = content;
  };

  const setUserId = (userId: string | undefined) => {
    if (userId == undefined) {
      return;
    }
    article.value.userId = userId;
  };

  const setArticle = (resArticle: Article | null): Article => {
    if (resArticle != null) {
      article.value = resArticle;
    }
    return article.value;
  };

  const resetArticle = () => {
    article.value = reactive<Article>({
      articleId: null,
      title: "",
      content: "",
      tags: [],
      userId: "",
      updatedAt: null,
    });
  };

  return {
    article: readonly(article),
    setTitle,
    setTags,
    setContent,
    setUserId,
    setArticle,
    resetArticle,
  };
};
