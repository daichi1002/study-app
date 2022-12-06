<script setup lang="ts">
import { Tag } from "~/types/tag";
import { Article } from "~/types/article";

const router = useRoute();
const { id } = router.params;

const { data: article } = await useFetch<Article>(
  `http://localhost:8080/article/${id}`
);

const { data: tags, error } = await useFetch<Tag[]>(
  `http://localhost:8080/article/create`
);

if (!tags.value || error.value) {
  throw createError({
    statusCode: 404,
    message: "Tags not found",
  });
}

const saveArticle = async () => {
  const { data, error } = await useFetch<Article>(
    `http://localhost:8080/article/edit/${id}`,
    {
      method: "PUT",
      body: tags,
    }
  );

  if (error.value) {
    throw createError({
      statusCode: 404,
      message: "failed to create",
    });
  }
};
</script>

<template>
  <div class="container mx-auto" v-if="article">
    <Form :value="article" :tags="tags" />
    <div class="mt-4 grid grid-cols-9 gap-9">
      <button
        class="bg-teal-300 hover:bg-teal-500 py-2 px-4 rounded-lg"
        @click="saveArticle()"
      >
        編集
      </button>
      <button class="bg-red-400 hover:bg-red-500 py-2 px-4 rounded-lg">
        削除
      </button>
    </div>
  </div>
</template>
