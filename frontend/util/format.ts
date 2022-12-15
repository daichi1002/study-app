import { ArticleTag } from "~/types/article";
import { User } from "~/types/user";

// 日付のフォーマット
export const formatDate = (date: string | null | undefined): string => {
  if (date == null || date == undefined) {
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
export const setUserName = (id: string, users: User[] | null) => {
  let name = "";
  users?.map((user) => {
    if (user.userId == id) {
      name = user.userName;
      return;
    }
  });
  return name;
};
