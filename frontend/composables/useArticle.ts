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

  const setTags = (tags: ArticleTag[]) => {
    // タグが既存の配列になければ、追加・あれば削除のロジックを書く
    article.value.tags = tags;
    console.log(tags);
    // console.log(article.value.tags);
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
