import { Tag } from "~/types/tag";

export const useTag = () => {
  const { data: tags, error } = useFetch<Tag[]>(`http://localhost:8080/tags`);

  if (!tags.value || error.value) {
    throw createError({
      statusCode: 404,
      message: "Tags not found",
    });
  }

  return tags;
};
