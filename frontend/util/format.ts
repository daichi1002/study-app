import { ArticleTag } from "~/types/article";

// 日付のフォーマット
export const formatDate = (date: string | null): string => {
  if (date == null) {
    return "";
  }
  const formatDate = new Date(date);
  return formatDate.toLocaleDateString();
};

// タグ名のフォーマット
export const setTagName = (tags: ArticleTag[]) => {
  let name = "";
  const len = tags.length;
  tags.forEach((tag, i) => {
    if (len == i + 1) {
      name += `${tag.tagName}`;
      return;
    }
    name += `${tag.tagName} / `;
  });
  return name;
};

// ユーザー名のフォーマット
export const setUserName = (id: string) => {
  const { users } = useUser();
  let name = "";
  users.value?.map((user) => {
    if (user.userId == id) {
      name = user.userName;
      return;
    }
  });
  return name;
};
